package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	// Construct full path to the template
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Example data you might pass to the template
	data := map[string]string{
		"Title":   "Go HTMX Tailwind Starter",
		"Message": "Go HTMX Tailwind Starter",
	}
	renderTemplate(w, "index.html", data)
}

func htmxHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "htmx.html", nil)
}

func main() {
	// Serve static files from the "static" directory at "/static/" URL path
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("GET /htmx", htmxHandler)

	log.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
