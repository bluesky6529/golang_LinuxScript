package cmd

import (
	"log"

	"github.com/bluesky6529/golang_LinuxScript/internal/git"

	"github.com/spf13/cobra"
)

var str string
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "ls -alh",
	Long:  "ls -alh function test",
	Run: func(cmd *cobra.Command, args []string) {
		//var pathstring string
		git.LsFunction(str)
		//git.LsFunction()
		log.Printf("%s", str)
	},
}

func init() {
	gitCmd.Flags().StringVarP(&str, "str", "s", ".", "輸入路徑內容(required)")
}
