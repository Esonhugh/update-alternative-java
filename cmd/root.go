package cmd

import (
	"github.com/esonhugh/update-alternative-java/lib/misc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "update-alternative-java",
	Short:   "Update alternative as macos java version manager",
	Long:    `Update alternative java Via Path injection`,
	Version: misc.ToolVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.InfoLevel)
		log.Println("update alternative of java little tool")
		log.Info("Based on Path injection")
	},
}
