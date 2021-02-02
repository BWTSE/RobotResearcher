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
	idProducer      chan int
	dummyIdProducer chan int
	rand            *rand.Rand
	sequences       [][]database.Scenario
}

func NewService(previousParticipants int) (*Service, error) {
	idProducer := make(chan int)

	go func() {
		for i := previousParticipants; true; i++ {
			idProducer <- i
		}
	}()

	dummyIdProducer := make(chan int)

	go func() {
		for i := 0; true; i++ {
			dummyIdProducer <- i
		}
	}()

	files, err := ioutil.ReadDir("Scenarios")
	if err != nil {
		return nil, err
	}

	var artifacts []artifact

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(path.Base(file.Name()), ".") {
			a, err := newArtifact(path.Join("Scenarios", file.Name()))
			if err != nil {
				return nil, err
			}

			artifacts = append(artifacts, a)
		}
	}

	if len(artifacts) != 2 {
		return nil, fmt.Errorf("%d artifacts instead of 3", len(artifacts))
	}

	a1 := artifacts[0]
	a2 := artifacts[1]

	return &Service{
		idProducer:      idProducer,
		dummyIdProducer: dummyIdProducer,
		rand:            rand.New(rand.NewSource(time.Now().Unix())),
		sequences: [][]database.Scenario{
			{
				a1.lowDebtScenario(),
				a2.highDebtScenario(),
			},
			{
				a1.highDebtScenario(),
				a2.lowDebtScenario(),
			},
		},
	}, nil
}

func (s *Service) GetSequence() []database.Scenario {
	scenarios := s.sequences[<-s.idProducer%len(s.sequences)]

	s.rand.Shuffle(len(scenarios), func(i, j int) {
		scenarios[i], scenarios[j] = scenarios[j], scenarios[i]
	})

	return scenarios
}

func (s *Service) GetDummySequence() []database.Scenario {
	scenarios := s.sequences[<-s.dummyIdProducer%len(s.sequences)]

	s.rand.Shuffle(len(scenarios), func(i, j int) {
		scenarios[i], scenarios[j] = scenarios[j], scenarios[i]
	})

	return scenarios
}
