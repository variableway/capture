package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/bitable"
	"github.com/variableway/innate/capture/internal/feishu"
	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/store"
)

var syncDirection string

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync tasks with Feishu Bitable",
	Long: `Sync tasks between local storage and Feishu Bitable (多维表格).

Directions:
  push          - Push local tasks to Feishu Bitable
  pull          - Pull Feishu Bitable records to local
  bidirectional - Sync both ways (default)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		appID := os.Getenv("FEISHU_APP_ID")
		appSecret := os.Getenv("FEISHU_APP_SECRET")
		appToken := os.Getenv("FEISHU_BITABLE_APP_TOKEN")
		tableID := os.Getenv("FEISHU_BITABLE_TABLE_ID")

		if appID == "" || appSecret == "" {
			return fmt.Errorf("FEISHU_APP_ID and FEISHU_APP_SECRET are required")
		}
		if appToken == "" || tableID == "" {
			return fmt.Errorf("FEISHU_BITABLE_APP_TOKEN and FEISHU_BITABLE_TABLE_ID are required")
		}

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		client := feishu.NewClient(appID, appSecret)
		bitableClient := bitable.NewClient(client, appToken, tableID)
		engine := bitable.NewSyncEngine(bitableClient, dualStore)

		switch syncDirection {
		case "push":
			tasks, err := dualStore.ListTasks(cmd.Context(), model.TaskFilter{})
			if err != nil {
				return err
			}
			result, err := engine.PushToBitable(cmd.Context(), tasks)
			if err != nil {
				return err
			}
			fmt.Printf("Push complete: %d tasks synced\n", result.Pushed)

		case "pull":
			result, err := engine.PullFromBitable(cmd.Context())
			if err != nil {
				return err
			}
			fmt.Printf("Pull complete: %d records pulled\n", result.Pulled)

		case "bidirectional":
			result, err := engine.Sync(cmd.Context())
			if err != nil {
				return err
			}
			fmt.Printf("Sync complete: %d pushed, %d pulled\n", result.Pushed, result.Pulled)
			if len(result.Errors) > 0 {
				fmt.Println("\nErrors:")
				for _, e := range result.Errors {
					fmt.Printf("  - %s\n", e)
				}
			}

		default:
			return fmt.Errorf("invalid direction: %s (use: push, pull, bidirectional)", syncDirection)
		}

		return nil
	},
}

func init() {
	syncCmd.Flags().StringVarP(&syncDirection, "direction", "d", "bidirectional", "Sync direction: push, pull, bidirectional")
	rootCmd.AddCommand(syncCmd)
}
