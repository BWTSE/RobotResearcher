package main

import (
	"fmt"
	"os"

	"sidus.io/robotresearcher/internal/database"

	"sidus.io/robotresearcher/internal/code"
	"sidus.io/robotresearcher/internal/router"
)

func main() {
	codeService, err := code.NewService()
	if err != nil {
		fmt.Println(fmt.Errorf("while creating code service: %w", err))
		os.Exit(1)
	}

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println(fmt.Errorf("while creating code service: %w", err))
		os.Exit(1)
	}
	defer db.Close()

	r, err := router.NewRouter(codeService, db)
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
