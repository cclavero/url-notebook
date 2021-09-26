package pdf

import (
	"fmt"
	"path/filepath"

	"github.com/cclavero/url-notebook/cmd/config"
	"github.com/cclavero/url-notebook/cmd/task"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

const (
	DOCKERFILE = `<<EOF
FROM ubuntu:20.04

RUN apt-get update && \
    apt-get upgrade -y;
    
RUN apt-get install wget sudo -y && \
    wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb && \
    sudo apt install ./wkhtmltox_0.12.6-1.focal_amd64.deb -y;

ENTRYPOINT ["wkhtmltopdf"]

CMD ["-h"]
EOF
`
	DOCKER_IMAGE           = "url-notebook:local"
	DOCKER_CHECK_CMD       = "docker version"
	DOCKER_CHECK_IMAGE_CMD = "docker image inspect " + DOCKER_IMAGE
	DOCKER_BUILD_CMD       = "docker build --tag " + DOCKER_IMAGE + " -f - . %s"
	DOCKER_RUN_CMD         = "docker run -u %s:%s -v %s:/out %s --rm " + DOCKER_IMAGE + " %s %s /out/%s"
)

func BuildWkhtmltopdfDocker() error {
	// Check docker is installed
	if _, err := task.ExecSystemCommand(DOCKER_CHECK_CMD); err != nil {
		return fmt.Errorf("checking docker is installed: %s", err)
	}

	// Check if the docker image exists
	if _, err := task.ExecSystemCommand(DOCKER_CHECK_IMAGE_CMD); err != nil {
		dockerBuildCmd := fmt.Sprintf(DOCKER_BUILD_CMD, DOCKERFILE)
		_, err := task.ExecSystemCommand(dockerBuildCmd)
		if err != nil {
			return fmt.Errorf("building wkhtmltopdf docker image: %s", err)
		}
	}

	return nil
}

func PublishURLAsPDF(cmdConfig *config.CmdConfig, publishURL config.PublishURL) error {
	dockerRunCmd := fmt.Sprintf(DOCKER_RUN_CMD, cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPath,
		cmdConfig.DockerExtraParams, cmdConfig.PublishData.WkhtmltopdfParams, publishURL.URL, publishURL.File)
	if _, err := task.ExecSystemCommand(dockerRunCmd); err != nil {
		return fmt.Errorf("generating PDF file: %s", err)
	}
	return nil
}

func MergePDFFiles(cmdConfig *config.CmdConfig) error {
	inFiles := []string{}
	for _, pub := range cmdConfig.PublishData.URLList {
		inFiles = append(inFiles, filepath.Join(cmdConfig.TargetPath, pub.File))
	}
	fmt.Printf("- PDF files to merge: %+v\n", inFiles)
	outPublishFile := filepath.Join(cmdConfig.TargetPath, cmdConfig.PublishData.File)
	if err := pdfcpu.MergeCreateFile(inFiles, outPublishFile, nil); err != nil {
		return fmt.Errorf("merging PDF files: %s", err)
	}
	return nil
}
