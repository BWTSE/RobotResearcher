package main

import (
	"fmt"
	"os"

	"sidus.io/robotresearcher/internal/code"
	"sidus.io/robotresearcher/internal/router"
)

func main() {
	codeService, err := code.NewService()
	if err != nil {
		fmt.Println(fmt.Errorf("while creating code service: %w", err))
		os.Exit(1)
	}

	r, err := router.NewRouter(codeService)
	if err != nil {
		fmt.Println(fmt.Errorf("while creating router: %w", err))
		os.Exit(1)
	}

	err = r.Serve(true, "8080")
	if err != nil {
		fmt.Println(fmt.Errorf("while serving router: %w", err))
		os.Exit(1)
	}
}
