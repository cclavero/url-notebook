package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/onsi/ginkgo"
)

type TestCtx struct {
	stdout    *os.File
	stderr    *os.File
	outReader *os.File
	outWriter *os.File
	errReader *os.File
	errWriter *os.File
}

func NewTestCtx() *TestCtx {
	return &TestCtx{}
}

func (testCtx *TestCtx) OpenOutCapture() {
	var err error

	testCtx.stdout = os.Stdout
	testCtx.stderr = os.Stderr

	if testCtx.outReader, testCtx.outWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}
	if testCtx.errReader, testCtx.errWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	os.Stdout = testCtx.outWriter
	os.Stderr = testCtx.errWriter
}

func (testCtx *TestCtx) CloseOutCapture(logResults bool, limit int) (result string, errResult string) {
	var out []byte
	var outError []byte
	var err error

	testCtx.outWriter.Close()
	if out, err = ioutil.ReadAll(testCtx.outReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	testCtx.errWriter.Close()
	if outError, err = ioutil.ReadAll(testCtx.errReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	os.Stdout = testCtx.stdout
	os.Stderr = testCtx.stderr

	result = string(out)
	errResult = string(outError)

	if logResults {
		testCtx.logResult(result, limit)
		testCtx.logErrorResult(errResult)
	}

	return result, errResult
}

func (testCtx *TestCtx) logResult(message string, limit int) {
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

func (testCtx *TestCtx) logErrorResult(message string) {
	if message == "" {
		log.Print("Error: No error\n\n")
	} else {
		log.Print(fmt.Sprintf("Error:\n\n%s\n\n", message))
	}
}
