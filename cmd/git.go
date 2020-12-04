package cmd

import (
	"log"

	_ "github.com/bluesky6529/golang_LinuxScript/internal/git"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git1",
	Long:  "git2",
	Run: func(cmd *cobra.Command, args []string) {
		var pathstring string
		i := git.lsfunction()
		log.Printf("%s", i)
	},
}

var path string

func init() {
	gitCmd.Flags().StringVarP(&path, "path", "p", " ", "123")
}
