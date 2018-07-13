package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type pageVariables struct {
	PageTitle string
}

type Userinput struct {
	Dimension float64
	Length    float64
	Width     float64
	Height    float64
	Joint     float64
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", handler)
	http.HandleFunc("/result", showResult)
	fmt.Println("listening and serving requests..")
	http.ListenAndServe(getPort(), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	pageVariables := pageVariables{}
	renderPage(w, "brix.html", pageVariables)
}

func showResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var input Userinput
	for i, value := range r.Form {
		switch i {
		case "length":
			input.Length = handleFloatParsingErrors(value[0])
		case "width":
			input.Width = handleFloatParsingErrors(value[0])
		case "height":
			input.Height = handleFloatParsingErrors(value[0])
		case "joint":
			input.Joint = handleFloatParsingErrors(value[0])
		case "dimension":
			input.Dimension = handleFloatParsingErrors(value[0])
		}
	}
}

func handleFloatParsingErrors(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Print("error parsing float", err)
		return 0
	} else {
		log.Printf("successfully parsed float %f", f)
		return f
	}
}
