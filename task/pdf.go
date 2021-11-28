package task

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/config"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

const (
	DockerImage         = "wkhtmltopdf"
	dockerCheckCmd      = "docker version"
	dockerCheckImageCmd = "docker image inspect %s"
	dockerBuildCmd      = "docker build --tag %s -f - . %s"
	dockerRunCmd        = "docker run -u %%s:%%s -v %%s:/out %%s --rm %s %%s %%s /out/%%s"
	dockerFile          = `<<EOF
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
)

type PDFTask struct {
	dockerImage         string
	dockerCheckCmd      string
	dockerCheckImageCmd string
	dockerBuildCmd      string
	dockerRunCmd        string
}

func NewPDFTask(dockerImageTag string) *PDFTask {
	dockerImageWithTag := fmt.Sprintf("%s:%s", DockerImage, dockerImageTag)
	return &PDFTask{
		dockerImage:         dockerImageWithTag,
		dockerCheckCmd:      dockerCheckCmd,
		dockerCheckImageCmd: fmt.Sprintf(dockerCheckImageCmd, dockerImageWithTag),
		dockerBuildCmd:      fmt.Sprintf(dockerBuildCmd, dockerImageWithTag, dockerFile),
		dockerRunCmd:        fmt.Sprintf(dockerRunCmd, dockerImageWithTag),
	}
}

func (pdfTask *PDFTask) CheckWkhtmltoPDFDocker() error {
	// Check docker is installed
	if _, err := pdfTask.execDockerCommand(pdfTask.dockerCheckCmd); err != nil {
		return fmt.Errorf("checking docker is installed: %s", err)
	}

	// Check if the docker image exists; if not build it
	if _, err := pdfTask.execDockerCommand(pdfTask.dockerCheckImageCmd); err != nil {
		fmt.Println("Building 'wkhtmltopdf' docker image ...")
		_, err := pdfTask.execDockerCommand(pdfTask.dockerBuildCmd)
		if err != nil {
			return fmt.Errorf("building wkhtmltopdf docker image: %s", err)
		}
	}

	return nil
}

func (pdfTask *PDFTask) PublishURLsAsPDF(cmdConfig *config.CmdConfig) error {
	for index, pub := range cmdConfig.PublishData.URLList {
		targetFile := filepath.Join(cmdConfig.TargetPathURL, pub.File)
		fmt.Printf("\t[%d] Publishing %s to %s ...\n", index+1, pub.URL, targetFile)
		dockerRunCmd := fmt.Sprintf(pdfTask.dockerRunCmd, cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPathURL,
			cmdConfig.PublishData.DockerParams, cmdConfig.PublishData.WkhtmltopdfParams, pub.URL, pub.File)
		if _, err := pdfTask.execDockerCommand(dockerRunCmd); err != nil {
			return fmt.Errorf("generating PDF file: %s", err)
		}
	}
	return nil
}

func (pdfTask *PDFTask) execDockerCommand(cmdStr string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("system command: '%s'; %s", cmdStr, err)
	}
	return string(stdout), nil
}

func (pdfTask *PDFTask) MergePDFFiles(cmdConfig *config.CmdConfig) error {
	inFiles := []string{}
	for _, pub := range cmdConfig.PublishData.URLList {
		inFiles = append(inFiles, filepath.Join(cmdConfig.TargetPathURL, pub.File))
	}
	outPublishFile := filepath.Join(cmdConfig.TargetPath, cmdConfig.PublishData.File)
	if err := pdfcpu.MergeCreateFile(inFiles, outPublishFile, nil); err != nil {
		return fmt.Errorf("merging PDF files: %s", err)
	}
	return nil
}
