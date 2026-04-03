package cmd

import (
	"github.com/spf13/cobra"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Feishu Bot management",
}

func init() {
	rootCmd.AddCommand(botCmd)
}
