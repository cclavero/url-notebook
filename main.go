package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/spf13/viper"
)

const (
	Version         = "1.0"
	PublishURLsFile = "publish-urls.yaml"
	DockerCmd       = "docker run -u %s:%s -v %s:/out %s url-notebook:1.0 %s %s /out/%s"
)

type CmdConfig struct {
	UserUID           string
	UserGID           string
	TargetPath        string
	DockerExtraParams string
}

type PublishURL struct {
	URL  string `yaml:"url"`
	File string `yaml:"file"`
}

type PublishData struct {
	File              string       `yaml:"file"`
	URLList           []PublishURL `mapstructure:"urls"`
	WkhtmltopdfParams string       `yaml:"wkhtmltopdfParams"`
}

var (
	cmdConfig   *CmdConfig
	publishData *PublishData
)

func init() {
	// Info
	fmt.Println(fmt.Sprintf("PublishPDF: v%s", Version))

	// Get config
	var err error
	if cmdConfig, err = getCmdConfig(); err != nil {
		panic(fmt.Sprintf("Error: Getting cmd config: %s\n", err))
	}

	// Get publish data
	if publishData, err = getPublishData(); err != nil {
		panic(fmt.Sprintf("Error: Reading publish URLs file: %s\n", err))
	}
}

func main() {
	// Info
	fmt.Printf("\n- Starting\n\n")
	fmt.Printf("- Command config: %+v\n", cmdConfig)
	fmt.Printf("- PublishData: %+v\n", publishData)

	// Remove old PDF files
	fmt.Printf("\n> Remove old PDF files ...\n\n")
	if err := removePDFFiles(); err != nil {
		panic(fmt.Sprintf("Error: Removing old PDF files: %s\n", err))
	}

	// Publish all URLs
	fmt.Printf("> Publish all URLs ...\n\n")
	if err := publishURLs(); err != nil {
		panic(fmt.Sprintf("Error: Publishing all URLs: %s\n", err))
	}

	// Merge all PDF files
	fmt.Printf("> Merge all PDF files ...\n\n")
	if err := mergePDFFiles(); err != nil {
		panic(fmt.Sprintf("Error: Merging all PDF files: %s\n", err))
	}

	// End
	fmt.Printf("\n- Done\n\n")
}

func getCmdConfig() (*CmdConfig, error) {
	args := os.Args[1:]
	params := make(map[string]string)
	for _, param := range args {
		paramArr := strings.Split(param, "=")
		if len(paramArr) != 2 {
			return nil, fmt.Errorf("bag arguments syntax: 'arg=value'")
		}
		params[paramArr[0]] = paramArr[1]
	}

	cmdConfig := &CmdConfig{
		UserUID:           params["userUID"],
		UserGID:           params["userGID"],
		TargetPath:        params["targetPath"],
		DockerExtraParams: params["dockerExtraParams"],
	}

	if cmdConfig.UserUID == "" || cmdConfig.UserGID == "" || cmdConfig.TargetPath == "" || cmdConfig.DockerExtraParams == "" {
		return nil, fmt.Errorf("missing arguments: userUID, userGID, targetPath, dockerExtraParams")
	}

	return cmdConfig, nil
}

func getPublishData() (*PublishData, error) {
	viper.SetConfigName(PublishURLsFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error in config file: %s", err)
	}

	publishData := &PublishData{}
	if err := viper.UnmarshalKey("publish", publishData); err != nil {
		return nil, fmt.Errorf("fatal unmarshalling config file: %s", err)
	}

	if publishData.File == "" || len(publishData.URLList) < 1 {
		return nil, fmt.Errorf("empty values in config file: file, urls")
	}

	return publishData, nil
}

func removePDFFiles() error {
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
	return nil
}

func publishURLs() error {
	var cmdStr string
	var cmd *exec.Cmd
	for index, pub := range publishData.URLList {
		cmdStr = fmt.Sprintf(DockerCmd, cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPath,
			cmdConfig.DockerExtraParams, publishData.WkhtmltopdfParams, pub.URL, pub.File)
		fmt.Printf("%d. Publishing %s to %s ...\n", index+1, pub.URL, pub.File)
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
		fmt.Printf("%s\n", cmdStr)
		stdout, err := cmd.CombinedOutput()
		fmt.Printf("%s\n\n", stdout)
		if err != nil {
			return fmt.Errorf("error publishing: %s", err)
		}
	}
	return nil
}

func mergePDFFiles() error {
	inFiles := []string{}
	for _, pub := range publishData.URLList {
		inFiles = append(inFiles, filepath.Join(cmdConfig.TargetPath, pub.File))
	}
	fmt.Printf("> PDF files to merge: %+v\n", inFiles)
	outPublishFile := filepath.Join(cmdConfig.TargetPath, publishData.File)
	if err := pdfcpu.MergeCreateFile(inFiles, outPublishFile, nil); err != nil {
		return fmt.Errorf("merging PDF files: %s", err)
	}
	return nil
}
