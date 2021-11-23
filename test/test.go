package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/onsi/ginkgo"
)

const (
	TestCmdName = "ws-pdf-publish"

	// TEMPORAL
	//TestBasePath = "../build/test"
)

type TestRuntime struct {
	// Exit
	err        error
	exitStatus int

	// TEMPORAL
	// Out capture
	Stdout    *os.File
	Stderr    *os.File
	OutReader *os.File
	OutWriter *os.File
	ErrReader *os.File
	ErrWriter *os.File
}

func NewTestRuntime() *TestRuntime {
	return &TestRuntime{
		exitStatus: 0,
	}
}
func (testRuntime *TestRuntime) SetError(err error) {
	testRuntime.err = err
	testRuntime.exitStatus = 1
}

func (testRuntime *TestRuntime) Exit() {
	if testRuntime.err != nil { // Has error
		fmt.Fprintf(os.Stderr, "Error: %s\n", testRuntime.err.Error())
	}
	if testRuntime.exitStatus != 0 {
		fmt.Printf("exit status %d", testRuntime.exitStatus)
	}
}

// TEMPORAL
func LogResult(message string, limit int) {
	if message == "" {
		log.Print("Result: Empty string\n\n")
	} else {
		if limit != 0 && len(message) > limit {
			log.Print(fmt.Sprintf("Result:\n\n%s[TRUNCATED]\n\n", message[0:limit]))
		} else {
			log.Print(fmt.Sprintf("Result:\n\n%s\n\n", message))
		}
	}
}

func LogError(message string) {
	if message == "" {
		log.Print("Error: No error\n\n")
	} else {
		log.Print(fmt.Sprintf("Error:\n\n%s\n\n", message))
	}
}

/*
type OutCapture struct {
	Stdout    *os.File
	Stderr    *os.File
	OutReader *os.File
	OutWriter *os.File
	ErrReader *os.File
	ErrWriter *os.File
}
*/

func (testRuntime *TestRuntime) StartOutCapture() {
	var err error

	// TEMPORAL
	/*
		outCap := &OutCapture{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
	*/
	testRuntime.Stdout = os.Stdout
	testRuntime.Stderr = os.Stderr

	if testRuntime.OutReader, testRuntime.OutWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	if testRuntime.ErrReader, testRuntime.ErrWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	os.Stdout = testRuntime.OutWriter
	os.Stderr = testRuntime.ErrWriter
}

func (testRuntime *TestRuntime) CloseOutCapture() (string, string) {
	var out []byte
	var outError []byte
	var err error
	testRuntime.OutWriter.Close()
	if out, err = ioutil.ReadAll(testRuntime.OutReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	testRuntime.ErrWriter.Close()
	if outError, err = ioutil.ReadAll(testRuntime.ErrReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	os.Stdout = testRuntime.Stdout
	os.Stderr = testRuntime.Stderr
	return string(out), string(outError)
}

/*
func (outCap OutCapture) Close() (string, string) {
	var out []byte
	var outError []byte
	var err error
	outCap.OutWriter.Close()
	if out, err = ioutil.ReadAll(outCap.OutReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	outCap.ErrWriter.Close()
	if outError, err = ioutil.ReadAll(outCap.ErrReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	os.Stdout = outCap.Stdout
	os.Stderr = outCap.Stderr
	return string(out), string(outError)
}
*/
