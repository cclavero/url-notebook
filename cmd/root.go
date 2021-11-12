package cmd

import (
	"fmt"
	"os"

	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/spf13/cobra"
)

var (
	Version = "devel"
	rootCmd = &cobra.Command{
		Use:   config.WSPDFpublishCmd,
		Short: "WebSite PDF Publish simple command line tool to publish HTML pages to PDF.",
		Long: `WebSite PDF Publish is a simple command line tool to publish some set of pages from a WebSite to PDF, using a 'ws-pub-pdf.yaml' configuration file.
Internally, uses the wkhtmltopdf utility.`,
		Run: execRoot,
	}
)

func init() {
	rootCmd.Flags().StringP(config.PublishFileFlag, "", "", "set the 'ws-pub-pdf.yaml' publish file, including absolute or relative path.")
	rootCmd.MarkFlagRequired(config.PublishFileFlag)
	rootCmd.Flags().StringP(config.TargetPathFlag, "", "", "set the target path for publishing partial and final PDF files.")
	rootCmd.MarkFlagRequired(config.TargetPathFlag)
	rootCmd.Version = Version
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func execRoot(cmd *cobra.Command, args []string) {
	var cmdConfig *config.CmdConfig
	var err error

	// Get the config
	if cmdConfig, err = config.GetCmdConfig(cmd); err != nil {
		exitWithError(fmt.Errorf("Error: Getting cmd config: %s\n", err))
	}

	// TEMPORAL
	fmt.Printf("--------------------->%+v\n\n", cmdConfig)

	// Info
	fmt.Println("Starting ...")

	//fmt.Printf("Command config: %s", config.GetCmdConfigInfo(cmdConfig))

	/*

		// Build 'wkhtmltopdf' docker image
		fmt.Println("Building 'wkhtmltopdf' docker image ...")
		if err := pdf.BuildWkhtmltoPDFDocker(); err != nil {
			panic(fmt.Sprintf("Error: Building wkhtmltopdf docker image: %s", err))
		}

		// Init target path
		targetFile := filepath.Join(cmdConfig.TargetPath, cmdConfig.PublishData.File)
	*/

	fmt.Println("Init target path ...")

	// TEMPORAL
	targetFile := "final"

	/*
		if err := task.InitTargetPath(cmdConfig); err != nil {
			panic(fmt.Sprintf("Error: Initializing target path: %s\n", err))
		}
	*/

	// Publish all URLs
	fmt.Println("Publish all URLs:")

	/*
		for index, pub := range cmdConfig.PublishData.URLList {
			if err := pdf.PublishURLAsPDF(cmdConfig, index+1, pub); err != nil {
				panic(fmt.Sprintf("Error: Publishing URL as PDF: %s\n", err))
			}
		}
	*/

	// Merge all PDF files
	fmt.Println("Merge all PDF files ...")

	/*
		if err := pdf.MergePDFFiles(cmdConfig); err != nil {
			panic(fmt.Sprintf("Error: Merging all PDF files: %s\n", err))
		}
	*/

	// End
	fmt.Printf("\nDone, full PDF file generated at: %s\n\n", targetFile)
}
