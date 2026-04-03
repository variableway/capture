package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a config value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		val := viper.Get(args[0])
		if val == nil {
			return fmt.Errorf("key %s not found", args[0])
		}
		fmt.Println(val)
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a config value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set(args[0], args[1])
		if err := viper.WriteConfig(); err != nil {
			// If config file doesn't exist, create it
			dir := getDataDir()
			return viper.WriteConfigAs(dir + "/config.yaml")
		}
		fmt.Printf("Set %s = %s\n", args[0], args[1])
		return nil
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	rootCmd.AddCommand(configCmd)
}
