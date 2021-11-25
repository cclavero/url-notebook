package config_test

import (
	"fmt"
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	publishFileFlag = test.TestBasePath + "/ws-pub-pdf-test.yaml"
	targetPathFlag  = test.TestBasePath + "/out"
)

var _ = Describe("Config", func() {

	Context("Get command config", func() {

		BeforeEach(func() { // Before each 'It' block
		})

		AfterEach(func() { // After each 'It' block
		})

		When("trying to get command config", func() {

			It("should work with valid values", func() {

				rootCmd, _ := cmd.NewRootCmd()
				rootCmd.Flags().Set(config.PublishFileFlag, publishFileFlag)
				rootCmd.Flags().Set(config.TargetPathFlag, targetPathFlag)
				cmdConfig, err := config.GetCmdConfig(rootCmd)

				Expect(cmdConfig).To(Not(BeNil()))
				Expect(err).To(BeNil())

				Expect(cmdConfig.UserUID).To(Not(BeNil()))
				Expect(cmdConfig.UserGID).To(Not(BeNil()))
				targetPath, _ := filepath.Abs(targetPathFlag)
				Expect(cmdConfig.TargetPath).To(Equal(targetPath))
				Expect(cmdConfig.TargetPathURL).To(Equal(targetPath + "/url"))
				targetFile, _ := filepath.Abs(publishFileFlag)
				//Expect(cmdConfig.TargetFile).To(Equal(targetFile))

				// TEMPORAL
				fmt.Printf("--------------->%+v,%+v\n\n", cmdConfig, targetFile)

			})

		})

	})

})
