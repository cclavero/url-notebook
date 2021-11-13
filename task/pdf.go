package task

import (
	"fmt"
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/config"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

const (
	dockerImageVersion  = "ws-pdf-publish"
	dockerImage         = "wkhtmltopdf:" + dockerImageVersion
	dockerCheckCmd      = "docker version"
	dockerCheckImageCmd = "docker image inspect " + dockerImage
	dockerBuildCmd      = "docker build --tag " + dockerImage + " -f - . %s"
	dockerRunCmd        = "docker run -u %s:%s -v %s:/out %s --name " + dockerImageVersion + " --rm " + dockerImage + " %s %s /out/%s"
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

func CheckWkhtmltoPDFDocker() error {
	// Check docker is installed
	if _, err := execSystemCommand(dockerCheckCmd); err != nil {
		return fmt.Errorf("checking docker is installed: %s", err)
	}

	// Check if the docker image exists; if not build it
	if _, err := execSystemCommand(dockerCheckImageCmd); err != nil {
		fmt.Println("Building 'wkhtmltopdf' docker image ...")
		dockerBuildCmdCmd := fmt.Sprintf(dockerBuildCmd, dockerFile)
		_, err := execSystemCommand(dockerBuildCmdCmd)
		if err != nil {
			return fmt.Errorf("building wkhtmltopdf docker image: %s", err)
		}
	}

	return nil
}

func PublishURLAsPDF(cmdConfig *config.CmdConfig, index int, publishURL config.PublishURL) error {
	targetFile := filepath.Join(cmdConfig.TargetPathURL, publishURL.File)
	fmt.Printf("\t[%d] Publishing %s to %s ...\n", index, publishURL.URL, targetFile)
	dockerRunCmdCmd := fmt.Sprintf(dockerRunCmd, cmdConfig.UserUID, cmdConfig.UserGID, cmdConfig.TargetPathURL,
		cmdConfig.PublishData.DockerParams, cmdConfig.PublishData.WkhtmltopdfParams, publishURL.URL, publishURL.File)
	if _, err := execSystemCommand(dockerRunCmdCmd); err != nil {
		return fmt.Errorf("generating PDF file: %s", err)
	}
	return nil
}

func MergePDFFiles(cmdConfig *config.CmdConfig) error {
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
