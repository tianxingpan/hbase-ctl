// Package cmd provides add custom CTL command
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCMD = &cobra.Command{
	Use:   "get",
	Short: "Get row or cell contents",
	Long: `Get row or cell contents; pass table name, row, and optionally
	a dictionary of column(s), timestamp, timerange and versions. Examples:
	
	  hbase-ctl get 'ns1:t1', 'r1'
	  hbase-ctl get 't1', 'r1'
	  hbase-ctl get 't1', 'r1', {TIMERANGE => [ts1, ts2]}
	  hbase-ctl get 't1', 'r1', {COLUMN => 'c1'}
	  hbase-ctl get 't1', 'r1', {COLUMN => ['c1', 'c2', 'c3']}
	  hbase-ctl get 't1', 'r1', {COLUMN => 'c1', TIMESTAMP => ts1}
	  hbase-ctl get 't1', 'r1', {COLUMN => 'c1', TIMERANGE => [ts1, ts2], VERSIONS => 4}
	  hbase-ctl get 't1', 'r1', {COLUMN => 'c1', TIMESTAMP => ts1, VERSIONS => 4}
	  hbase-ctl get 't1', 'r1', {FILTER => "ValueFilter(=, 'binary:abc')"}
	  hbase-ctl get 't1', 'r1', 'c1'
	  hbase-ctl get 't1', 'r1', 'c1', 'c2'
	  hbase-ctl get 't1', 'r1', ['c1', 'c2']
	  hbase-ctl get 't1', 'r1', {COLUMN => 'c1', ATTRIBUTES => {'mykey'=>'myvalue'}}
	  hbase-ctl get 't1', 'r1', {COLUMN => 'c1', AUTHORIZATIONS => ['PRIVATE','SECRET']}
	  hbase-ctl get 't1', 'r1', {CONSISTENCY => 'TIMELINE'}
	  hbase-ctl get 't1', 'r1', {CONSISTENCY => 'TIMELINE', REGION_REPLICA_ID => 1}
	
	Besides the default 'toStringBinary' format, 'get' also supports custom formatting by
	column.  A user can define a FORMATTER by adding it to the column name in the get
	specification.  The FORMATTER can be stipulated:
	
	 1. either as a org.apache.hadoop.hbase.util.Bytes method name (e.g, toInt, toString)
	 2. or as a custom class followed by method name: e.g. 'c(MyFormatterClass).format'.
	
	Example formatting cf:qualifier1 and cf:qualifier2 both as Integers:
	  hbase-ctl get 't1', 'r1' {COLUMN => ['cf:qualifier1:toInt',
		'cf:qualifier2:c(org.apache.hadoop.hbase.util.Bytes).toInt'] }
	
	Note that you can specify a FORMATTER by column only (cf:qualifier). You can set a
	formatter for all columns (including, all key parts) using the "FORMATTER"
	and "FORMATTER_CLASS" options. The default "FORMATTER_CLASS" is
	"org.apache.hadoop.hbase.util.Bytes".
	
	  hbase-ctl get 't1', 'r1', {FORMATTER => 'toString'}
	  hbase-ctl get 't1', 'r1', {FORMATTER_CLASS => 'org.apache.hadoop.hbase.util.Bytes', FORMATTER => 'toString'}`,
	Run: getCMDHandler,
}

func init() {
	RootCMD.AddCommand(getCMD)
}

// Processing logic of subcommand get
func getCMDHandler(cmd *cobra.Command, args []string) {
	fmt.Println(buildVersion)
}
