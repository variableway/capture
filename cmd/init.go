package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/config"
	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/store"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize capture data directory and configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()
		fmt.Printf("Initializing capture at %s\n", dir)

		// Create directory structure
		dirs := []string{
			dir,
			filepath.Join(dir, "tasks"),
			filepath.Join(dir, "logs"),
		}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", d, err)
			}
		}

		// Create default config if not exists
		configPath := filepath.Join(dir, "config.yaml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			cfg := model.DefaultConfig()
			cfg.App.DataDir = dir
			if err := config.Save(dir, cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}
			fmt.Println("Created default config.yaml")
		} else {
			fmt.Println("Config already exists, skipping")
		}

		// Initialize SQLite database
		db, err := store.NewSQLiteStore(dir)
		if err != nil {
			return fmt.Errorf("failed to initialize database: %w", err)
		}
		db.Close()
		fmt.Println("Initialized database")

		fmt.Println("\nCapture initialized successfully!")
		fmt.Println("\nQuick start:")
		fmt.Println("  capture add \"我的第一个想法\"")
		fmt.Println("  capture list")
		fmt.Println("  capture kanban")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
