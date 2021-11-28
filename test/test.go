package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/onsi/ginkgo"
)

const (
	TestBasePath = "../build/test"
)

// Global

func GetAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		ginkgo.Fail(err.Error())
	}
	return absPath
}

func RemoveAbsPath(path string) {
	if err := os.RemoveAll(path); err != nil {
		ginkgo.Fail(err.Error())
	}
}

func CopyFileToAbsPath(srcFile string, path string, trgFileName string) {
	trgFilePath := filepath.Join(path, trgFileName)
	bytesRead, err := ioutil.ReadFile(srcFile)
	if err != nil {
		ginkgo.Fail(err.Error())
	}
	if err = ioutil.WriteFile(trgFilePath, bytesRead, 0644); err != nil {
		ginkgo.Fail(err.Error())
	}
}

func ExecSysCommand(cmdStr string) string {
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		ginkgo.Fail(err.Error())
	}
	return string(stdout)
}

// TestCmdCtx

type TestCmdCtx struct {
	stdout    *os.File
	stderr    *os.File
	outReader *os.File
	outWriter *os.File
	errReader *os.File
	errWriter *os.File
}

func NewTestCmdCtx() *TestCmdCtx {
	return &TestCmdCtx{}
}

func (testCmdCtx *TestCmdCtx) OpenOutCapture() {
	var err error

	testCmdCtx.stdout = os.Stdout
	testCmdCtx.stderr = os.Stderr

	if testCmdCtx.outReader, testCmdCtx.outWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}
	if testCmdCtx.errReader, testCmdCtx.errWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	os.Stdout = testCmdCtx.outWriter
	os.Stderr = testCmdCtx.errWriter
}

func (testCmdCtx *TestCmdCtx) CloseOutCapture(logResults bool, limit int) (result string, errResult string) {
	var out []byte
	var outError []byte
	var err error

	testCmdCtx.outWriter.Close()
	if out, err = ioutil.ReadAll(testCmdCtx.outReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	testCmdCtx.errWriter.Close()
	if outError, err = ioutil.ReadAll(testCmdCtx.errReader); err != nil {
		ginkgo.Fail(fmt.Sprintf("unexpected error: %s", err))
	}

	os.Stdout = testCmdCtx.stdout
	os.Stderr = testCmdCtx.stderr

	result = string(out)
	errResult = string(outError)

	if logResults {
		testCmdCtx.logResult(result, limit)
		testCmdCtx.logErrorResult(errResult)
	}

	return result, errResult
}

func (testCmdCtx *TestCmdCtx) logResult(message string, limit int) {
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

func (testCmdCtx *TestCmdCtx) logErrorResult(message string) {
	if message == "" {
		log.Print("Error: No error\n\n")
	} else {
		log.Print(fmt.Sprintf("Error:\n\n%s\n\n", message))
	}
}
