package request

import (
	view "example.com/socks/views"
)

type Assignment struct {
	Id       uint64 `json:"id"`
	CourseId uint64 `json:"course_id"`
	Name     string `json:"name"`
	DueDate  string `json:"due_at"`
	URL      string `json:"html_url"`
}

type Assignments struct {
	Assignments []Assignment
	View        view.TableModel
	Error       error
}

func GetAssignments(config Config) (Assignments, error) {
	return Assignments{}, nil
}
