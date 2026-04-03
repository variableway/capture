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
	addDesc     string
	addTags     string
	addPriority string
)

var addCmd = &cobra.Command{
	Use:   "add <title>",
	Short: "Add a new idea/task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.Join(args, " ")
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		opts := []service.TaskOption{
			service.WithDescription(addDesc),
		}

		if addPriority != "" {
			if !model.IsValidPriority(addPriority) {
				return fmt.Errorf("invalid priority: %s (valid: high, medium, low)", addPriority)
			}
			opts = append(opts, service.WithPriority(model.TaskPriority(addPriority)))
		}

		if addTags != "" {
			tags := strings.Split(addTags, ",")
			for i, t := range tags {
				tags[i] = strings.TrimSpace(t)
			}
			opts = append(opts, service.WithTags(tags))
		}

		task, err := svc.Create(cmd.Context(), title, opts...)
		if err != nil {
			return err
		}

		fmt.Printf("Created: %s - %s\n", task.ID, task.Title)
		fmt.Printf("  Status: %s | Priority: %s\n", task.Status, task.Priority)
		if len(task.Tags) > 0 {
			fmt.Printf("  Tags: %s\n", strings.Join(task.Tags, ", "))
		}
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&addDesc, "description", "d", "", "Task description")
	addCmd.Flags().StringVarP(&addTags, "tags", "t", "", "Tags (comma-separated)")
	addCmd.Flags().StringVarP(&addPriority, "priority", "p", "medium", "Priority: high, medium, low")
	rootCmd.AddCommand(addCmd)
}
