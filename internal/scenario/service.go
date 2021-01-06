package scenario

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"strings"
	"time"

	"sidus.io/robotresearcher/internal/database"
)

type Service struct {
	a1         artifact
	a2         artifact
	a3         artifact
	idProducer chan int
	rand       *rand.Rand
}

func NewService(previousParticipants int) (*Service, error) {
	idProducer := make(chan int)

	go func() {
		for i := previousParticipants; true; i++ {
			idProducer <- i
		}
	}()

	files, err := ioutil.ReadDir("artifacts")
	if err != nil {
		return nil, err
	}

	var artifacts []artifact

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(path.Base(file.Name()), ".") {
			a, err := newArtifact(path.Join("artifacts", file.Name()))
			if err != nil {
				return nil, err
			}

			artifacts = append(artifacts, a)
		}
	}

	if len(artifacts) != 3 {
		return nil, fmt.Errorf("%d artifacts instead of 3", len(artifacts))
	}

	return &Service{
		a1:         artifacts[0],
		a2:         artifacts[1],
		a3:         artifacts[2],
		idProducer: idProducer,
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
	}, nil
}

func (s *Service) GetSequence() []database.Scenario {
	sequences := [][]database.Scenario{
		{
			s.a1.lowDebtScenario(),
			s.a2.highDebtScenario(),
			s.a3.lowDebtScenario(),
		},
		{
			s.a1.highDebtScenario(),
			s.a2.lowDebtScenario(),
			s.a3.highDebtScenario(),
		},
		{
			s.a1.highDebtScenario(),
			s.a2.highDebtScenario(),
			s.a3.lowDebtScenario(),
		},
		{
			s.a1.lowDebtScenario(),
			s.a2.lowDebtScenario(),
			s.a3.highDebtScenario(),
		},
		{
			s.a1.lowDebtScenario(),
			s.a2.highDebtScenario(),
			s.a3.highDebtScenario(),
		},
		{
			s.a1.highDebtScenario(),
			s.a2.lowDebtScenario(),
			s.a3.lowDebtScenario(),
		},
	}

	scenarios := sequences[<-s.idProducer%len(sequences)]

	s.rand.Shuffle(len(scenarios), func(i, j int) {
		scenarios[i], scenarios[j] = scenarios[j], scenarios[i]
	})

	return scenarios
}
