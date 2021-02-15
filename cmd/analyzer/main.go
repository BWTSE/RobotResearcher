package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"sidus.io/robotresearcher/internal/database"
)

func main() {
	db, err := database.NewDatabase(os.Getenv("DB_URI"))
	if err != nil {
		fmt.Println(fmt.Errorf("while creating code service: %w", err))
		os.Exit(1)
	}
	defer db.Close()

	sessions, err := db.GetSessions("our-hero")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Total sessions found: %d\n", len(sessions))

	var startedSessions []database.Session

	for _, session := range sessions {
		if !session.Scenarios[0].SubmittedAt.IsZero() {
			startedSessions = append(startedSessions, session)
		}
	}

	fmt.Printf("After filtering out useless: %d\n", len(startedSessions))

	for _, session := range startedSessions {
		for _, scenario := range session.Scenarios {
			if scenario.SubmittedAt.IsZero() {
				continue
			}

			srcDir := path.Join("submissions", session.ID.Hex())
			dir := path.Join(srcDir, scenario.Name)

			_, err := ioutil.ReadDir(dir)
			if err == nil { // dir exists
				fmt.Printf("skipping %s \n", dir)
				continue
			}

			err = os.MkdirAll(dir, 0775)
			if err != nil {
				panic(err)
			}

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

				err = ioutil.WriteFile(path.Join(dir, name), code, 0644)
				if err != nil {
					fmt.Println(err)
					break
				}
			}

			for file := range scenario.Submitted {
				cmd := exec.Command(
					"javac",
					path.Join(scenario.Name, file),
				)
				cmd.Dir = srcDir

				out, err := cmd.CombinedOutput()
				if err != nil {
					_, ok := err.(*exec.ExitError)
					if ok {
						err := ioutil.WriteFile(path.Join(dir, "error.compiletime"), out, 0644)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
						break
					}
				}
			}

			cd, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			cmd := exec.CommandContext(
				cd,
				"java",
				"-Djava.security.manager",
				//			"-Djava.security.debug=all",
				//			"-Djava.security.policy==./../../untrusted.policy",
				"-Xmx15M",
				scenario.Name+".Main",
			)
			cmd.Dir = srcDir

			out, err := cmd.CombinedOutput()
			if err != nil {
				_, ok := err.(*exec.ExitError)
				if ok {
					err := ioutil.WriteFile(path.Join(dir, "error.runtime"), out, 0644)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
					break
				}
			} else {
				err = ioutil.WriteFile(path.Join(dir, "output"), out, 0644)
				if err != nil {
					fmt.Println(err)
				}
			}

			cmd = exec.CommandContext(
				cd,
				"sonar-scanner",
				fmt.Sprintf("-Dsonar.projectKey=%s-%s-%s", scenario.Name, debtString(scenario.HasHighDebt), session.ID.Hex()),
				"-Dsonar.scm.exclusions.disabled=true",
				"-Dsonar.java.binaries="+scenario.Name,
				"-Dsonar.sources="+scenario.Name,
			)
			cmd.Dir = srcDir

			out, err = cmd.CombinedOutput()
			if err != nil {
				_, ok := err.(*exec.ExitError)
				if ok {
					err := ioutil.WriteFile(path.Join(dir, "error.sonar"), out, 0644)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
					break
				}
			} else {
				err = ioutil.WriteFile(path.Join(dir, "output.sonar"), out, 0644)
				if err != nil {
					fmt.Println(err)
				}
			}
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
