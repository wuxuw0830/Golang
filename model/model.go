package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Title   string
	Content string
}

func test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./home.html"))
	data := new(IndexData)
	data.Title = "Home Page"
	data.Content = "First web application"
	tmpl.Execute(w, data)

}
func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/home", test)
	err := http.ListenAndServe(":3002", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
