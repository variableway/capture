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
	listStatus string
	listStage  string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		filter := model.TaskFilter{}
		if listStatus != "" {
			if !model.IsValidStatus(listStatus) {
				return fmt.Errorf("invalid status: %s", listStatus)
			}
			s := model.TaskStatus(listStatus)
			filter.Status = &s
		}
		if listStage != "" {
			if !model.IsValidStage(listStage) {
				return fmt.Errorf("invalid stage: %s", listStage)
			}
			stage := model.TaskStage(listStage)
			filter.Stage = &stage
		}

		tasks, err := svc.List(cmd.Context(), filter)
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		fmt.Printf("%-12s %-12s %-12s %-10s %-20s %s\n", "ID", "STATUS", "STAGE", "PRIORITY", "CREATED", "TITLE")
		fmt.Println(strings.Repeat("-", 96))
		for _, t := range tasks {
			fmt.Printf("%-12s %-12s %-12s %-10s %-20s %s\n",
				t.ID,
				t.Status,
				t.Stage,
				t.Priority,
				t.CreatedAt.Format("2006-01-02 15:04"),
				truncate(t.Title, 30),
			)
		}
		fmt.Printf("\nTotal: %d task(s)\n", len(tasks))
		return nil
	},
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func init() {
	listCmd.Flags().StringVarP(&listStatus, "status", "s", "", "Filter by status (todo, in_progress, done, cancelled, archived)")
	listCmd.Flags().StringVar(&listStage, "stage", "", "Filter by stage (inbox, mindstorm, analysis, planning, prd, tasks, dispatch, execution, review)")
	rootCmd.AddCommand(listCmd)
}
