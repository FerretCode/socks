package request

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/socks/views"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type Course struct {
	Id uint64 `json:"id"`
	Name string `json:"name"` 
}

type Courses struct {
	Courses []Course	
	View view.TableModel
	Error error
}

func GetCourses(config Config) (Courses, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://%s.instructure.com/api/v1/courses", config.Domain),
		nil,
	)

	if err != nil {
		return Courses{Error: err}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))

	res, err := client.Do(req)

	if err != nil {
		return Courses{Error: err}, err
	}

	parser := ResponseParser{
		Response: *res,
	}

	body, err := parser.ParseRequest(*res)

	if err != nil {
		return Courses{Error: err}, err
	}

	courses := Courses{}

	if jsonErr := json.Unmarshal(body, &courses.Courses); err != nil {
		log.Fatal(jsonErr)
	}

	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Name", Width: 50},
	}	

	rows := []table.Row{}

	for _, c := range courses.Courses {
		row := table.Row{fmt.Sprint(c.Id), fmt.Sprint(c.Name)}

		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	courses.View = view.TableModel{
		Table: t,
		Columns: columns,
		Row: rows,
	}

	return courses, nil
}
