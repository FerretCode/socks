package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/list"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type command struct {
	title, desc string
}

func (c command) Title() string       { return c.title }
func (c command) Description() string { return c.desc }
func (c command) FilterValue() string { return c.title }

type Model struct {
	list list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
} 

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "q":
					return m, tea.Quit
			}

			case tea.WindowSizeMsg:
				h, v := docStyle.GetFrameSize()
				m.list.SetSize(msg.Width-h, msg.Height-v)

		default:
			var cmd tea.Cmd
			return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	s := ""

	s += docStyle.Render(m.list.View())
	
	m.list.Select(m.list.Cursor())
	
	s += m.list.SelectedItem().FilterValue()

	return s
}

func main() {
	commands := []list.Item{
		command{
			title: "Get Courses", 
			desc: "Fetch all courses from Canvas",
		},
		command{
			title: "Get Assignments for Course",
			desc: "Fetch all assignments for a given course",
		},
	} 

	m := Model{list: list.New(commands, list.NewDefaultDelegate(), 0, 0),}
	m.list.Title = "Commands"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Printf("There was an error: %v", err)
		os.Exit(1)
	}
}
