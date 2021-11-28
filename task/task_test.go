// +build test

package task_test

import (
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/task"
	"github.com/cclavero/ws-pdf-publish/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testOutPathITP = "out-task-itp"
)

var _ = Describe("Task", func() {

	Context("Init target path", func() {

		var (
			testTargetPath string
		)

		BeforeEach(func() { // Before each 'It' block
			testTargetPath = test.GetAbsPath(filepath.Join(test.TestBasePath, testOutPathITP))
		})

		AfterEach(func() { // After each 'It' block
			test.RemoveAbsPath(testTargetPath)
		})

		When("trying to init target path", func() {

			It("should work with not existing target path", func() {

				test.RemoveAbsPath(testTargetPath)

				Expect(testTargetPath).Should(Not(BeADirectory()))

				testTargetPathURL := filepath.Join(testTargetPath, config.URLFolder)

				cmdConfig := &config.CmdConfig{
					TargetPath:    testTargetPath,
					TargetPathURL: testTargetPathURL,
				}

				err := task.InitTargetPath(cmdConfig)

				Expect(err).To(BeNil())

				Expect(cmdConfig.TargetPath).Should(BeADirectory())
				Expect(cmdConfig.TargetPathURL).Should(BeADirectory())

			})

			It("should work with existing target path", func() {

				test.RemoveAbsPath(testTargetPath)

				Expect(testTargetPath).Should(Not(BeADirectory()))

				testTargetPathURL := filepath.Join(testTargetPath, config.URLFolder)
				initTaskOutPath(testTargetPath, testTargetPathURL)

				cmdConfig := &config.CmdConfig{
					TargetPath:    testTargetPath,
					TargetPathURL: testTargetPathURL,
				}

				err := task.InitTargetPath(cmdConfig)

				Expect(err).To(BeNil())

				Expect(cmdConfig.TargetPath).Should(BeADirectory())
				Expect(cmdConfig.TargetPathURL).Should(BeADirectory())

			})

		})

	})

})
