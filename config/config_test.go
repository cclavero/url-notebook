package config_test

import (
	"fmt"

	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/cclavero/ws-pdf-publish/config"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

const (
	// TEMPORAL
	testBasePath = "build/test"
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
				rootCmd.SetArgs([]string{"--publishFile", testBasePath + "/ws-pub-pdf-test.yaml",
					"--targetPath", testBasePath + "/out"})

				cmdConfig, err := config.GetCmdConfig(rootCmd)

				fmt.Printf("------------------------------->%+v,%+v\n\n", cmdConfig, err)

			})

		})

	})

})
