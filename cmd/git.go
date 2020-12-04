package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "ls",
	Short: "ls",
	Long:  `ls -alt`,
	Run: func(cmd *cobra.Command, args []string) {
		a := git.ls(args[0])
		log.Printf(git.ls())
	},
}
