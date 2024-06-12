package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

type Response struct {
	Name string
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "templates/index.html")
			// t, err := template.new("")
		}else if r.Method == http.MethodPost{
			r.ParseForm()
			name := r.FormValue("user_input")

			data := Response {
				Name: name,
			}

			tmpl, err := template.ParseFiles("templates/index.html")

			if err != nil {
				tmpl.Execute(w, err)
			}

			tmpl.Execute(w, data)
		}

	})

	fs := http.FileServer(http.Dir("static/style"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
