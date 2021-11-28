// +build test

package cmd_test

import (
	"path/filepath"

	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/cclavero/ws-pdf-publish/config"
	"github.com/cclavero/ws-pdf-publish/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cmd", func() {

	Context("Execute command", func() {

		BeforeEach(func() { // Before each 'It' block
		})

		AfterEach(func() { // After each 'It' block
		})

		When("trying to invoke the command", func() {

			It("should work with '-h' parameter and show the help result", func() {

				testCmdCtx := test.NewTestCmdCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"-h"})
				testCmdCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCmdCtx.CloseOutCapture(true, 360)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("WebSite PDF Publish is a simple command line tool"))
				Expect(result).Should(ContainSubstring("Usage:\n"))
				Expect(result).Should(ContainSubstring("Flags:\n"))

			})

			It("should work with '-v' parameter and show version", func() {

				testCmdCtx := test.NewTestCmdCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"-v"})
				testCmdCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCmdCtx.CloseOutCapture(true, 200)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("ws-pdf-publish version devel"))

			})

			It("should fail because setting bad flag", func() {

				testCmdCtx := test.NewTestCmdCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"--bad-flag"})
				testCmdCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCmdCtx.CloseOutCapture(true, 100)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix("Error: unknown flag: --bad-flag"))
				Expect(errResult).Should(ContainSubstring("Usage:\n"))

			})

			It("should fail because missing required flags", func() {

				testCmdCtx := test.NewTestCmdCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{})
				testCmdCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCmdCtx.CloseOutCapture(true, 100)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix(`Error: required flag(s) "publishFile", "targetPath" not set`))
				Expect(errResult).Should(ContainSubstring("Usage:\n"))

			})

			It("should fail with invalid flags", func() {

				testCmdCtx := test.NewTestCmdCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"--publishFile", test.TestBasePath + "/not-valid-ws-pub-pdf.yaml",
					"--targetPath", test.TestBasePath + "/out-cmd"})
				testCmdCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCmdCtx.CloseOutCapture(true, 500)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix("Error: getting cmd config: getting publish data:"))
				Expect(errResult).Should(ContainSubstring(`Config File "not-valid-ws-pub-pdf.yaml" Not Found`))

			})

			It("should work with valid flags", func() {

				testCmdCtx := test.NewTestCmdCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				targetPathOut := filepath.Join(test.TestBasePath, "out")
				rootCmd.SetArgs([]string{"--publishFile", test.TestBasePath + "/ws-pub-pdf-test.yaml",
					"--targetPath", targetPathOut})
				testCmdCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCmdCtx.CloseOutCapture(true, 500)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("Starting"))
				Expect(result).Should(ContainSubstring("Config:"))
				Expect(result).Should(ContainSubstring("/test.pdf"))

				Expect(targetPathOut).Should(BeADirectory())
				targetPathOutURL := filepath.Join(targetPathOut, config.URLFolder)
				Expect(targetPathOutURL).Should(BeADirectory())
				Expect(targetPathOutURL + "/boe.pdf").Should(BeARegularFile())
				Expect(targetPathOut + "/test.pdf").Should(BeARegularFile())

			})

		})

	})

})
