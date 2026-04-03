package tui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2)

	columnStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			Width(30)

	columnTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#7D56F4"))

	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#585858")).
			Padding(0, 1).
			Margin(0, 0, 1, 0)

	selectedCardStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#FF6B6B")).
				Padding(0, 1).
				Margin(0, 0, 1, 0)

	priorityHighStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF6B6B")).Bold(true)

	priorityMediumStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFD93D"))

	priorityLowStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6BCB77"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	tagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4D96FF"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Padding(1, 0)

	detailTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#7D56F4")).
				Margin(0, 0, 1, 0)
)
