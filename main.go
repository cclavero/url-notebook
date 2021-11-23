package main

import (
	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/cclavero/ws-pdf-publish/config"
)

func main() {
	runtime := config.NewRuntime()

	rootCmd := cmd.NewRootCmd(runtime)
	if err := rootCmd.Execute(); err != nil {
		runtime.SetError(err)
	}

	runtime.Exit()
}
