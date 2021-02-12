package main

import (
	"html/template"
	"log"
	"net/http"
)

type TodoList struct {
	Title string
	Items []Todo
}

type Todo struct {
	Done bool
	Job  string
}

func index(w http.ResponseWriter, r *http.Request) {
	data := TodoList{
		Title: "Todo List of Today",
		Items: []Todo{
			{Job: "Go to school", Done: true},
			{Job: "sleep", Done: false},
		},
	}
	templ := template.Must(template.ParseFiles("web/index.html"))
	templ.Execute(w, data)

	//fmt.Printf("hello world",w)
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
