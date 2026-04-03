package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var deleteForce bool

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		svc := service.NewTaskService(dualStore, dir)

		if !deleteForce {
			task, err := svc.Get(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			fmt.Printf("Delete \"%s\" (%s)? [y/N]: ", task.Title, task.ID)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled.")
				return nil
			}
		}

		if err := svc.Delete(cmd.Context(), args[0]); err != nil {
			return err
		}

		fmt.Printf("Deleted: %s\n", args[0])
		return nil
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "Skip confirmation")
	rootCmd.AddCommand(deleteCmd)
}
