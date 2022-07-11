// Package cmd provides add custom CTL command
package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	buildVersion = "0.1.0"
	RootCMD      = &cobra.Command{
		Use:   "hbase-ctl",
		Short: "hbase a cli tool",
		Long:  `hbase a cli tool, which is a command line tool for operating HBase.`,
	}
)

// Execute 执行root命令
func Execute() error {
	return RootCMD.Execute()
}

func init() {
	// 定义版本号
	RootCMD.Version = fmt.Sprintf("%s %s/%s", buildVersion, runtime.GOOS, runtime.GOARCH)
}
