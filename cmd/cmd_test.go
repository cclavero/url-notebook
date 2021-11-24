package cmd_test

import (
	"fmt"

	"github.com/cclavero/ws-pdf-publish/cmd"
	"github.com/cclavero/ws-pdf-publish/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cmd", func() {

	Context("Execute command", func() {

		var (
			testRun *test.TestRuntime
		)

		BeforeEach(func() { // Before each 'It' block
			testRun = test.NewTestRuntime()
		})

		AfterEach(func() { // After each 'It' block
		})

		When("trying to invoke the command", func() {

			It("should work with '-h' parameter and show the help result", func() {

				rootCmd := cmd.NewRootCmd(testRun)
				rootCmd.SetArgs([]string{"-h"})
				testRun.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testRun.CloseOutCapture(true, 360)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("WebSite PDF Publish is a simple command line tool"))
				Expect(result).Should(ContainSubstring("Usage:\n"))
				Expect(result).Should(ContainSubstring("Flags:\n"))

			})

			It("should work with '-v' parameter and show version", func() {

				rootCmd := cmd.NewRootCmd(testRun)
				rootCmd.SetArgs([]string{"-v"})
				testRun.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testRun.CloseOutCapture(true, 200)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("ws-pdf-publish version devel"))

			})

			It("should fail because setting bad flag", func() {

				rootCmd := cmd.NewRootCmd(testRun)
				rootCmd.SetArgs([]string{"--bad-flag"})
				testRun.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testRun.CloseOutCapture(true, 100)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix("Error: unknown flag: --bad-flag"))
				Expect(errResult).Should(ContainSubstring("Usage:\n"))

			})

			It("should fail because missing required flags", func() {

				rootCmd := cmd.NewRootCmd(testRun)
				rootCmd.SetArgs([]string{})
				testRun.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testRun.CloseOutCapture(true, 100)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix(`Error: required flag(s) "publishFile", "targetPath" not set`))
				Expect(errResult).Should(ContainSubstring("Usage:\n"))

			})

			It("should fail with invalid flags", func() {

				rootCmd := cmd.NewRootCmd(testRun)
				rootCmd.SetArgs([]string{"--publishFile", test.TestBasePath + "/not-valid-ws-pub-pdf.yaml",
					"--targetPath", test.TestBasePath + "/out"})
				testRun.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testRun.CloseOutCapture(true, 500)

				//Expect(result).To(Equal(""))
				//Expect(errResult).To(Not(Equal("")))

				// TEMPORAL
				fmt.Printf("--------------------------->%+v,%+v\n\n", result, errResult)

			})

			It("should work with valid flags", func() {

				rootCmd := cmd.NewRootCmd(testRun)
				rootCmd.SetArgs([]string{"--publishFile", test.TestBasePath + "/ws-pub-pdf-test.yaml",
					"--targetPath", test.TestBasePath + "/out"})
				testRun.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testRun.CloseOutCapture(true, 500)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("Starting"))
				Expect(result).Should(ContainSubstring("Config:"))
				Expect(result).Should(ContainSubstring("/test.pdf"))

				Expect(test.TestBasePath + "/out").Should(BeADirectory())
				Expect(test.TestBasePath + "/out/test.pdf").Should(BeARegularFile())
				Expect(test.TestBasePath + "/out/url").Should(BeADirectory())
				Expect(test.TestBasePath + "/out/url/boe.pdf").Should(BeARegularFile())

			})

		})

	})

})
