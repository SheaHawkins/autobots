/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/SheaHawkins/autobots/internal/cmd"
	"github.com/SheaHawkins/autobots/internal/cmd/shared"
)

func main() {
	deps := shared.Dependencies{}
	rootCmd := cmd.NewRoot(deps)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
