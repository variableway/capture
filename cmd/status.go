package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var statusCmd = &cobra.Command{
	Use:   "status <id> <new-status>",
	Short: "Change task status",
	Long:  "Change task status. Valid statuses: todo, in_progress, done, cancelled, archived",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		newStatus := model.TaskStatus(args[1])
		if !model.IsValidStatus(string(newStatus)) {
			return fmt.Errorf("invalid status: %s\nValid statuses: todo, in_progress, done, cancelled, archived", args[1])
		}

		task, err := svc.SetStatus(cmd.Context(), args[0], newStatus)
		if err != nil {
			return err
		}

		fmt.Printf("Status changed: %s -> %s (%s)\n", args[0], newStatus, task.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
