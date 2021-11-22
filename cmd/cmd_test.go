package cmd_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

// TEMPORAL
// https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra

var _ = Describe("Cmd", func() {

	Context("Execute command", func() {

		BeforeEach(func() { // Before each 'It' block
		})

		AfterEach(func() { // After each 'It' block
		})

		When("trying to invoke the command", func() {

			It("should work with '-h' parameter and show the help result", func() {

				/*
					os.Args = []string{test.TestCmdName, "-h"}

					outCap := test.NewOutCapture()
					cmd.Execute()
					result, errResult := outCap.Close()

					test.LogResult(result, 300)
					test.LogError(errResult)

					Expect(errResult).To(Equal(""))
					Expect(result).To(Not(Equal("")))

					Expect(result).Should(HavePrefix("WebSite PDF Publish is a simple command line tool"))
					Expect(result).Should(ContainSubstring("Usage:\n"))
					Expect(result).Should(ContainSubstring("Flags:\n"))
				*/

			})

			It("should work with '-v' parameter and show version", func() {

				/*
					os.Args = []string{test.TestCmdName, "-v"}

					outCap := test.NewOutCapture()
					cmd.Execute()
					result, errResult := outCap.Close()

					test.LogResult(result, 200)
					test.LogError(errResult)

					Expect(errResult).To(Equal(""))
					Expect(result).To(Not(Equal("")))

					Expect(result).Should(HavePrefix("mc3words version devel"))
				*/

			})

			/*
				It("should work with 'test-short.txt' file", func() {

					os.Args = []string{test.TestCmdName, test.TestBasePath + "/resource/test-short.txt"}

					outCap := test.NewOutCapture()
					cmd.Execute()
					result, errResult := outCap.Close()

					test.LogResult(result, 500)
					test.LogError(errResult)

					Expect(result).To(Not(Equal("")))
					Expect(errResult).To(Equal(""))

					Expect(result).Should(HavePrefix("Starting ..."))
					Expect(result).Should(ContainSubstring("test-short.txt"))
					Expect(result).Should(ContainSubstring("Processing words for input file #1: 412 words"))
					Expect(result).Should(ContainSubstring("Found partial results #1: 401 sequences found"))
					Expect(result).Should(ContainSubstring("1    i'm your father                          5"))
					Expect(result).Should(ContainSubstring("2    father i'm your                          4"))
					Expect(result).Should(ContainSubstring("3    your father i'm                          3"))

				})

				It("should work with 2 files", func() {

					os.Args = []string{test.TestCmdName, test.TestBasePath + "/resource/test-short.txt",
						test.TestBasePath + "/resource/test-short.txt"}

					outCap := test.NewOutCapture()
					cmd.Execute()
					result, errResult := outCap.Close()

					test.LogResult(result, 700)
					test.LogError(errResult)

					Expect(result).To(Not(Equal("")))
					Expect(errResult).To(Equal(""))

					Expect(result).Should(HavePrefix("Starting ..."))
					Expect(result).Should(ContainSubstring("test-short.txt"))
					Expect(result).Should(ContainSubstring("Processing words for input file #2: 412 words"))
					Expect(result).Should(ContainSubstring("Found partial results #2: 401 sequences found"))
					Expect(result).Should(ContainSubstring("1    i'm your father                          10"))
					Expect(result).Should(ContainSubstring("2    father i'm your                          8"))

				})*/

		})

	})

})
