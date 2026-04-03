package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	dataDir string
)

var rootCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture - 捕捉灵感，管理任务",
	Long: `Capture 是一个想法捕捉与任务管理工具。
通过 Terminal 或飞书 Bot 快速记录灵感，保存为 Markdown 文件并同步到飞书多维表格。`,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.capture/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "", "data directory (default is $HOME/.capture)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		viper.AddConfigPath(home + "/.capture")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CAPTURE")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func getDataDir() string {
	if dataDir != "" {
		return dataDir
	}
	if d := viper.GetString("data_dir"); d != "" {
		return d
	}
	home, _ := os.UserHomeDir()
	return home + "/.capture"
}
