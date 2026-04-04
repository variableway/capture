package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/variableway/innate/capture/internal/model"
	"github.com/variableway/innate/capture/internal/service"
)

type viewState int

const (
	viewKanban viewState = iota
	viewDetail
)

type column struct {
	title string
	width int
	tasks []*model.Task
}

// App is the root bubbletea Model for the TUI application.
type App struct {
	taskSvc *service.TaskService
	columns []column
	cursor  struct {
		col int
		row int
	}
	state    viewState
	help     help.Model
	keys     keyMap
	selected *model.Task
	err      error
	width    int
	height   int
}

func NewApp(taskSvc *service.TaskService) *App {
	return &App{
		taskSvc: taskSvc,
		columns: []column{
			{title: "TODO", width: 30},
			{title: "IN PROGRESS", width: 30},
			{title: "DONE", width: 30},
		},
		help:  help.New(),
		keys:  keys,
		state: viewKanban,
	}
}

func (a *App) Init() tea.Cmd {
	return a.loadTasks()
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.help.Width = msg.Width
		return a, nil

	case tasksLoadedMsg:
		a.organizeTasks(msg.tasks)
		a.err = nil
		return a, nil

	case errorMsg:
		a.err = msg.err
		return a, nil
	}

	switch a.state {
	case viewKanban:
		return a.updateKanban(msg)
	case viewDetail:
		return a.updateDetail(msg)
	}

	return a, nil
}

func (a *App) updateKanban(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, a.keys.Quit):
			return a, tea.Quit
		case key.Matches(msg, a.keys.Up):
			if a.cursor.row > 0 {
				a.cursor.row--
			}
		case key.Matches(msg, a.keys.Down):
			col := &a.columns[a.cursor.col]
			if a.cursor.row < len(col.tasks)-1 {
				a.cursor.row++
			}
		case key.Matches(msg, a.keys.Left):
			if a.cursor.col > 0 {
				a.cursor.col--
				a.cursor.row = 0
			}
		case key.Matches(msg, a.keys.Right):
			if a.cursor.col < len(a.columns)-1 {
				a.cursor.col++
				a.cursor.row = 0
			}
		case key.Matches(msg, a.keys.Enter):
			col := &a.columns[a.cursor.col]
			if a.cursor.row < len(col.tasks) {
				a.selected = col.tasks[a.cursor.row]
				a.state = viewDetail
			}
		case key.Matches(msg, a.keys.NewTask):
			// TODO: open input dialog for new task
			return a, nil
		case key.Matches(msg, a.keys.Help):
			a.help.ShowAll = !a.help.ShowAll
		}
	}
	return a, nil
}

func (a *App) updateDetail(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, a.keys.Quit), key.Matches(msg, a.keys.Enter):
			a.state = viewKanban
			a.selected = nil
		}
	}
	return a, nil
}

func (a *App) View() string {
	if a.width == 0 {
		return "Loading..."
	}

	switch a.state {
	case viewKanban:
		return a.viewKanban()
	case viewDetail:
		return a.viewDetail()
	}
	return ""
}

func (a *App) viewKanban() string {
	header := titleStyle.Render(" Capture Kanban ")

	var columns []string
	for i, col := range a.columns {
		var sb strings.Builder
		sb.WriteString(columnTitleStyle.Render(fmt.Sprintf("%s (%d)", col.title, len(col.tasks))))
		sb.WriteString("\n")

		if len(col.tasks) == 0 {
			sb.WriteString(statusStyle.Render("  (empty)"))
		}
		for j, task := range col.tasks {
			card := a.renderCard(task, i == a.cursor.col && j == a.cursor.row)
			sb.WriteString(card)
			sb.WriteString("\n")
		}

		columns = append(columns, columnStyle.Render(sb.String()))
	}

	kanban := lipgloss.JoinHorizontal(lipgloss.Top, columns...)

	footer := a.helpView()

	if a.err != nil {
		footer = fmt.Sprintf("\nError: %v", a.err)
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, kanban, footer)
}

func (a *App) renderCard(task *model.Task, selected bool) string {
	style := cardStyle
	if selected {
		style = selectedCardStyle
	}

	var parts []string
	parts = append(parts, fmt.Sprintf("%s %s", task.ID, task.Title))
	parts = append(parts, fmt.Sprintf("stage: %s", task.Stage))

	priorityStr := string(task.Priority)
	switch task.Priority {
	case model.PriorityHigh:
		priorityStr = priorityHighStyle.Render(priorityStr)
	case model.PriorityMedium:
		priorityStr = priorityMediumStyle.Render(priorityStr)
	case model.PriorityLow:
		priorityStr = priorityLowStyle.Render(priorityStr)
	}
	parts = append(parts, priorityStr)

	if len(task.Tags) > 0 {
		tagStrs := make([]string, len(task.Tags))
		for i, t := range task.Tags {
			tagStrs[i] = tagStyle.Render("#" + t)
		}
		parts = append(parts, strings.Join(tagStrs, " "))
	}
	if task.Dispatch.Agent != "" {
		parts = append(parts, fmt.Sprintf("agent: %s", task.Dispatch.Agent))
	}

	return style.Render(strings.Join(parts, "\n"))
}

func (a *App) viewDetail() string {
	if a.selected == nil {
		return "No task selected"
	}

	t := a.selected
	var sb strings.Builder

	sb.WriteString(detailTitleStyle.Render(fmt.Sprintf("Task %s", t.ID)))
	sb.WriteString(fmt.Sprintf("\n\nTitle:    %s", t.Title))
	sb.WriteString(fmt.Sprintf("\nStatus:   %s", t.Status))
	sb.WriteString(fmt.Sprintf("\nStage:    %s", t.Stage))
	sb.WriteString(fmt.Sprintf("\nPriority: %s", t.Priority))
	sb.WriteString(fmt.Sprintf("\nSource:   %s", t.Source))
	sb.WriteString(fmt.Sprintf("\nCreated:  %s", t.CreatedAt.Format("2006-01-02 15:04:05")))
	sb.WriteString(fmt.Sprintf("\nUpdated:  %s", t.UpdatedAt.Format("2006-01-02 15:04:05")))
	if t.Dispatch.Agent != "" {
		sb.WriteString(fmt.Sprintf("\nAgent:    %s", t.Dispatch.Agent))
	}
	if t.Dispatch.Repository != "" {
		sb.WriteString(fmt.Sprintf("\nRepo:     %s", t.Dispatch.Repository))
	}

	if len(t.Tags) > 0 {
		sb.WriteString(fmt.Sprintf("\nTags:     %s", strings.Join(t.Tags, ", ")))
	}
	if t.Description != "" {
		sb.WriteString(fmt.Sprintf("\n\n%s", t.Description))
	}

	sb.WriteString(helpStyle.Render("\n\nPress Enter or Esc to go back"))

	return sb.String()
}

func (a *App) helpView() string {
	return helpStyle.Render(a.help.View(a.keys))
}

func (a *App) loadTasks() tea.Cmd {
	return func() tea.Msg {
		tasks, err := a.taskSvc.List(context.Background(), model.TaskFilter{})
		if err != nil {
			return errorMsg{err}
		}
		return tasksLoadedMsg{tasks}
	}
}

func (a *App) organizeTasks(tasks []*model.Task) {
	for i := range a.columns {
		a.columns[i].tasks = nil
	}

	for _, t := range tasks {
		var colIdx int
		switch t.Status {
		case model.StatusTodo:
			colIdx = 0
		case model.StatusInProgress:
			colIdx = 1
		case model.StatusDone:
			colIdx = 2
		default:
			continue
		}
		a.columns[colIdx].tasks = append(a.columns[colIdx].tasks, t)
	}
}

type tasksLoadedMsg struct {
	tasks []*model.Task
}

type errorMsg struct {
	err error
}
