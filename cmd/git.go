package cmd

import (
	"github.com/bluesky6529/golang_LinuxScript/internal/git"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git1",
	Long:  "git2",
	Run: func(cmd *cobra.Command, args []string) {

		git.LsFunction()
		//log.Printf("%s", i)
	},
}
