package main

import (
	"fmt"
	"os"

	"github.com/cclavero/ws-pdf-publish/cmd"
)

func main() {

	// TEMPORAL
	runtimeEnv := NewRuntimeEnv()

	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		exitWithError(err)
	}

	// TEMPORAL
	//runtimeSer.Exit()

}

// TEMPORAL
func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// TEMPORAL
/*
// RuntimeServiceInt interface for the Runtime service
type RuntimeServiceInt interface {
	SetError(err error)
	GetError() string
	Print(message string)
	GetExitStatus() int
	SetExitCmd(cmd func(exitStatus int))
	Exit()
}

// RuntimeService type for the Runtime service
type RuntimeService struct {
	runError   string
	exitStatus int
	exitCmd    func(exitStatus int)
}

// NewRuntimeService constructor to get a new Runtime service instance
func NewRuntimeService() RuntimeServiceInt {
	return &RuntimeService{
		runError:   "",
		exitStatus: 0,
		exitCmd: func(exitStatus int) {
			os.Exit(exitStatus)
		},
	}
}

/*
// SetError method to set an execution error
func (runtimeSer *RuntimeService) SetError(err error) {
	runtimeSer.runError = err.Error()
	runtimeSer.exitStatus = 1
}

// GetError method to get an execution error
func (runtimeSer *RuntimeService) GetError() string {
	return runtimeSer.runError
}

// Print method to print some message to Stdout
func (runtimeSer *RuntimeService) Print(message string) {
	// Important: Print all buffer
	fmt.Fprintf(os.Stdout, "%s", message)
}

// GetExitStatus method to get the exit status
func (runtimeSer *RuntimeService) GetExitStatus() int {
	return runtimeSer.exitStatus
}

// SetExitCmd method to set the final exit command
func (runtimeSer *RuntimeService) SetExitCmd(cmd func(exitStatus int)) {
	runtimeSer.exitCmd = cmd
}

// Exit method to execit the command execution
func (runtimeSer *RuntimeService) Exit() {
	if runtimeSer.exitStatus != 0 { // Has error
		fmt.Fprintf(os.Stderr, "Error: %s\n", runtimeSer.runError)
	}
	runtimeSer.exitCmd(runtimeSer.exitStatus)
}
*/
