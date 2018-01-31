package main

import (
	"html/template"
	"net/http"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/index/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}
