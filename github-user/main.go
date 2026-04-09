package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`

	Actor Actor `json:"actor"`
	Repo  Repo  `json:"repo"`

	Payload map[string]interface{} `json:"payload"`
}

type Actor struct {
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	Url          string `json:"url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// outsite of a function can't use := to create a variable
var EventTypes = make(map[string]int)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <username>")
		return
	}
	username := os.Args[1]

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Create request
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	var events []Event

	err = json.NewDecoder(req.Body).Decode(&events)
	if err != nil {
		panic(err)
	}

	for _, e := range events {

		_, ok := EventTypes[e.Type]
		if ok {
			// Key exists, use value
			EventTypes[e.Type] += 1
		} else {
			// Key does not exist
			EventTypes[e.Type] = 1
		}
	}

	fmt.Printf("Events done by %s:\n", username)
	for eventT, times := range EventTypes {

		fmt.Printf("%s = %d\n", eventT, times)

	}

}

func separator(times int, sep string) {
	for i := 0; i < times; i++ {
		fmt.Print(sep)
	}

}
