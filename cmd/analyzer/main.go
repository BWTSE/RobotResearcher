package main

import (
	"context"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/mohae/struct2csv"

	"github.com/pkg/errors"

	"sidus.io/robotresearcher/internal/database"
)

var rootDir = "submissions"

func structToFile(a interface{}, file string) error {
	data, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return errors.Wrap(err, "could not marshal struct")
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return errors.Wrap(err, "could not write file")
	}

	return nil
}

type Session struct {
	SignupCode        string               `json:"signup_code"`
	BackgroundAnswers BackgroundSubmission `json:"background_answers"`
	RepeatOf          string               `json:"repeat_of"`
}

type BackgroundSubmission struct {
	EducationLevel string `json:"education_level"`
	EducationField string `json:"education_field"`

	ProgrammingExperience string `json:"work_experience"`
	JavaExperience        string `json:"work_experience_java"`

	WorkDomain string `json:"work_domain"`

	CompanyUsesCodeReviews     bool `json:"workplace_peer_review"`
	CompanyUsesPairProgramming bool `json:"workplace_pair_programming"`
	CompanyTracksTechnicalDebt bool `json:"workplace_td_tracking"`
	CompanyHasCodingStandards  bool `json:"workplace_coding_standards"`
}

type SurveySubmission struct {
	ScenarioQuality   int `json:"quality_pre_task"`
	SubmissionQuality int `json:"quality_post_task"`
}

type Submission struct {
	NotSubmitted      bool              `json:"not_submitted"`
	HighDebtVersion   bool              `json:"high_debt_version"`
	Time              time.Duration     `json:"time"`
	Order             int               `json:"order"`
	ReflectionAnswers *SurveySubmission `json:"reflection_answers"`
}

type SubmissionInspection struct {
	TaskCompletion string `json:"task_completion"`

	ReusedLogicConstructor bool `json:"reused_logic_constructor"`
	ReusedLogicValidation  bool `json:"reused_logic_validation"`
	ReusedLogicEquals      bool `json:"reused_logic_equals"`
	ReusedLogicHashcode    bool `json:"reused_logic_hashcode"`

	CopiedVariableNamesAll []string `json:"copied_variable_names_all"`
	CopiedVariableNamesBad []string `json:"copied_variable_names_bad"`
	NewVariableNamesAll    []string `json:"new_variable_names_all"`
	NewVariableNamesBad    []string `json:"new_variable_names_bad"`
	EditedVariableNamesAll []string `json:"edited_variable_names_all"`
	EditedVariableNamesBad []string `json:"edited_variable_names_bad"`

	HasEquals   bool `json:"has_equals"`
	HasHashCode bool `json:"has_hashcode"`

	DocumentationState string `json:"documentation_state"`

	LargeStructureChange bool `json:"large_structure_change"`

	Done bool `json:"inspection_done"`
}

type SessionInspection struct {
	WorkExperience     float64 `json:"work_experience"`
	WorkExperienceJava float64 `json:"work_experience_java"`

	Ignore bool `json:"ignore"`

	Done bool `json:"inspection_done"`
}

type Issue struct {
	Rule     string `json:"rule"`
	File     string `json:"file"`
	Severity string `json:"severity"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	Line     int    `json:"line"`
}

type Rule struct {
	Count          int    `json:"count"`
	Severity       string `json:"severity"`
	Type           string `json:"type"`
	ExampleMessage string `json:"example_message"`
	Ignored        bool   `json:"ignored"`
	IgnoredComment string `json:"ignored_comment"`
}

type Groupings struct {
	EducationFields map[string]string `json:"education_fields"`
	WorkDomains     map[string]string `json:"work_domains"`
	SignupCodes     map[string]string `json:"signup_codes"`
}

type TableRow struct {
	Session                    string  `csv:"session"`
	Group                      string  `csv:"group"`
	EducationLevel             string  `csv:"education_level"`
	EducationField             string  `csv:"education_field"`
	ProgrammingExperience      float64 `csv:"work_experience_programming"`
	JavaExperience             float64 `csv:"work_experience_java"`
	WorkDomain                 string  `csv:"work_domain"`
	CompanyUsesCodeReviews     bool    `csv:"workplace_peer_review"`
	CompanyUsesPairProgramming bool    `csv:"workplace_pair_programming"`
	CompanyTracksTechnicalDebt bool    `csv:"workplace_td_tracking"`
	CompanyHasCodingStandards  bool    `csv:"workplace_coding_standards"`

	Submission      string `csv:"scenario"`
	Order           int    `csv:"order"`
	HighDebtVersion bool   `csv:"high_debt_version"`
	QualityPreTask  int    `csv:"quality_pre_task"`
	QualityPostTask int    `csv:"quality_post_task"`
	Time            int    `csv:"time"`

	SonarqubeIssuesMajor    int `csv:"sonarqube_issues_major"`
	SonarqubeIssuesInfo     int `csv:"sonarqube_issues_info"`
	SonarqubeIssuesMinor    int `csv:"sonarqube_issues_minor"`
	SonarqubeIssuesCritical int `csv:"sonarqube_issues_critical"`

	ModifiedLines           int    `csv:"modified_lines"`
	TaskCompletion          string `csv:"task_completion"`
	ReusedLogicConstructor  bool   `csv:"reused_logic_constructor"`
	ReusedLogicValidation   bool   `csv:"reused_logic_validation"`
	CopiedVariableNamesAll  int    `csv:"var_names_copied_all"`
	CopiedVariableNamesGood int    `csv:"var_names_copied_good"`
	NewVariableNamesAll     int    `csv:"var_names_new_all"`
	NewVariableNamesGood    int    `csv:"var_names_new_good"`
	EditedVariableNamesAll  int    `csv:"var_names_edited_all"`
	EditedVariableNamesGood int    `csv:"var_names_edited_good"`
	EqualsState             string `csv:"equals_state"`
	HashCodeState           string `csv:"hashcode_state"`
	DocumentationState      string `csv:"documentation_state"`
	LargeStructureChange    bool   `csv:"large_structure_change"`
}

func fileExists(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		panic(err)
	}

	f.Close()

	return true
}

func download() error {
	db, err := database.NewDatabase(os.Getenv("DB_URI"))
	if err != nil {
		return fmt.Errorf("while creating code service: %w", err)
	}
	defer db.Close()

	sessions, err := db.GetRealSessions()
	if err != nil {
		return errors.Wrap(err, "while getting sessions")
	}

	fmt.Printf("Total sessions found: %d\n", len(sessions))

	var startedSessions []database.Session

	for _, session := range sessions {
		if !session.Scenarios[0].StartedAt.IsZero() {
			startedSessions = append(startedSessions, session)
		}
	}

	fmt.Printf("After filtering out not started: %d\n", len(startedSessions))

	for _, session := range startedSessions {
		sessionDir := path.Join(rootDir, session.ID.Hex())

		if (session.Scenarios[0].SubmittedAt.IsZero() || session.Scenarios[1].SubmittedAt.IsZero()) &&
			time.Since(session.StartedAt) < 24*time.Hour {
			fmt.Printf("download skipped (session was opened less than 24 hours ago and not yet submitted): %s\n", session.ID.Hex())
			continue
		}

		if fileExists(sessionDir) { // dir exists
			fmt.Printf("download skipped (already exists): %s\n", session.ID.Hex())
			continue
		}

		fmt.Printf("download: %s\n", sessionDir)

		err = os.Mkdir(sessionDir, 0775)
		if err != nil {
			return errors.Wrap(err, "could not create session directory")
		}

		err = structToFile(Session{
			SignupCode: session.RegisterCode,
			BackgroundAnswers: BackgroundSubmission{
				EducationLevel:             session.BackgroundAnswers.EducationLevel,
				EducationField:             session.BackgroundAnswers.EducationField,
				ProgrammingExperience:      session.BackgroundAnswers.ProgrammingExperience,
				JavaExperience:             session.BackgroundAnswers.JavaExperience,
				WorkDomain:                 session.BackgroundAnswers.WorkDomain,
				CompanyUsesCodeReviews:     session.BackgroundAnswers.CompanyUsesCodeReviews,
				CompanyUsesPairProgramming: session.BackgroundAnswers.CompanyUsesPairProgramming,
				CompanyTracksTechnicalDebt: session.BackgroundAnswers.CompanyTracksTechnicalDebt,
				CompanyHasCodingStandards:  session.BackgroundAnswers.CompanyHasCodingStandards,
			},
			RepeatOf: session.FirstID,
		}, path.Join(sessionDir, ".session.json"))
		if err != nil {
			return err
		}

		for i, scenario := range session.Scenarios {
			scenarioDir := path.Join(sessionDir, scenario.Name)
			srcDir := path.Join(scenarioDir, ".submission", scenario.Name)

			err = os.MkdirAll(scenarioDir, 0775)
			if err != nil {
				return errors.Wrap(err, "could not create scenario directory")
			}

			if !scenario.SubmittedAt.IsZero() {
				err = os.MkdirAll(srcDir, 0775)
				if err != nil {
					return errors.Wrap(err, "could not create submission directory")
				}

				// save all submitted files
				for name, content := range scenario.Submitted {
					if strings.Contains(name, string(os.PathSeparator)) {
						break
					}

					if !strings.HasSuffix(name, ".java") {
						break
					}

					if content == "" {
						break
					}

					code, err := base64.StdEncoding.DecodeString(content)
					if err != nil {
						fmt.Println(err)
						break
					}

					code = []byte(strings.ToValidUTF8(string(code), "?"))

					err = ioutil.WriteFile(path.Join(srcDir, name), code, 0644)
					if err != nil {
						fmt.Println(err)
						break
					}
				}
			}

			var reflectionAnswers SurveySubmission

			if session.SurveyAnswers != nil {
				if scenario.Name == "tickets" {
					reflectionAnswers.ScenarioQuality = session.SurveyAnswers.ScenarioTicketsQuality
					reflectionAnswers.SubmissionQuality = session.SurveyAnswers.ScenarioTicketsQualityChange
				} else if scenario.Name == "booking" {
					reflectionAnswers.ScenarioQuality = session.SurveyAnswers.ScenarioBookingQuality
					reflectionAnswers.SubmissionQuality = session.SurveyAnswers.ScenarioBookingQualityChange
				} else {
					panic("very fatal")
				}
			}

			err = structToFile(Submission{
				HighDebtVersion:   scenario.HasHighDebt,
				NotSubmitted:      scenario.SubmittedAt.IsZero(),
				Time:              scenario.SubmittedAt.Sub(scenario.StartedAt),
				Order:             i,
				ReflectionAnswers: &reflectionAnswers,
			}, path.Join(scenarioDir, ".submission.json"))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func forEachSession(f func(string, Session) error) error {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			sessionPath := path.Join(rootDir, file.Name())

			sessionData, err := ioutil.ReadFile(path.Join(sessionPath, ".session.json"))
			if err != nil {
				return err
			}

			var session Session

			err = json.Unmarshal(sessionData, &session)
			if err != nil {
				return err
			}

			err = f(sessionPath, session)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func forEachSubmission(f func(string, Session, Submission) error) error {
	return forEachSession(func(sessionPath string, session Session) error {
		for _, submissionName := range []string{"booking", "tickets"} {
			submissionPath := path.Join(sessionPath, submissionName)

			if !fileExists(submissionPath) {
				continue
			}

			submissionData, err := ioutil.ReadFile(path.Join(submissionPath, ".submission.json"))
			if err != nil {
				return err
			}

			var submission Submission
			err = json.Unmarshal(submissionData, &submission)
			if err != nil {
				return err
			}

			err = f(submissionPath, session, submission)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func diff() error {
	return forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		if fileExists(path.Join(submissionPath, ".diff")) || fileExists(path.Join(submissionPath, ".diff.error")) {
			fmt.Printf("diff skipped: %s\n", submissionPath)
			return nil
		}

		if submission.NotSubmitted {
			fmt.Printf("diff skipped: %s\n", submissionPath)
			return nil
		}

		cmd := exec.Command(
			"diff",
			"-rdBEZN",
			path.Join(
				"..",
				"..",
				"..",
				"Scenarios",
				path.Base(submissionPath),
				debtString(submission.HighDebtVersion)+"_debt",
				path.Base(submissionPath),
			),
			path.Join(".submission", path.Base(submissionPath)),
		)
		cmd.Dir = submissionPath

		fmt.Printf("diff: %s\n", submissionPath)

		out, err := cmd.CombinedOutput()
		if err != nil {
			exitCodeError, ok := err.(*exec.ExitError)
			if ok {
				if exitCodeError.ExitCode() == 1 {
					err = ioutil.WriteFile(path.Join(submissionPath, ".diff"), out, 0644)
					if err != nil {
						return err
					}
				} else {
					err := ioutil.WriteFile(
						path.Join(submissionPath, ".diff.error"),
						[]byte(fmt.Sprintf(
							"exit code: %d\n%s",
							exitCodeError.ExitCode(),
							string(out),
						)),
						0644,
					)
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			err = ioutil.WriteFile(path.Join(submissionPath, ".diff"), out, 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func compile() error {
	return forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		srcDir := path.Join(submissionPath, ".submission", path.Base(submissionPath))

		if fileExists(path.Join(submissionPath, ".out")) || fileExists(path.Join(submissionPath, ".compilation.error")) {
			fmt.Printf("compilation skipped (.out or .compilation.error present): %s\n", submissionPath)
			return nil
		}

		if submission.NotSubmitted {
			fmt.Printf("compilation skipped: %s\n", submissionPath)
			return nil
		}

		fmt.Printf("compilation: %s\n", submissionPath)

		files, err := ioutil.ReadDir(srcDir)
		if err != nil {
			return err
		}

		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".java") {
				continue
			}

			cmd := exec.Command(
				"javac",
				"-d", "../.out",
				path.Join(path.Base(submissionPath), file.Name()),
			)
			cmd.Dir = path.Join(submissionPath, ".submission")

			out, err := cmd.CombinedOutput()
			if err != nil {
				exitCodeError, ok := err.(*exec.ExitError)
				if ok {
					err := ioutil.WriteFile(
						path.Join(submissionPath, ".compilation.error"),
						[]byte(fmt.Sprintf(
							"exit code: %d\n%s",
							exitCodeError.ExitCode(),
							string(out),
						)),
						0644,
					)
					if err != nil {
						return err
					}

					return nil
				} else {
					return err
				}
			}
		}

		return nil
	})
}

func execution() error {
	return forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		if fileExists(path.Join(submissionPath, ".execution")) || fileExists(path.Join(submissionPath, ".execution.error")) {
			fmt.Printf("execution skipped (exection or execution.error already exists): %s\n", submissionPath)
			return nil
		}

		if fileExists(path.Join(submissionPath, ".compilation.error")) {
			fmt.Printf("execution skipped (compilation failed): %s\n", submissionPath)
			return nil
		}

		if submission.NotSubmitted {
			fmt.Printf("compilation skipped: %s\n", submissionPath)
			return nil
		}

		if !fileExists(path.Join(submissionPath, ".compilation.error")) && !fileExists(path.Join(submissionPath, ".out")) {
			return errors.New("must compile before execution")
		}

		fmt.Printf("execution: %s\n", submissionPath)

		cmd := exec.CommandContext(
			ctx,
			"java",
			"-Djava.security.manager",
			//			"-Djava.security.debug=all",
			//			"-Djava.security.policy==./../../untrusted.policy",
			"-Xmx15M",
			path.Base(submissionPath)+".Main",
		)
		cmd.Dir = path.Join(submissionPath, ".out")

		out, err := cmd.CombinedOutput()
		if err != nil {
			exitCodeError, ok := err.(*exec.ExitError)
			if ok {
				err := ioutil.WriteFile(
					path.Join(submissionPath, ".execution.error"),
					[]byte(fmt.Sprintf(
						"exit code: %d\n%s",
						exitCodeError.ExitCode(),
						string(out),
					)),
					0644,
				)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			err = ioutil.WriteFile(path.Join(submissionPath, ".execution"), out, 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func scan() error {
	return forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		sessionPath, submissionName := path.Split(submissionPath)

		if submission.NotSubmitted {
			fmt.Printf("scanning skipped (not submitted): %s\n", submissionPath)
			return nil
		}

		if !fileExists(path.Join(submissionPath, ".out")) && !fileExists(path.Join(submissionPath, ".compilation.error")) {
			return errors.New("must compile before scan")
		}

		if fileExists(path.Join(submissionPath, ".compilation.error")) {
			fmt.Printf("scanning skipped (was not compiled successfully): %s\n", submissionPath)
			return nil
		}

		if fileExists(path.Join(submissionPath, ".sonarscanner")) || fileExists(path.Join(submissionPath, ".sonarscanner.error")) {
			fmt.Printf("scanning skipped (.sonarscanner or .sonarscanner.error present): %s\n", submissionPath)
			return nil
		}

		fmt.Printf("scanning: %s\n", submissionPath)

		cmd := exec.Command(
			"sonar-scanner",
			fmt.Sprintf("-Dsonar.projectKey=%s-%s-%s", path.Base(submissionPath), debtString(submission.HighDebtVersion), path.Base(sessionPath)),
			"-Dsonar.scm.exclusions.disabled=true",
			"-Dsonar.java.binaries="+path.Join(".out", submissionName),
			"-Dsonar.sources="+path.Join(".submission", submissionName),
		)
		cmd.Dir = path.Join(submissionPath)

		out, err := cmd.CombinedOutput()
		if err != nil {
			exitCodeError, ok := err.(*exec.ExitError)
			if ok {
				err := ioutil.WriteFile(
					path.Join(submissionPath, ".sonarscanner.error"),
					[]byte(fmt.Sprintf(
						"exit code: %d\n%s",
						exitCodeError.ExitCode(),
						string(out),
					)),
					0644,
				)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			err = ioutil.WriteFile(path.Join(submissionPath, ".sonarscanner"), out, 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func issues() error {
	client := http.DefaultClient

	return forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		if fileExists(path.Join(submissionPath, ".issues.json")) || fileExists(path.Join(submissionPath, ".issues.error")) {
			fmt.Printf("issues skipped (.issues.json or .issues.error present): %s\n", submissionPath)
			return nil
		}

		if !fileExists(path.Join(submissionPath, ".sonarscanner")) {
			fmt.Printf("issues skipped (no scan output): %s\n", submissionPath)
			return nil
		}

		if submission.NotSubmitted {
			fmt.Printf("issues skipped: %s\n", submissionPath)
			return nil
		}

		fmt.Printf("issues: %s\n", submissionPath)

		sessionPath, _ := path.Split(submissionPath)

		projectKey := fmt.Sprintf("%s-%s-%s", path.Base(submissionPath), debtString(submission.HighDebtVersion), path.Base(sessionPath))

		resp, err := client.Get(fmt.Sprintf("https://sonarqube.sidus.io/api/issues/search?componentKeys=%s&ps=500&p=1", projectKey))
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(
				path.Join(submissionPath, ".issues.error"),
				[]byte(fmt.Sprintf(
					"status code: %d\n%s",
					resp.StatusCode,
					string(bytes),
				)),
				0644,
			)

			return err
		}

		var response struct {
			TotalItems int `json:"total"`
			SentItems  int `json:"ps"`
			Issues     []struct {
				Rule      string `json:"rule"`
				Component string `json:"component"`
				Message   string `json:"message"`
				Severity  string `json:"severity"`
				Type      string `json:"type"`
				Line      int    `json:"line"`
			} `json:"issues"`
		}

		dec := json.NewDecoder(resp.Body)

		err = dec.Decode(&response)
		if err != nil {
			return err
		}

		if response.SentItems < response.TotalItems {
			return errors.New("page size exceeded")
		}

		var issuesToSave []Issue
		for _, issue := range response.Issues {
			componentParts := strings.Split(issue.Component, ":")
			var fileName string
			if len(componentParts) == 0 {
				return errors.New("not two components: " + issue.Component)
			} else if len(componentParts) == 2 {
				_, fileName = path.Split(componentParts[1])
			} else {
				fileName = componentParts[0]
			}

			if strings.ToLower(fileName) != "main.java" {
				issuesToSave = append(issuesToSave, Issue{
					Rule:     issue.Rule,
					File:     fileName,
					Severity: issue.Severity,
					Type:     issue.Type,
					Message:  issue.Message,
					Line:     issue.Line,
				})
			}
		}

		err = structToFile(issuesToSave, path.Join(submissionPath, ".issues.json"))
		if err != nil {
			return err
		}
		return nil
	})
}

func issuesRules() error {
	rules := make(map[string]Rule)

	if fileExists(path.Join(rootDir, ".manual_rules.json")) {
		data, err := ioutil.ReadFile(path.Join(rootDir, ".manual_rules.json"))
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &rules)
		if err != nil {
			return err
		}
	}

	for key, rule := range rules {
		rule.Count = 0
		rules[key] = rule
	}

	err := forEachSubmission(func(submissionPath string, session Session, submisison Submission) error {
		if !fileExists(path.Join(submissionPath, ".issues.json")) {
			return nil
		}

		data, err := ioutil.ReadFile(path.Join(submissionPath, ".issues.json"))
		if err != nil {
			return err
		}

		var issues []Issue

		err = json.Unmarshal(data, &issues)
		if err != nil {
			return err
		}

		for _, issue := range issues {
			rule := rules[issue.Rule]
			rule.Severity = issue.Severity
			rule.Type = issue.Type
			if rule.ExampleMessage == "" && issue.Message != "" {
				rule.ExampleMessage = issue.Message
			}
			rule.Count++
			rules[issue.Rule] = rule
		}
		return nil
	})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(rules, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(rootDir, ".manual_rules.json"), data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func groups() error {
	groupings := Groupings{
		EducationFields: make(map[string]string),
		WorkDomains:     make(map[string]string),
		SignupCodes:     make(map[string]string),
	}

	if fileExists(path.Join(rootDir, ".manual_groupings.json")) {
		data, err := ioutil.ReadFile(path.Join(rootDir, ".manual_groupings.json"))
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &groupings)
		if err != nil {
			return err
		}
	}

	err := forEachSession(func(s string, session Session) error {
		if groupings.WorkDomains[session.BackgroundAnswers.WorkDomain] == "" {
			groupings.WorkDomains[session.BackgroundAnswers.WorkDomain] = "TODO"
		}

		if groupings.EducationFields[session.BackgroundAnswers.EducationField] == "" {
			groupings.EducationFields[session.BackgroundAnswers.EducationField] = "TODO"
		}

		if groupings.SignupCodes[session.SignupCode] == "" {
			groupings.SignupCodes[session.SignupCode] = "TODO"
		}
		return nil
	})
	if err != nil {
		return err
	}

	return structToFile(groupings, path.Join(rootDir, ".manual_groupings.json"))
}

func submissionInspection() error {
	return forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		inspection := SubmissionInspection{}

		if fileExists(path.Join(submissionPath, ".manual_inspection.json")) {
			data, err := ioutil.ReadFile(path.Join(submissionPath, ".manual_inspection.json"))
			if err != nil {
				return err
			}

			err = json.Unmarshal(data, &inspection)
			if err != nil {
				return err
			}
		}

		if inspection.CopiedVariableNamesAll == nil {
			inspection.CopiedVariableNamesAll = []string{}
		}
		if inspection.CopiedVariableNamesBad == nil {
			inspection.CopiedVariableNamesBad = []string{}
		}
		if inspection.NewVariableNamesAll == nil {
			inspection.NewVariableNamesAll = []string{}
		}
		if inspection.NewVariableNamesBad == nil {
			inspection.NewVariableNamesBad = []string{}
		}
		if inspection.EditedVariableNamesAll == nil {
			inspection.EditedVariableNamesAll = []string{}
		}
		if inspection.EditedVariableNamesBad == nil {
			inspection.EditedVariableNamesBad = []string{}
		}

		if inspection.DocumentationState == "" {
			inspection.DocumentationState = "None/Incorrect/Correct"
		}

		if inspection.TaskCompletion == "" {
			if !submission.NotSubmitted {
				inspection.TaskCompletion = "Not submitted/Does not compile/Invalid solution/Completed"
			} else {
				inspection.TaskCompletion = "Not submitted"
			}
		}

		return structToFile(inspection, path.Join(submissionPath, ".manual_inspection.json"))
	})
}

func sessionInspection() error {
	return forEachSession(func(sessionPath string, session Session) error {
		inspection := SessionInspection{}

		if fileExists(path.Join(sessionPath, ".manual_inspection.json")) {
			data, err := ioutil.ReadFile(path.Join(sessionPath, ".manual_inspection.json"))
			if err != nil {
				return err
			}

			err = json.Unmarshal(data, &inspection)
			if err != nil {
				return err
			}
		}

		return structToFile(inspection, path.Join(sessionPath, ".manual_inspection.json"))
	})
}

var EmptyLineRE = regexp.MustCompile(`(?m)^[><|][\s}]*$`)
var AllLinesRE = regexp.MustCompile(`(?m)^[><|].*$`)

func sum() error {
	var table []TableRow

	var manualGroupings Groupings

	data, err := ioutil.ReadFile(path.Join(rootDir, ".manual_groupings.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &manualGroupings)
	if err != nil {
		return err
	}

	manualRules := make(map[string]Rule)

	data, err = ioutil.ReadFile(path.Join(rootDir, ".manual_rules.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &manualRules)
	if err != nil {
		return err
	}

	/* print latex list of rules
	san := func(s string) string {
		s = strings.ReplaceAll(s, "_", "\\_")
		s = strings.ReplaceAll(s, "$", "\\$")
		s = strings.ReplaceAll(s, "%", "\\%")
		s = strings.ReplaceAll(s, "^", "\\^{}")
		s = strings.ReplaceAll(s, "\"", "''")
		return s
	}

	for key, rule := range manualRules {
		if rule.Ignored {
			fmt.Printf("\\sonarruleignored{%s}{%s}{%s}\n", san(key), san(rule.ExampleMessage), san(rule.IgnoredComment))
		} else {
			fmt.Printf("\\sonarrule{%s}{%s}{%s}\n", san(key), san(rule.ExampleMessage), san(rule.IgnoredComment))
		}
	}
	*/

	err = forEachSubmission(func(submissionPath string, session Session, submission Submission) error {
		sessionPath, submissionName := path.Split(submissionPath)
		sessionName := path.Base(sessionPath)

		var sessionInspection SessionInspection

		data, err = ioutil.ReadFile(path.Join(sessionPath, ".manual_inspection.json"))
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &sessionInspection)
		if err != nil {
			return err
		}

		var submissionInspection SubmissionInspection

		data, err = ioutil.ReadFile(path.Join(submissionPath, ".manual_inspection.json"))
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &submissionInspection)
		if err != nil {
			return err
		}

		var issues []Issue

		if fileExists(path.Join(submissionPath, ".issues.json")) {
			data, err = ioutil.ReadFile(path.Join(submissionPath, ".issues.json"))
			if err != nil {
				return err
			}

			err = json.Unmarshal(data, &issues)
			if err != nil {
				return err
			}
		}

		issuesMajor := 0
		issuesInfo := 0
		issuesMinor := 0
		issuesCritical := 0
		skippedExisting := false
		if submissionName == "tickets" && submission.HighDebtVersion == true {
			skippedExisting = true
		}
		for _, issue := range issues {
			rule := manualRules[issue.Rule]

			if !rule.Ignored {
				if issue.Rule == "fb-contrib:OCP_OVERLY_CONCRETE_PARAMETER" &&
					issue.File == "Interval.java" &&
					submissionName == "booking" &&
					!skippedExisting {
					skippedExisting = true
					continue
				} else if issue.Rule == "java:S1172" &&
					issue.File == "TicketType.java" &&
					!skippedExisting {
					skippedExisting = true
					continue
				}
				switch issue.Severity {
				case "MAJOR":
					issuesMajor++
				case "INFO":
					issuesInfo++
				case "MINOR":
					issuesMinor++
				case "CRITICAL":
					issuesCritical++
				default:
					panic(issue.Severity)
				}
			}
		}

		modifiedLines := 0
		if fileExists(path.Join(submissionPath, ".diff")) {
			data, err := ioutil.ReadFile(path.Join(submissionPath, ".diff"))
			if err != nil {
				return err
			}

			modifiedLines = len(AllLinesRE.FindAll(data, -1)) - len(EmptyLineRE.FindAll(data, -1))
		}

		educationField := manualGroupings.EducationFields[session.BackgroundAnswers.EducationField]
		if educationField == "TODO" || educationField == "" {
			fmt.Printf("sum, skipping, no educationField mapping %s (%s) \n", session.BackgroundAnswers.EducationField, submissionPath)
			return nil
		}

		workDomain := manualGroupings.WorkDomains[session.BackgroundAnswers.WorkDomain]
		if workDomain == "TODO" || workDomain == "" {
			fmt.Printf("sum, skipping, no workdomain mapping %s (%s) \n", session.BackgroundAnswers.WorkDomain, submissionPath)
			return nil
		}

		group := manualGroupings.SignupCodes[session.SignupCode]
		if group == "TODO" || group == "" {
			fmt.Printf("sum, skipping, no signup code mapping %s (%s) \n", session.SignupCode, submissionPath)
			return nil
		}

		if !sessionInspection.Done {
			fmt.Printf("sum, skipping, no session inspection(%s) \n", submissionPath)
			return nil
		}

		if !submissionInspection.Done {
			fmt.Printf("sum, skipping, no submission inspection(%s) \n", submissionPath)
			return nil
		}

		time := int(submission.Time.Seconds())
		if time < 0 {
			time = 0
		}

		equalsState := "Not implemented"
		if submissionInspection.HasEquals {
			equalsState = "Duplicated"
			if submissionInspection.ReusedLogicEquals {
				equalsState = "Good"
			}
		}

		hashCodeState := "Not implemented"
		if submissionInspection.HasHashCode {
			hashCodeState = "Duplicated"
			if submissionInspection.ReusedLogicHashcode {
				hashCodeState = "Good"
			}
		}

		fmt.Printf("sum (%s)\n", submissionPath)

		tr := TableRow{
			Session:                    sessionName,
			Group:                      group,
			EducationLevel:             session.BackgroundAnswers.EducationLevel,
			EducationField:             educationField,
			ProgrammingExperience:      sessionInspection.WorkExperience,
			JavaExperience:             sessionInspection.WorkExperienceJava,
			WorkDomain:                 workDomain,
			CompanyUsesCodeReviews:     session.BackgroundAnswers.CompanyUsesCodeReviews,
			CompanyUsesPairProgramming: session.BackgroundAnswers.CompanyUsesPairProgramming,
			CompanyTracksTechnicalDebt: session.BackgroundAnswers.CompanyTracksTechnicalDebt,
			CompanyHasCodingStandards:  session.BackgroundAnswers.CompanyHasCodingStandards,
			Submission:                 submissionName,
			Order:                      submission.Order,
			HighDebtVersion:            submission.HighDebtVersion,
			QualityPreTask:             submission.ReflectionAnswers.ScenarioQuality,
			QualityPostTask:            submission.ReflectionAnswers.SubmissionQuality,
			Time:                       time,
			SonarqubeIssuesMajor:       issuesMajor,
			SonarqubeIssuesMinor:       issuesMinor,
			SonarqubeIssuesInfo:        issuesInfo,
			SonarqubeIssuesCritical:    issuesCritical,
			ModifiedLines:              modifiedLines,
			TaskCompletion:             submissionInspection.TaskCompletion,
			ReusedLogicConstructor:     submissionInspection.ReusedLogicConstructor,
			ReusedLogicValidation:      submissionInspection.ReusedLogicValidation,
			CopiedVariableNamesAll:     len(submissionInspection.CopiedVariableNamesAll),
			CopiedVariableNamesGood:    len(submissionInspection.CopiedVariableNamesAll) - len(submissionInspection.CopiedVariableNamesBad),
			NewVariableNamesAll:        len(submissionInspection.NewVariableNamesAll),
			NewVariableNamesGood:       len(submissionInspection.NewVariableNamesAll) - len(submissionInspection.NewVariableNamesBad),
			EditedVariableNamesAll:     len(submissionInspection.EditedVariableNamesAll),
			EditedVariableNamesGood:    len(submissionInspection.EditedVariableNamesAll) - len(submissionInspection.EditedVariableNamesBad),
			EqualsState:                equalsState,
			HashCodeState:              hashCodeState,
			DocumentationState:         submissionInspection.DocumentationState,
			LargeStructureChange:       submissionInspection.LargeStructureChange,
		}

		table = append(table, tr)
		return nil
	})
	if err != nil {
		return err
	}

	var tmp []TableRow

	for _, tr := range table {
		if tr.Order == 1 {
			var previous TableRow

			for _, otr := range table {
				if otr.Session == tr.Session && otr.Order == 0 {
					previous = otr
					break
				}
			}

			if previous.TaskCompletion == "Not submitted" && tr.TaskCompletion == "Not submitted" {
				continue
			}
		}

		tmp = append(tmp, tr)
	}

	table = tmp

	encoder := struct2csv.New()

	csvTable, err := encoder.Marshal(table)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(rootDir, "data.csv"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	err = writer.WriteAll(csvTable)
	if err != nil {
		return err
	}

	writer.Flush()

	return nil
}

func main() {
	stages := []func() error{
		download,
		diff,
		compile,
		execution,
		scan,
		issues,
		issuesRules,
		groups,
		submissionInspection,
		sessionInspection,
		sum,
	}

	for _, stage := range stages {
		err := stage()
		if err != nil {
			panic(err)
		}
	}
}

func debtString(hasHighDebt bool) string {
	if hasHighDebt {
		return "high"
	} else {
		return "low"
	}
}
