package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Show task details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		task, err := svc.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		fmt.Printf("ID:          %s\n", task.ID)
		fmt.Printf("Title:       %s\n", task.Title)
		fmt.Printf("Status:      %s\n", task.Status)
		fmt.Printf("Stage:       %s\n", task.Stage)
		fmt.Printf("Priority:    %s\n", task.Priority)
		fmt.Printf("Source:      %s\n", task.Source)
		fmt.Printf("Created:     %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated:     %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
		if task.Dispatch.Agent != "" {
			fmt.Printf("Agent:       %s\n", task.Dispatch.Agent)
		}
		if task.Dispatch.Model != "" {
			fmt.Printf("Model:       %s\n", task.Dispatch.Model)
		}
		if task.Dispatch.Repository != "" {
			fmt.Printf("Repository:  %s\n", task.Dispatch.Repository)
		}
		if task.Dispatch.Worktree != "" {
			fmt.Printf("Worktree:    %s\n", task.Dispatch.Worktree)
		}
		if task.Dispatch.TerminalSession != "" {
			fmt.Printf("Terminal:    %s\n", task.Dispatch.TerminalSession)
		}
		if task.Dispatch.AssignedAt != nil {
			fmt.Printf("Assigned:    %s\n", task.Dispatch.AssignedAt.Format("2006-01-02 15:04:05"))
		}

		if len(task.Tags) > 0 {
			fmt.Printf("Tags:        %s\n", strings.Join(task.Tags, ", "))
		}
		if task.Description != "" {
			fmt.Printf("\nDescription:\n%s\n", task.Description)
		}
		if task.Body != "" {
			fmt.Printf("\n%s\n", task.Body)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
