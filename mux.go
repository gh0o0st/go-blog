package main

import (
	"html/template"
	"log"
	"net/http"
)

type URLs map[string]http.Handler

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))
	blog := Blog{}
	tmpl.Execute(w, blog)
}

//Mux is a ServeMux with custom method
type Mux struct {
	http.ServeMux
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	m.ServeMux.ServeHTTP(w, r)
}

func NewMux() *Mux {
	return &Mux{*http.NewServeMux()}
}

func main() {
	mux := NewMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)
}
