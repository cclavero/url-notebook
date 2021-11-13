package cmd

import (
	"fmt"
	"os"

	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/task"
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
	if err := rootCmd.MarkFlagRequired(config.PublishFileFlag); err != nil {
		exitWithError(fmt.Errorf("making flag required: %s\n", err))
	}
	rootCmd.Flags().StringP(config.TargetPathFlag, "", "", "set the target path for publishing partial and final PDF files.")
	if err := rootCmd.MarkFlagRequired(config.TargetPathFlag); err != nil {
		exitWithError(fmt.Errorf("making flag required: %s\n", err))
	}
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
		exitWithError(fmt.Errorf("getting cmd config: %s\n", err))
	}

	// Info
	fmt.Println("Starting ...")
	fmt.Printf("Config: %s", fmt.Sprintf(config.ConfigInfoStr, cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPath,
		cmdConfig.TargetPathURL, cmdConfig.TargetFile, cmdConfig.PublishData))

	// Build 'wkhtmltopdf' docker image
	fmt.Println("Checking 'wkhtmltopdf' docker image ...")
	if err = task.CheckWkhtmltoPDFDocker(); err != nil {
		exitWithError(fmt.Errorf("building wkhtmltopdf docker image: %s", err))
	}

	// Init target path
	fmt.Println("Init target path ...")
	if err = task.InitTargetPath(cmdConfig); err != nil {
		exitWithError(fmt.Errorf("initializing target path: %s", err))
	}

	// Publish all URLs
	fmt.Println("Publish all URLs:")
	for index, pub := range cmdConfig.PublishData.URLList {
		if err = task.PublishURLAsPDF(cmdConfig, index+1, pub); err != nil {
			exitWithError(fmt.Errorf("publishing URL as PDF: %s", err))
		}
	}

	// Merge all PDF files
	fmt.Println("Merge all PDF files ...")
	if err = task.MergePDFFiles(cmdConfig); err != nil {
		exitWithError(fmt.Errorf("merging all PDF files: %s", err))
	}

	// End
	fmt.Printf("\nDone, full PDF file generated at: %s\n\n", cmdConfig.TargetFile)
}
