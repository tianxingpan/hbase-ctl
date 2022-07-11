// package main provides hbase ctl
package main

import (
	"fmt"
	"os"

	"github.com/tianxingpan/hbase-ctl/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("Execute err: ", err.Error())
		os.Exit(-1)
	}
}
