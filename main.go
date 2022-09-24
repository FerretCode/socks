package main

import (
	"fmt"
	"log"
	"os"

	"example.com/socks/requests"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	text string
}

func (m Model) Init() tea.Cmd {
	return nil
} 

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
		cmds []tea.Cmd
	)

	config := request.Config{}

	token, exists := os.LookupEnv("SOCKS_TOKEN")

	if !exists {
		log.Fatal("SOCKS_TOKEN does not exist as an environment variable!")
	}

	domain, exists := os.LookupEnv("SOCKS_DOMAIN")

	if !exists {
		log.Fatal("SOCKS_DOMAIN does not exist as an environment variable!")
	}

	config.Token = token
	config.Domain = domain

	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "q":
					return m, tea.Quit

				case "enter", " ":
					switch m.list.SelectedItem().FilterValue() {
					case "Get Courses":
						courses, err := request.GetCourses(config)

						if err != nil {
							log.Fatal(err)
						}

						m.text = courses.View.View()
					}

					return m, nil
			}

			case tea.WindowSizeMsg:
				h, v := docStyle.GetFrameSize()
				m.list.SetSize(msg.Width-h, msg.Height-v)

		default:
			var cmd tea.Cmd
			return m, cmd
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := ""

	s += docStyle.Render(m.list.View())
	
	m.list.Select(m.list.Cursor())

	if m.text != "" {
		s = m.text
	}

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
