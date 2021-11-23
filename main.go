package main

import (
	"github.com/cclavero/ws-pdf-publish/cmd"
)

func main() {
	runtime := cmd.NewRuntime()

	rootCmd := cmd.NewRootCmd(runtime)
	if err := rootCmd.Execute(); err != nil {
		runtime.SetError(err)
	}

	runtime.Exit()
}
