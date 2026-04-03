package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var (
	editTitle       string
	editDesc        string
	editTags        string
	editPriority    string
)

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "Edit a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		opts := []service.TaskOption{}
		if editTitle != "" {
			opts = append(opts, func(t *model.Task) { t.Title = editTitle })
		}
		if editDesc != "" {
			opts = append(opts, service.WithDescription(editDesc))
		}
		if editPriority != "" {
			if !model.IsValidPriority(editPriority) {
				return fmt.Errorf("invalid priority: %s", editPriority)
			}
			opts = append(opts, service.WithPriority(model.TaskPriority(editPriority)))
		}
		if editTags != "" {
			tags := strings.Split(editTags, ",")
			for i, t := range tags {
				tags[i] = strings.TrimSpace(t)
			}
			opts = append(opts, service.WithTags(tags))
		}

		task, err := svc.Update(cmd.Context(), args[0], opts...)
		if err != nil {
			return err
		}

		fmt.Printf("Updated: %s - %s\n", task.ID, task.Title)
		return nil
	},
}

func init() {
	editCmd.Flags().StringVarP(&editTitle, "title", "T", "", "New title")
	editCmd.Flags().StringVarP(&editDesc, "description", "d", "", "New description")
	editCmd.Flags().StringVarP(&editTags, "tags", "t", "", "New tags (comma-separated)")
	editCmd.Flags().StringVarP(&editPriority, "priority", "p", "", "New priority (high, medium, low)")
	rootCmd.AddCommand(editCmd)
}
