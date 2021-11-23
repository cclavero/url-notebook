package cmd

import (
	"fmt"
	"os"
)

type RuntimeInt interface {
	SetError(err error)
	Exit()
}

type Runtime struct {
	err        error
	exitStatus int
}

func NewRuntime() *Runtime {
	return &Runtime{
		exitStatus: 0,
	}
}

func (runtime *Runtime) SetError(err error) {
	runtime.err = err
	runtime.exitStatus = 1
}

func (runtime *Runtime) Exit() {
	if runtime.err != nil { // Has error
		fmt.Fprintf(os.Stderr, "Error: %s\n", runtime.err.Error())
	}
	os.Exit(runtime.exitStatus)
}
