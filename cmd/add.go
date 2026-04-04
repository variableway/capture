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
	addStage    string
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

		if addStage != "" {
			if !model.IsValidStage(addStage) {
				return fmt.Errorf("invalid stage: %s (valid: inbox, mindstorm, analysis, planning, prd, tasks, dispatch, execution, review)", addStage)
			}
			opts = append(opts, service.WithStage(model.TaskStage(addStage)))
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
		fmt.Printf("  Status: %s | Stage: %s | Priority: %s\n", task.Status, task.Stage, task.Priority)
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
	addCmd.Flags().StringVar(&addStage, "stage", "inbox", "Task stage: inbox, mindstorm, analysis, planning, prd, tasks, dispatch, execution, review")
	rootCmd.AddCommand(addCmd)
}
