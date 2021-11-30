package task

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/config"
)

// InitTargetPath function to init the target path for the generated PDF files
func InitTargetPath(cmdConfig *config.CmdConfig) error {
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
			return fmt.Errorf("deleting the target path URL folder: %s", err)
		}
	}
	if err := os.MkdirAll(cmdConfig.TargetPathURL, os.ModePerm); err != nil {
		return fmt.Errorf("creating the target path URL folder: %s", err)
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
