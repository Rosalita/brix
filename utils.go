package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func renderPage(w http.ResponseWriter, filename string, vars pageVariables) {
	template := parseTemplate(filename)
	executeTemplate(template, w, vars)
}

func parseTemplate(filename string) (*template.Template) {
	template, err := template.ParseFiles(filename)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	log.Printf("successfully parsed template %s", filename)
	return template
}

func executeTemplate(t *template.Template, w http.ResponseWriter, v pageVariables){
	err := t.Execute(w, v)
	if err != nil {
		log.Print("template executing error ", err)
	} else{
		log.Printf("successfully executed template")
	}
}
