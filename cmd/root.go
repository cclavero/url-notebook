package cmd

import (
	"fmt"
	"time"

	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/task"
	"github.com/spf13/cobra"
)

var (
	Version = "devel"
	runtime RuntimeInt
)

func NewRootCmd(runtimeIn RuntimeInt) *cobra.Command {
	runtime = runtimeIn

	rootCmd := &cobra.Command{
		Use:   config.WSPDFpublishCmd,
		Short: "WebSite PDF Publish simple command line tool to publish HTML pages to PDF.",
		Long: `WebSite PDF Publish is a simple command line tool to publish some set of pages from a WebSite to PDF, using a 'ws-pub-pdf.yaml' configuration file.
Internally, uses the wkhtmltopdf utility.`,
		Run: execRoot,
	}

	rootCmd.Flags().StringP(config.PublishFileFlag, "", "", "set the 'ws-pub-pdf.yaml' publish file, including absolute or relative path.")
	if err := rootCmd.MarkFlagRequired(config.PublishFileFlag); err != nil {
		runtime.SetError(fmt.Errorf("making flag required: %s", err))
		runtime.Exit()
	}
	rootCmd.Flags().StringP(config.TargetPathFlag, "", "", "set the target path for publishing partial and final PDF files.")
	if err := rootCmd.MarkFlagRequired(config.TargetPathFlag); err != nil {
		runtime.SetError(fmt.Errorf("making flag required: %s", err))
		runtime.Exit()
	}
	rootCmd.Version = Version

	return rootCmd
}

func execRoot(cmd *cobra.Command, args []string) {
	var cmdConfig *config.CmdConfig
	var err error

	// Time
	start := time.Now()

	// Get the config
	if cmdConfig, err = config.GetCmdConfig(cmd); err != nil {
		runtime.SetError(fmt.Errorf("getting cmd config: %s", err))
		runtime.Exit()
	}

	// Info
	fmt.Println("Starting ...")
	fmt.Printf("Config: %s", fmt.Sprintf(config.ConfigInfoStr, cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPath,
		cmdConfig.TargetPathURL, cmdConfig.TargetFile, cmdConfig.PublishData))

	// Build 'wkhtmltopdf' docker image
	fmt.Println("Checking 'wkhtmltopdf' docker image ...")
	if err = task.CheckWkhtmltoPDFDocker(); err != nil {
		runtime.SetError(fmt.Errorf("building wkhtmltopdf docker image: %s", err))
		runtime.Exit()
	}

	// Init target path
	fmt.Println("Init target path ...")
	if err = task.InitTargetPath(cmdConfig); err != nil {
		runtime.SetError(fmt.Errorf("initializing target path: %s", err))
		runtime.Exit()
	}

	// Publish all URLs
	fmt.Println("Publish all URLs:")
	for index, pub := range cmdConfig.PublishData.URLList {
		if err = task.PublishURLAsPDF(cmdConfig, index+1, pub); err != nil {
			runtime.SetError(fmt.Errorf("publishing URL as PDF: %s", err))
			runtime.Exit()
		}
	}

	// Merge all PDF files
	fmt.Println("Merge all PDF files ...")
	if err = task.MergePDFFiles(cmdConfig); err != nil {
		runtime.SetError(fmt.Errorf("merging all PDF files: %s", err))
		runtime.Exit()
	}

	// Time
	fmt.Printf("\nProcess time: %s\n", time.Since(start))

	// End
	fmt.Printf("\nDone, full PDF file generated at: %s\n\n", cmdConfig.TargetFile)
}
