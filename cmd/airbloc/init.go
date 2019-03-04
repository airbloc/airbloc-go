package main

import (
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize Airbloc node",
		Run:   runInit,
	}
)

func runInit(cmd *cobra.Command, args []string) {
}
