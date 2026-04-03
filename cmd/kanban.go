package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
	"github.com/variableway/innate/capture/internal/tui"
)

var kanbanCmd = &cobra.Command{
	Use:   "kanban",
	Short: "Launch TUI kanban board",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return err
		}
		defer dualStore.Close()

		taskSvc := service.NewTaskService(dualStore, dir)
		app := tui.NewApp(taskSvc)

		p := tea.NewProgram(app, tea.WithAltScreen())
		_, err = p.Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(kanbanCmd)
}
