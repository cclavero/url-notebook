// +build test

package config_test

import (
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Context("Get command config", func() {

		BeforeEach(func() { // Before each 'It' block
		})

		AfterEach(func() { // After each 'It' block
		})

		When("trying to get command config", func() {

			It("should work with valid values", func() {

				rootCmd := getNewRootCmd()
				publishFileFlag := test.TestBasePath + "/ws-pub-pdf-test.yaml"
				targetPathFlag := test.TestBasePath + "/out-config"
				rootCmd.Flags().Set(config.PublishFileFlag, publishFileFlag)
				rootCmd.Flags().Set(config.TargetPathFlag, targetPathFlag)
				cmdConfig, err := config.GetCmdConfig(rootCmd)

				Expect(cmdConfig).To(Not(BeNil()))
				Expect(err).To(BeNil())

				Expect(cmdConfig.UserUID).To(Not(BeNil()))
				Expect(cmdConfig.UserGID).To(Not(BeNil()))
				targetPath := test.GetAbsPath(targetPathFlag)
				Expect(cmdConfig.TargetPath).To(Equal(targetPath))
				Expect(cmdConfig.TargetPathURL).To(Equal(targetPath + "/url"))
				Expect(cmdConfig.PublishData).To(Not(BeNil()))
				Expect(cmdConfig.PublishData.File).To(Equal("test.pdf"))
				Expect(len(cmdConfig.PublishData.URLList)).To(Equal(1))
				targetFile := test.GetAbsPath(filepath.Join(targetPathFlag, cmdConfig.PublishData.File))
				Expect(cmdConfig.TargetFile).To(Equal(targetFile))

			})

			It("should fail with empty flag", func() {

				rootCmd := getNewRootCmd()
				cmdConfig, err := config.GetCmdConfig(rootCmd)

				Expect(cmdConfig).To(BeNil())
				Expect(err).To(Not(BeNil()))

				Expect(err.Error()).Should(HavePrefix("getting 'publishFile' empty flag value"))

			})

			It("should fail with invalid publish file", func() {

				rootCmd := getNewRootCmd()

				// Inexistent
				publishFileFlag := test.TestBasePath + "/ws-pub-pdf-test-inexistent.yaml"
				targetPathFlag := test.TestBasePath + "/out-config"
				rootCmd.Flags().Set(config.PublishFileFlag, publishFileFlag)
				rootCmd.Flags().Set(config.TargetPathFlag, targetPathFlag)
				cmdConfig, err := config.GetCmdConfig(rootCmd)

				Expect(cmdConfig).To(BeNil())
				Expect(err).To(Not(BeNil()))

				Expect(err.Error()).Should(ContainSubstring(`Config File "ws-pub-pdf-test-inexistent.yaml" Not Found`))

				// Bad JSON
				publishFileFlag = test.TestBasePath + "/ws-pub-pdf-bad-json.yaml"
				rootCmd.Flags().Set(config.PublishFileFlag, publishFileFlag)
				cmdConfig, err = config.GetCmdConfig(rootCmd)

				Expect(cmdConfig).To(BeNil())
				Expect(err).To(Not(BeNil()))

				Expect(err.Error()).Should(ContainSubstring("While parsing config: yaml"))

				// Bad JSON values
				publishFileFlag = test.TestBasePath + "/ws-pub-pdf-bad-json-values.yaml"
				rootCmd.Flags().Set(config.PublishFileFlag, publishFileFlag)
				cmdConfig, err = config.GetCmdConfig(rootCmd)

				Expect(cmdConfig).To(BeNil())
				Expect(err).To(Not(BeNil()))

				Expect(err.Error()).Should(ContainSubstring("empty values in config file: file, urls"))

			})

		})

	})

})
