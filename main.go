package main

import (
	"fmt"

	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd *cobra.Command
	var err error

	if rootCmd, err = cmd.NewRootCmd(); err != nil {

		// TEMPORAL
		fmt.Printf("-------------------->ERROR:%s", err.Error())

	}
	if err := rootCmd.Execute(); err != nil {

		// TEMPORAL
		fmt.Printf("-------------------->ERROR:%s", err.Error())

		//runtime.SetError(err)
	}

	//runtime.Exit()
}
