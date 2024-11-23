package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var data map[string]interface{}

type Author struct {
	Name string
	Bio  string
}
type Post struct {
	Title   string
	Content string
	Author  Author
	Tags    []string
}

func main() {
	author := &Author{
		"Tom",
		"My name is Tom",
	}
	post := &Post{
		"My first post",
		"This is my first post",
		*author,
		[]string{"Newbie", "Positive"},
	}
	jsonData, marshalErr := json.Marshal(post)
	if marshalErr != nil {
		fmt.Println(marshalErr)
	}
	parts := strings.Split(string(jsonData)[1:len(jsonData)-1], ",")
	for _, part := range parts {
		fmt.Println(part)
	}
	start := strings.Index(string(jsonData), "Author")
	end := strings.Index(string(jsonData)[start:], "}") + start
	authorString := "{" + string(jsonData)[start-1:end+1] + "}"
	var jsonAuthor Author
	unmarshalErr := json.Unmarshal([]byte(authorString), &jsonAuthor)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}
	fmt.Printf("Name: %s\nBio: %s\n", author.Name, author.Bio)
}
