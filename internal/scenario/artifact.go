package scenario

import (
	"encoding/base64"
	"io/ioutil"
	"path"
	"strings"

	"sidus.io/robotresearcher/internal/database"
)

type artifact struct {
	instructions    string
	lowDebtVersion  map[string]string
	highDebtVersion map[string]string
	name            string
}

func newArtifact(dir string) (artifact, error) {
	instructions, err := ioutil.ReadFile(path.Join(dir, "instructions.md"))
	if err != nil {
		return artifact{}, err
	}

	lowDebt, err := getFiles(path.Join(dir, "low_debt", path.Base(dir)))
	if err != nil {
		return artifact{}, err
	}

	highDebt, err := getFiles(path.Join(dir, "high_debt", path.Base(dir)))
	if err != nil {
		return artifact{}, err
	}

	return artifact{
		instructions:    string(instructions),
		lowDebtVersion:  lowDebt,
		highDebtVersion: highDebt,
		name:            path.Base(dir),
	}, nil
}

func getFiles(dir string) (map[string]string, error) {
	result := make(map[string]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".java") || strings.HasSuffix(file.Name(), ".txt") {
			data, err := ioutil.ReadFile(path.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}

			result[file.Name()] = base64.StdEncoding.EncodeToString(data)
		}
	}

	return result, nil
}

func (a artifact) lowDebtScenario() database.Scenario {
	return database.Scenario{
		HasHighDebt:  false,
		Instructions: a.instructions,
		Starting:     a.lowDebtVersion,
		Name:         a.name,
	}
}

func (a artifact) highDebtScenario() database.Scenario {
	return database.Scenario{
		HasHighDebt:  true,
		Instructions: a.instructions,
		Starting:     a.highDebtVersion,
		Name:         a.name,
	}
}
