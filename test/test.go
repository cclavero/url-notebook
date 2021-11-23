package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/onsi/ginkgo"
)

const (
	TestCmdName  = "ws-pdf-publish"
	TestBasePath = "../build/test"
)

type TestRuntime struct {
	// Exit
	err        error
	exitStatus int
	// Out capture
	stdout    *os.File
	stderr    *os.File
	outReader *os.File
	outWriter *os.File
	errReader *os.File
	errWriter *os.File
}

func NewTestRuntime() *TestRuntime {
	return &TestRuntime{
		exitStatus: 0,
	}
}
func (testRun *TestRuntime) SetError(err error) {
	testRun.err = err
	testRun.exitStatus = 1
}

func (testRun *TestRuntime) Exit() {
	if testRun.err != nil { // Has error
		fmt.Fprintf(os.Stderr, "Error: %s\n", testRun.err.Error())
	}
	if testRun.exitStatus != 0 {
		fmt.Printf("exit status %d", testRun.exitStatus)

		// TEMPORAL
		//ginkgo.Fail("hola")

	}
}

func (testRun *TestRuntime) OpenOutCapture() {
	var err error

	testRun.stdout = os.Stdout
	testRun.stderr = os.Stderr

	if testRun.outReader, testRun.outWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}
	if testRun.errReader, testRun.errWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	os.Stdout = testRun.outWriter
	os.Stderr = testRun.errWriter
}

func (testRun *TestRuntime) CloseOutCapture(logResults bool, limit int) (result string, errResult string) {
	var out []byte
	var outError []byte
	var err error

	testRun.outWriter.Close()
	if out, err = ioutil.ReadAll(testRun.outReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	testRun.errWriter.Close()
	if outError, err = ioutil.ReadAll(testRun.errReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	os.Stdout = testRun.stdout
	os.Stderr = testRun.stderr

	result = string(out)
	errResult = string(outError)

	if logResults {
		testRun.logResult(result, limit)
		testRun.logErrorResult(errResult)
	}

	return result, errResult
}

func (testRun *TestRuntime) logResult(message string, limit int) {
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

func (testRun *TestRuntime) logErrorResult(message string) {
	if message == "" {
		log.Print("Error: No error\n\n")
	} else {
		log.Print(fmt.Sprintf("Error:\n\n%s\n\n", message))
	}
}
