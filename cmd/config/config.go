package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

const (
	urlFolder = "url"
)

type PublishURL struct {
	URL  string `yaml:"url"`
	File string `yaml:"file"`
}

type PublishData struct {
	File              string       `yaml:"file"`
	URLList           []PublishURL `mapstructure:"urls"`
	WkhtmltopdfParams string       `yaml:"wkhtmltopdfParams"`
}

type CmdConfig struct {
	UserUID           string
	UserGID           string
	TargetPath        string
	TargetPathURL     string
	DockerExtraParams string
	PublishData       *PublishData
}

func GetCmdConfig() (*CmdConfig, error) {
	args := os.Args[1:]
	params := make(map[string]string)
	for _, param := range args {
		paramArr := strings.Split(param, "=")
		if len(paramArr) != 2 {
			return nil, fmt.Errorf("bag arguments syntax: 'arg=value'")
		}
		params[paramArr[0]] = paramArr[1]
	}

	targetPath, err := filepath.Abs(params["targetPath"])
	if err != nil {
		return nil, fmt.Errorf("bad value for 'targetPath': '%s'; %s", targetPath, err)
	}

	urlNotebookFile, err := filepath.Abs(params["urlNotebookFile"])
	if err != nil {
		return nil, fmt.Errorf("bad value for 'urlNotebookFile': '%s'; %s", urlNotebookFile, err)
	}

	cmdConfig := &CmdConfig{
		UserUID:           strconv.Itoa(os.Getuid()),
		UserGID:           strconv.Itoa(os.Getgid()),
		TargetPath:        targetPath,
		TargetPathURL:     filepath.Join(targetPath, urlFolder),
		DockerExtraParams: params["dockerExtraParams"],
	}

	if cmdConfig.TargetPath == "" || urlNotebookFile == "" {
		return nil, fmt.Errorf("missing arguments: targetPath, urlNotebookFile")
	}

	cmdConfig.PublishData, err = getPublishData(urlNotebookFile)
	if err != nil {
		return nil, fmt.Errorf("error getting publish data: '%s'; %s", urlNotebookFile, err)
	}

	return cmdConfig, nil
}

func getPublishData(urlNotebookFile string) (*PublishData, error) {
	viper.AddConfigPath(filepath.Dir(urlNotebookFile))
	viper.SetConfigName(filepath.Base(urlNotebookFile))
	viper.SetConfigType("yaml")
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

func GetCmdConfigInfo(cmdConfig *CmdConfig) string {
	return fmt.Sprintf("\n- userUID: %s\n- userGID: %s\n- targetPath: %s\n- targetPathURL: %s\n- dockerExtraParams: %s\n- publishData: %+v\n",
		cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPath, cmdConfig.TargetPathURL,
		cmdConfig.DockerExtraParams, cmdConfig.PublishData)
}
