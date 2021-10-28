package task

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cclavero/url-notebook/cmd/config"
)

const (
	urlFolder = "url"
)

func InitTargetPath(cmdConfig *config.CmdConfig) error {
	var files []string
	var err error
	if files, err = filepath.Glob(filepath.Join(cmdConfig.TargetPath, "*.pdf")); err != nil {
		return fmt.Errorf("error getting pdf files: %s", err)
	}
	for _, file := range files {
		if err = os.RemoveAll(file); err != nil {
			return fmt.Errorf("error deleting file: %s; %s", file, err)
		}
	}

	// TEMPORAL
	targetURLFolder := filepath.Join(cmdConfig.TargetPath, urlFolder)
	if err := os.MkdirAll(targetURLFolder, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Error: Creating the target url folder: %s", err))
	}

	return nil
}

func ExecSystemCommand(cmdStr string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("system command: '%s'; %s", cmdStr, err)
	}
	return string(stdout), nil
}
