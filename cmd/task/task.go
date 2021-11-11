package task

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cclavero/url-notebook/cmd/config"
)

func InitTargetPath(cmdConfig *config.CmdConfig) error {
	fmt.Printf("Check target path: %s\n", cmdConfig.TargetPath)

	// Target path
	if _, err := os.Stat(cmdConfig.TargetPath); os.IsNotExist(err) {
		if err := os.MkdirAll(cmdConfig.TargetPath, os.ModePerm); err != nil {
			return fmt.Errorf("creating the target path: %s", err)
		}
	}
	if err := removeFilesFromPath(cmdConfig.TargetPath, "*.pdf"); err != nil {
		return err
	}

	// Target path url
	if _, err := os.Stat(cmdConfig.TargetPathURL); !os.IsNotExist(err) {
		if err := removeFilesFromPath(cmdConfig.TargetPathURL, "*.pdf"); err != nil {
			return err
		}
		if err := os.Remove(cmdConfig.TargetPathURL); err != nil {
			return fmt.Errorf("deleting the target path URL: %s", err)
		}
	}
	if err := os.MkdirAll(cmdConfig.TargetPathURL, os.ModePerm); err != nil {
		return fmt.Errorf("creating the target url folder: %s", err)
	}

	return nil
}

func removeFilesFromPath(path string, extFilter string) error {
	var files []string
	var err error
	if files, err = filepath.Glob(filepath.Join(path, extFilter)); err != nil {
		return fmt.Errorf("error getting pdf files: %s", err)
	}
	for _, file := range files {
		if err = os.RemoveAll(file); err != nil {
			return fmt.Errorf("error deleting file: %s; %s", file, err)
		}
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
