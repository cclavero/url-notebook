package main

import (
	"fmt"

	"github.com/cclavero/url-notebook/cmd/config"
	"github.com/cclavero/url-notebook/cmd/pdf"
	"github.com/cclavero/url-notebook/cmd/task"
)

const (
	Version = "1.0"
)

func main() {
	var cmdConfig *config.CmdConfig

	// Info
	fmt.Println(fmt.Sprintf("PublishPDF: v%s", Version))

	// Get config
	var err error
	if cmdConfig, err = config.GetCmdConfig(); err != nil {
		panic(fmt.Sprintf("Error: Getting cmd config: %s\n", err))
	}

	// Info
	fmt.Printf("\n- Starting\n\n")
	fmt.Printf("- Command config: %+v\n", cmdConfig)

	// Build 'wkhtmltopdf' docker image
	fmt.Printf("\n> Building 'wkhtmltopdf' docker image ...\n")
	if err := pdf.BuildWkhtmltopdfDocker(); err != nil {
		panic(fmt.Sprintf("Error: Building wkhtmltopdf docker image: %s", err))
	}

	// Execute clean task
	fmt.Printf("\n> Execute clean task ...\n")
	if err := task.ExecCleanTask(cmdConfig); err != nil {
		panic(fmt.Sprintf("Error: Removing old PDF files: %s\n", err))
	}

	// Publish all URLs
	fmt.Printf("\n> Publish all URLs ...\n")
	for index, pub := range cmdConfig.PublishData.URLList {
		fmt.Printf("%d. Publishing %s to %s ...\n", index+1, pub.URL, pub.File)
		if err := pdf.PublishURLAsPDF(cmdConfig, pub); err != nil {
			panic(fmt.Sprintf("Error: Publishing URL as PDF: %s\n", err))
		}
	}

	// Merge all PDF files
	fmt.Printf("\n> Merge all PDF files ...\n")
	if err := pdf.MergePDFFiles(cmdConfig); err != nil {
		panic(fmt.Sprintf("Error: Merging all PDF files: %s\n", err))
	}

	// End
	fmt.Printf("\n- Done\n\n")
}
