// +build test

package task_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cclavero/ws-pdf-publish/test"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTask(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Task Suite")
}

var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {
})

func initTaskOutPath(testTargetPath string, testTargetPathURL string) {
	if err := os.MkdirAll(testTargetPathURL, os.ModePerm); err != nil {
		Fail(err.Error())
	}
	absSrcFile := test.GetAbsPath(filepath.Join(test.TestBasePath, "sample.pdf"))
	test.CopyFileToAbsPath(absSrcFile, testTargetPathURL, "sample.pdf")
	test.CopyFileToAbsPath(absSrcFile, testTargetPath, "test.pdf")
}

func removeDockerImage(dockerImage string) {
	dockerRemoveCmd := fmt.Sprintf("docker rmi %s || true", dockerImage)
	test.ExecSysCommand(dockerRemoveCmd)
}

func existsDockerImage(dockerImage string) bool {
	dockerImagePart := strings.Split(dockerImage, ":")
	if len(dockerImagePart) != 2 {
		ginkgo.Fail("bag image tag format")
	}
	dockerListCmd := fmt.Sprintf("docker images | grep -E '%s.*%s' || true", dockerImagePart[0], dockerImagePart[1])
	result := test.ExecSysCommand(dockerListCmd)
	return strings.Contains(result, dockerImagePart[0]) && strings.Contains(result, dockerImagePart[1])
}
