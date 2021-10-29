package main

import (
	"fmt"
	"path/filepath"

	"github.com/cclavero/url-notebook/cmd/config"
	"github.com/cclavero/url-notebook/cmd/pdf"
	"github.com/cclavero/url-notebook/cmd/task"
)

var (
	version = "devel"
)

func main() {
	var cmdConfig *config.CmdConfig

	// Info
	fmt.Println(fmt.Sprintf("\n> URLNotebook: v%s", version))

	// Get config
	var err error
	if cmdConfig, err = config.GetCmdConfig(); err != nil {
		panic(fmt.Sprintf("Error: Getting cmd config: %s\n", err))
	}

	// Info
	fmt.Println("Starting ...")
	fmt.Printf("Command config: %s", config.GetCmdConfigInfo(cmdConfig))

	// Build 'wkhtmltopdf' docker image
	fmt.Println("Building 'wkhtmltopdf' docker image ...")
	if err := pdf.BuildWkhtmltopdfDocker(); err != nil {
		panic(fmt.Sprintf("Error: Building wkhtmltopdf docker image: %s", err))
	}

	// Init target path
	targetFile := filepath.Join(cmdConfig.TargetPath, cmdConfig.PublishData.File)

	fmt.Println("Init target path ...")
	if err := task.InitTargetPath(cmdConfig); err != nil {
		panic(fmt.Sprintf("Error: Initializing target path: %s\n", err))
	}

	// Publish all URLs
	fmt.Println("Publish all URLs:")
	for index, pub := range cmdConfig.PublishData.URLList {
		if err := pdf.PublishURLAsPDF(cmdConfig, index+1, pub); err != nil {
			panic(fmt.Sprintf("Error: Publishing URL as PDF: %s\n", err))
		}
	}

	/*
		// Merge all PDF files
		fmt.Println("Merge all PDF files ...")
		if err := pdf.MergePDFFiles(cmdConfig); err != nil {
			panic(fmt.Sprintf("Error: Merging all PDF files: %s\n", err))
		}
	*/

	// End
	fmt.Printf("\nDone, full PDF file generated at: %s\n\n", targetFile)
}
