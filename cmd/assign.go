package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var (
	assignAgent    string
	assignModel    string
	assignRepo     string
	assignWorktree string
	assignTerminal string
)

var assignCmd = &cobra.Command{
	Use:   "assign <id>",
	Short: "Assign a task to an AI agent workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if assignAgent == "" {
			return fmt.Errorf("agent is required")
		}

		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)
		assignedAt := time.Now()
		dispatch := model.TaskDispatch{
			Agent:           assignAgent,
			Model:           assignModel,
			Repository:      assignRepo,
			Worktree:        assignWorktree,
			TerminalSession: assignTerminal,
			AssignedAt:      &assignedAt,
		}

		task, err := svc.Update(
			cmd.Context(),
			args[0],
			service.WithDispatch(dispatch),
			service.WithStage(model.StageDispatch),
		)
		if err != nil {
			return err
		}

		fmt.Printf("Assigned: %s -> %s\n", task.ID, task.Dispatch.Agent)
		if task.Dispatch.Repository != "" {
			fmt.Printf("  Repository: %s\n", task.Dispatch.Repository)
		}
		if task.Dispatch.Worktree != "" {
			fmt.Printf("  Worktree: %s\n", task.Dispatch.Worktree)
		}
		if task.Dispatch.TerminalSession != "" {
			fmt.Printf("  Terminal: %s\n", task.Dispatch.TerminalSession)
		}
		fmt.Printf("  Stage: %s\n", task.Stage)
		return nil
	},
}

func init() {
	assignCmd.Flags().StringVar(&assignAgent, "agent", "", "Agent name")
	assignCmd.Flags().StringVar(&assignModel, "model", "", "Model name")
	assignCmd.Flags().StringVar(&assignRepo, "repo", "", "Repository path")
	assignCmd.Flags().StringVar(&assignWorktree, "worktree", "", "Worktree path")
	assignCmd.Flags().StringVar(&assignTerminal, "terminal", "", "Terminal session identifier")
	rootCmd.AddCommand(assignCmd)
}
