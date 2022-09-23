package request

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Course struct {
	Id uint64 `json:"id"`
}

type Courses struct {
	Courses []Course	
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

	return courses, nil
}
