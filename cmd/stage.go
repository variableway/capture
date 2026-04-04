package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var stageCmd = &cobra.Command{
	Use:   "stage <id> <new-stage>",
	Short: "Change task stage in the task center pipeline",
	Long:  "Change task stage. Valid stages: inbox, mindstorm, analysis, planning, prd, tasks, dispatch, execution, review",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		newStage := model.TaskStage(args[1])
		if !model.IsValidStage(string(newStage)) {
			return fmt.Errorf("invalid stage: %s\nValid stages: inbox, mindstorm, analysis, planning, prd, tasks, dispatch, execution, review", args[1])
		}

		task, err := svc.Update(cmd.Context(), args[0], service.WithStage(newStage))
		if err != nil {
			return err
		}

		fmt.Printf("Stage changed: %s -> %s (%s)\n", args[0], newStage, task.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stageCmd)
}
