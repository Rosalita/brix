package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Userinput holds the input given by the user
type Userinput struct {
	Dimension float64
	Length    float64
	Width     float64
	Height    float64
	Joint     float64
}

// Coordination holds a trio of CO, CO+ and CO- values
type Coordination struct {
	basic float64
	plus  float64
	minus float64
}

// horizontalResult stores the nearest horizontal coordinating sizes to the dimension
type horizontalResult struct {
	nfull  int
	fullCo float64
	nhalf  float64
	halfCo float64
}

// verticalResult stores the nearest vertical course sizes to the dimension
type verticalResult struct {
	nfirst   int
	firstCo  float64
	nsecond  int
	secondCo float64
}

// pageVariable are the values displayed by the html template
type pageVariables struct {
	PageTitle        string
	Dimension        float64
	Length           float64
	Width            float64
	Height           float64
	Joint            float64
	Result           bool
	WholeBricks      int
	Remainder        float64
	Nfull            int
	FullCo           float64
	FullCoPlus       float64
	FullCoMinus      float64
	Nhalf            float64
	HalfCo           float64
	HalfCoPlus       float64
	HalfCoMinus      float64
	VerticalResult   bool
	Courses          int
	CoursesRemainder float64
	NfirstVertical   int
	NsecondVertical  int
	FirstVCo         float64
	FirstVCoPlus     float64
	FirstVCoMinus    float64
	SecondVCo        float64
	SecondVCoPlus    float64
	SecondVCoMinus   float64
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
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
	result := calcHorizontalResult(remainder, wholeBricks, coSizeForHalfBrick, coSizeForFullBrick)

	fullCo, halfCo, firstVertCo, secondVertCo := Coordination{}, Coordination{}, Coordination{}, Coordination{}

	if result.nfull != 0 {
		fullCo = calcCoPlusAndMinus(result.fullCo, input.Joint)
	}

	if result.nhalf != 0 {
		halfCo = calcCoPlusAndMinus(result.halfCo, input.Joint)
	}

	verticalCoSize := calcCoSize(input.Height, input.Joint)
	courses := calcWholeBricksInDim(input.Dimension, verticalCoSize)
	coursesRemainder := calcRemainderFromDim(input.Dimension, verticalCoSize)
	verticalResult := calcVerticalResult(coursesRemainder, courses, verticalCoSize)

	if verticalResult.nfirst != 0 {
		firstVertCo = calcCoPlusAndMinus(verticalResult.firstCo, input.Joint)
	}

	if verticalResult.nsecond != 0 {
		secondVertCo = calcCoPlusAndMinus(verticalResult.secondCo, input.Joint)
	}

	fmt.Printf("%+v\n", verticalResult)

	pageVariables := pageVariables{
		PageTitle:        "Rosibrix v2.0",
		Dimension:        input.Dimension,
		Length:           input.Length,
		Width:            input.Width,
		Height:           input.Height,
		Joint:            input.Joint,
		Result:           true,
		WholeBricks:      wholeBricks,
		Remainder:        remainder,
		Nfull:            result.nfull,
		FullCo:           fullCo.basic,
		FullCoPlus:       fullCo.plus,
		FullCoMinus:      fullCo.minus,
		Nhalf:            result.nhalf,
		HalfCo:           halfCo.basic,
		HalfCoPlus:       halfCo.plus,
		HalfCoMinus:      halfCo.minus,
		VerticalResult:   true,
		Courses:          courses,
		CoursesRemainder: coursesRemainder,
		NfirstVertical:   verticalResult.nfirst,
		NsecondVertical:  verticalResult.nsecond,
		FirstVCo:         firstVertCo.basic,
		FirstVCoPlus:     firstVertCo.plus,
		FirstVCoMinus:    firstVertCo.minus,
		SecondVCo:        secondVertCo.basic,
		SecondVCoPlus:    secondVertCo.plus,
		SecondVCoMinus:   secondVertCo.minus,
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
	} 
	return f
}
