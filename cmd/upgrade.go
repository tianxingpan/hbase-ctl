// Package cmd provides add custom CTL command
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tianxingpan/hbase-ctl/api/util"
)

var upgradeCmd = &cobra.Command{
	Use:   "update",
	Short: "upgrade hbase-ctl",
	Long:  "upgrade hbase-ctl to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := util.RunCMD("GO111MODULE=on go get -u github.com/tianxingpan/hbase-ctl", "")
		if err != nil {
			fmt.Printf("Upgrade failed, errmsg:%s\n", err.Error())
			return
		}
		fmt.Print(info)
	},
}

func init() {
	RootCMD.AddCommand(upgradeCmd)
}
