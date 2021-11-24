package main

import (
	"fmt"
	"os"

	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd *cobra.Command
	var err error

	if rootCmd, err = cmd.NewRootCmd(); err != nil {
		exitWithError(err)
	}
	if err = rootCmd.Execute(); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
