package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	WSPDFpublishCmd = "ws-pdf-publish"
	PublishFileFlag = "publishFile"
	TargetPathFlag  = "targetPath"
	URLFolder       = "url"
	ConfigInfoStr   = "\n\t- userUID: %s\n\t- userGID: %s\n\t- targetPath: %s\n\t- targetPathURL: %s\n" +
		"\t- targetFile: %s\n\t- publishData: %+v\n"
)

type PublishURL struct {
	URL  string `yaml:"url"`
	File string `yaml:"file"`
}

type PublishData struct {
	File              string       `yaml:"file"`
	URLList           []PublishURL `mapstructure:"urls"`
	DockerParams      string       `yaml:"dockerParams"`
	WkhtmltopdfParams string       `yaml:"wkhtmltopdfParams"`
}

type CmdConfig struct {
	UserUID       string
	UserGID       string
	TargetPath    string
	TargetPathURL string
	TargetFile    string
	PublishData   *PublishData
}

func GetCmdConfig(cmd *cobra.Command) (*CmdConfig, error) {
	var publishFileFlag, publishFile, targetPathFlag, targetPath string
	var err error

	if publishFileFlag, err = getFlagValue(cmd, PublishFileFlag); err != nil {
		return nil, err
	}
	if publishFile, err = filepath.Abs(publishFileFlag); err != nil {
		return nil, fmt.Errorf("getting '%s' value: %s", PublishFileFlag, err)
	}

	if targetPathFlag, err = getFlagValue(cmd, TargetPathFlag); err != nil {
		return nil, err
	}
	if targetPath, err = filepath.Abs(targetPathFlag); err != nil {
		return nil, fmt.Errorf("getting '%s' value: %s", TargetPathFlag, err)
	}

	cmdConfig := &CmdConfig{
		UserUID:       strconv.Itoa(os.Getuid()),
		UserGID:       strconv.Itoa(os.Getgid()),
		TargetPath:    targetPath,
		TargetPathURL: filepath.Join(targetPath, URLFolder),
	}
	if cmdConfig.PublishData, err = getPublishData(publishFile); err != nil {
		return nil, fmt.Errorf("getting publish data: '%s'; %s", publishFile, err)
	}
	cmdConfig.TargetFile = filepath.Join(targetPath, cmdConfig.PublishData.File)

	return cmdConfig, nil
}

func getFlagValue(cmd *cobra.Command, flagID string) (string, error) {
	flagValue, err := cmd.Flags().GetString(flagID)
	if err != nil {
		return "", fmt.Errorf("getting '%s' flag value: %s", flagID, err)
	}
	if flagValue == "" { // Error
		return "", fmt.Errorf("getting '%s' empty flag value", flagID)
	}
	return flagValue, nil
}

func getPublishData(publishFile string) (*PublishData, error) {
	viper.AddConfigPath(filepath.Dir(publishFile))
	viper.SetConfigName(filepath.Base(publishFile))
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
