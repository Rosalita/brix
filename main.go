package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Userinput struct {
	Dimension float64
	Length    float64
	Width     float64
	Height    float64
	Joint     float64
}

type Cos struct {
	co      float64
	coPlus  float64
	coMinus float64
}

type result struct {
	nfull int
	fullCo float64
	nhalf float64
	halfCo float64
}

type pageVariables struct {
	PageTitle   string
	Dimension   float64
	Length      float64
	Width       float64
	Height      float64
	Joint       float64
	Result			bool
	WholeBricks int
	Remainder   float64
	Nfull       int
	FullCo      float64
	FullCoPlus  float64
	FullCoMinus float64
	Nhalf       float64
	HalfCo      float64
	HalfCoPlus  float64
	HalfCoMinus float64
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/result", handleResult)
	fmt.Println("listening and serving requests..")
	http.ListenAndServe(getPort(), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	pageVariables := pageVariables{}
	renderPage(w, "brix.html", pageVariables)
}

func handleResult(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	input := marshalFormValues(r.Form)
	coSizeForFullBrick := calcCoSize(input.Length, input.Joint)
	coSizeForHalfBrick := calcCoSize(calcHalfBrickSize(input.Length, input.Joint), input.Joint)
	wholeBricks := calcWholeBricksInDim(input.Dimension, coSizeForFullBrick)
	remainder := calcRemainderFromDim(input.Dimension, coSizeForFullBrick)

	result := calcResult(remainder, wholeBricks, coSizeForHalfBrick, coSizeForFullBrick)

	fullCos := calcCoPlusAndMinus(result.fullCo, input.Joint)
	halfCos := calcCoPlusAndMinus(result.halfCo, input.Joint)

	pageVariables := pageVariables{
		PageTitle:   "Rosibrix v2.0",
		Dimension:   input.Dimension,
		Length:      input.Length,
		Width:       input.Width,
		Height:      input.Height,
		Joint:       input.Joint,
		Result:      true,
		WholeBricks: wholeBricks,
		Remainder:   remainder,
		Nfull:        result.nfull,
		FullCo:      fullCos.co,
		FullCoPlus:  fullCos.coPlus,
		FullCoMinus: fullCos.coMinus,
		Nhalf:       result.nhalf,
		HalfCo:      halfCos.co,
		HalfCoPlus:  halfCos.coPlus,
		HalfCoMinus: halfCos.coMinus,
	}

	renderPage(w, "brix.html", pageVariables)
}

func marshalFormValues(values url.Values) Userinput {
	var input Userinput
	for key, value := range values {
		switch key {
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
	return input
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
