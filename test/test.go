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
	//TestBasePath = "../build/test"
)

type OutCapture struct {
	Stdout    *os.File
	Stderr    *os.File
	OutReader *os.File
	OutWriter *os.File
	ErrReader *os.File
	ErrWriter *os.File
}

func NewOutCapture() *OutCapture {
	var err error
	outCap := &OutCapture{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if outCap.OutReader, outCap.OutWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	if outCap.ErrReader, outCap.ErrWriter, err = os.Pipe(); err != nil {
		ginkgo.Fail(fmt.Sprintf("Unexpected error: %s", err))
	}
	os.Stdout = outCap.OutWriter
	os.Stderr = outCap.ErrWriter
	return outCap
}

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
