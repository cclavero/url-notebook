// +build test

package task_test

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/task"
	"github.com/cclavero/ws-pdf-publish/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testOutPathPub          = "out-task-pub"
	testDockerImageBuildTag = "ws-pdf-publish-test-build"
	testDockerImageBuild    = task.DockerImage + ":" + testDockerImageBuildTag
	testDockerImageRunTag   = "ws-pdf-publish-test-run"
	testDockerImageRun      = task.DockerImage + ":" + testDockerImageRunTag
	testOutPathMerge        = "out-task-mrg"
)

var _ = Describe("PDF", func() {

	Context("Docker image WkhtmltoPDF", func() {

		BeforeEach(func() { // Before each 'It' block
		})

		AfterEach(func() { // After each 'It' block
			// Clean up
			removeDockerImage(testDockerImageBuild)
			removeDockerImage(testDockerImageRun)
		})

		When("trying to check, build or use WkhtmltoPDF docker image", func() {

			It("should create a new docker image if not exists", func() {

				testCmdCtx := test.NewTestCmdCtx()
				pdfTask := task.NewPDFTask(testDockerImageBuildTag)

				// Create a new docker image
				removeDockerImage(testDockerImageBuild)

				testCmdCtx.OpenOutCapture()
				err := pdfTask.CheckWkhtmltoPDFDocker()
				result, errResult := testCmdCtx.CloseOutCapture(true, 100)

				Expect(err).To(BeNil())
				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("Building 'wkhtmltopdf' docker image"))
				Expect(existsDockerImage(testDockerImageBuild)).To(BeTrue())

				// Docker image exists
				testCmdCtx.OpenOutCapture()
				err = pdfTask.CheckWkhtmltoPDFDocker()
				result, errResult = testCmdCtx.CloseOutCapture(true, 100)

				Expect(err).To(BeNil())
				Expect(result).To(Equal(""))
				Expect(errResult).To(Equal(""))

			})

			It("trying to publish URL as PDF file", func() {

				testCmdCtx := test.NewTestCmdCtx()
				pdfTask := task.NewPDFTask(testDockerImageRunTag)

				if err := pdfTask.CheckWkhtmltoPDFDocker(); err != nil {
					Fail(err.Error())
				}

				Expect(existsDockerImage(testDockerImageRun)).To(BeTrue())

				testTargetPath := test.GetAbsPath(filepath.Join(test.TestBasePath, testOutPathPub))

				test.RemoveAbsPath(testTargetPath)

				Expect(testTargetPath).Should(Not(BeADirectory()))

				testTargetPathURL := filepath.Join(testTargetPath, config.URLFolder)

				cmdConfig := &config.CmdConfig{
					UserUID:       strconv.Itoa(os.Getuid()),
					UserGID:       strconv.Itoa(os.Getgid()),
					TargetPath:    testTargetPath,
					TargetPathURL: testTargetPathURL,
					TargetFile:    "",
					PublishData: &config.PublishData{
						File: "test.pdf",
						URLList: []config.PublishURL{
							{URL: "https://raw.githubusercontent.com/cclavero/ws-pdf-publish/master/README.md", File: "test.pdf"},
						},
						DockerParams:      "",
						WkhtmltopdfParams: "",
					},
				}

				if err := task.InitTargetPath(cmdConfig); err != nil {
					Fail(err.Error())
				}

				testCmdCtx.OpenOutCapture()
				err := pdfTask.PublishURLsAsPDF(cmdConfig)
				result, errResult := testCmdCtx.CloseOutCapture(true, 200)

				Expect(err).To(BeNil())
				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("	[1] Publishing https://raw.githubusercontent.com/cclavero/ws-pdf-publish/master/README.md"))
				Expect(result).Should(ContainSubstring("/test/out-task-pub/url/test.pdf"))
				targetFile := filepath.Join(cmdConfig.TargetPathURL, cmdConfig.PublishData.URLList[0].File)
				Expect(targetFile).Should(BeARegularFile())

				// Clean up
				test.RemoveAbsPath(testTargetPath)

			})

		})

	})

	Context("Merge PDF files", func() {

		var (
			testTargetPath string
		)

		BeforeEach(func() { // Before each 'It' block
			testTargetPath = test.GetAbsPath(filepath.Join(test.TestBasePath, testOutPathMerge))
		})

		AfterEach(func() { // After each 'It' block
			test.RemoveAbsPath(testTargetPath)
		})

		When("trying to merge PDF files", func() {

			It("should work with existing pdf files", func() {

				test.RemoveAbsPath(testTargetPath)

				Expect(testTargetPath).Should(Not(BeADirectory()))

				testTargetPathURL := filepath.Join(testTargetPath, config.URLFolder)

				// Fail because one missing file
				cmdConfig := &config.CmdConfig{
					UserUID:       strconv.Itoa(os.Getuid()),
					UserGID:       strconv.Itoa(os.Getgid()),
					TargetPath:    testTargetPath,
					TargetPathURL: testTargetPathURL,
					TargetFile:    "",
					PublishData: &config.PublishData{
						File: "test.pdf",
						URLList: []config.PublishURL{
							{URL: "https://sample-1.url", File: "sample-1.pdf"},
							{URL: "https://sample-2.url", File: "sample-2.pdf"},
						},
						DockerParams:      "",
						WkhtmltopdfParams: "",
					},
				}

				if err := task.InitTargetPath(cmdConfig); err != nil {
					Fail(err.Error())
				}

				Expect(cmdConfig.TargetPath).Should(BeADirectory())
				Expect(cmdConfig.TargetPathURL).Should(BeADirectory())

				absSrcFile := test.GetAbsPath(filepath.Join(test.TestBasePath, "sample.pdf"))
				test.CopyFileToAbsPath(absSrcFile, testTargetPathURL, "sample-1.pdf")

				Expect(filepath.Join(cmdConfig.TargetPathURL, "sample-1.pdf")).Should(BeARegularFile())

				pdfTask := task.NewPDFTask(testDockerImageRunTag)
				err := pdfTask.MergePDFFiles(cmdConfig)

				Expect(err).To(Not(BeNil()))

				Expect(err.Error()).Should(ContainSubstring("/out-task-mrg/url/sample-2.pdf: no such file or directory"))

				// Works because add the missing file
				test.CopyFileToAbsPath(absSrcFile, testTargetPathURL, "sample-2.pdf")

				Expect(filepath.Join(cmdConfig.TargetPathURL, "sample-2.pdf")).Should(BeARegularFile())

				err = pdfTask.MergePDFFiles(cmdConfig)

				Expect(err).To(BeNil())

				Expect(filepath.Join(cmdConfig.TargetPath, "test.pdf")).Should(BeARegularFile())

			})

		})

	})

})
