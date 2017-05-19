package main

import (
	"html/template"
	"net/http"
	"fmt"
	"path"
	"log"
)

func main() {
	fmt.Println("Starting server:")
	// Include folder
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	// Get that index
	http.HandleFunc("/", index)
	// ICON handler (So it does not run twice)
	http.HandleFunc("/favicon.ico", faviconHandler)
	// Listen on port
	err := http.ListenAndServe(":8090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	// Generate page and include web folder and index.html
	index := path.Join("web", "index.html")
	tmpl, start := template.ParseFiles(index)

	// Start the server with loaded data
	if start := tmpl.Execute(w, tmpl); start != nil {
		http.Error(w, start.Error(), http.StatusInternalServerError)
	}
	// Handle error
	err := r.ParseForm()
	if err != nil {
		http.Error(w, start.Error(), http.StatusInternalServerError)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/favicon.ico")
}
