// Package cmd provides add custom CTL command
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of hbase-ctl",
	Long:  `This subcommand is used to view the specific build version of hbase-ctl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildVersion)
	},
}

func init() {
	RootCMD.AddCommand(versionCmd)
}
