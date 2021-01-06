package main

import (
	"fmt"
	"os"

	"sidus.io/robotresearcher/internal/code"
	"sidus.io/robotresearcher/internal/database"
	"sidus.io/robotresearcher/internal/router"
	"sidus.io/robotresearcher/internal/scenario"
)

func main() {
	codeService, err := code.NewService()
	if err != nil {
		fmt.Println(fmt.Errorf("while creating code service: %w", err))
		os.Exit(1)
	}

	db, err := database.NewDatabase(os.Getenv("DB_URI"))
	if err != nil {
		fmt.Println(fmt.Errorf("while creating code service: %w", err))
		os.Exit(1)
	}
	defer db.Close()

	previousParticipants, err := db.CountSessions()
	if err != nil {
		fmt.Println(fmt.Errorf("while counting previous sessions: %w", err))
		os.Exit(1)
	}

	scenarioService, err := scenario.NewService(previousParticipants)
	if err != nil {
		fmt.Println(fmt.Errorf("while creating scenario servive: %w", err))
		os.Exit(1)
	}

	r, err := router.NewRouter(codeService, scenarioService, db, "/tmp")
	if err != nil {
		fmt.Println(fmt.Errorf("while creating router: %w", err))
		os.Exit(1)
	}

	err = r.Serve(true, "8000")
	if err != nil {
		fmt.Println(fmt.Errorf("while serving router: %w", err))
		os.Exit(1)
	}
}
