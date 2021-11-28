// +build test

package config_test

import (
	"testing"

	"github.com/cclavero/ws-pdf-publish/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = BeforeSuite(func() {
})

var _ = AfterSuite(func() {
})

func getNewRootCmd() *cobra.Command {
	rootCmd, err := cmd.NewRootCmd()
	if err != nil {
		Fail(err.Error())
	}
	return rootCmd
}
