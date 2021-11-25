package cmd_test

import (
	"github.com/cclavero/ws-pdf-publish/cmd"
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

				testCtx := test.NewTestCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"-h"})
				testCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCtx.CloseOutCapture(true, 360)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("WebSite PDF Publish is a simple command line tool"))
				Expect(result).Should(ContainSubstring("Usage:\n"))
				Expect(result).Should(ContainSubstring("Flags:\n"))

			})

			It("should work with '-v' parameter and show version", func() {

				testCtx := test.NewTestCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"-v"})
				testCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCtx.CloseOutCapture(true, 200)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("ws-pdf-publish version devel"))

			})

			It("should fail because setting bad flag", func() {

				testCtx := test.NewTestCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"--bad-flag"})
				testCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCtx.CloseOutCapture(true, 100)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix("Error: unknown flag: --bad-flag"))
				Expect(errResult).Should(ContainSubstring("Usage:\n"))

			})

			It("should fail because missing required flags", func() {

				testCtx := test.NewTestCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{})
				testCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCtx.CloseOutCapture(true, 100)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix(`Error: required flag(s) "publishFile", "targetPath" not set`))
				Expect(errResult).Should(ContainSubstring("Usage:\n"))

			})

			It("should fail with invalid flags", func() {

				testCtx := test.NewTestCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"--publishFile", test.TestBasePath + "/not-valid-ws-pub-pdf.yaml",
					"--targetPath", test.TestBasePath + "/out"})
				testCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCtx.CloseOutCapture(true, 500)

				Expect(result).To(Equal(""))
				Expect(errResult).To(Not(Equal("")))

				Expect(errResult).Should(HavePrefix("Error: getting cmd config: getting publish data:"))
				Expect(errResult).Should(ContainSubstring(`Config File "not-valid-ws-pub-pdf.yaml" Not Found`))

			})

			It("should work with valid flags", func() {

				testCtx := test.NewTestCtx()

				rootCmd, err := cmd.NewRootCmd()

				Expect(rootCmd).To(Not(BeNil()))
				Expect(err).To(BeNil())

				rootCmd.SetArgs([]string{"--publishFile", test.TestBasePath + "/ws-pub-pdf-test.yaml",
					"--targetPath", test.TestBasePath + "/out"})
				testCtx.OpenOutCapture()
				rootCmd.Execute()
				result, errResult := testCtx.CloseOutCapture(true, 500)

				Expect(result).To(Not(Equal("")))
				Expect(errResult).To(Equal(""))

				Expect(result).Should(HavePrefix("Starting"))
				Expect(result).Should(ContainSubstring("Config:"))
				Expect(result).Should(ContainSubstring("/test.pdf"))

				Expect(test.TestBasePath + "/out").Should(BeADirectory())
				Expect(test.TestBasePath + "/out/url").Should(BeADirectory())
				Expect(test.TestBasePath + "/out/url/boe.pdf").Should(BeARegularFile())
				Expect(test.TestBasePath + "/out/test.pdf").Should(BeARegularFile())

			})

		})

	})

})
