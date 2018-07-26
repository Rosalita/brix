package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"net/url"
)

func TestParseFloatLogsSFailure(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	handleFloatParsingErrors("cat")
	logMessage := buffer.String()
	assert.Contains(t, logMessage, "error parsing floatstrconv.ParseFloat: parsing \"cat\": invalid syntax", "unexpected log message")
}

func TestCanMarshallFormValues(t *testing.T){
	v := url.Values{}
	v.Set("length", "215")
	v.Set("width", "102.5")
	v.Set("height", "65")
	v.Set("joint", "10")
	v.Set("dimension", "3400")

	input := marshalFormValues(v)
	assert.Equal(t, float64(215), input.Length, "unexpected value for length marshalled from form")
	assert.Equal(t, float64(102.5), input.Width, "unexpected value for width marshalled from form")
	assert.Equal(t, float64(65), input.Height, "unexpected value for height marshalled from form")
	assert.Equal(t, float64(10), input.Joint, "unexpected value for joint marshalled from form")
	assert.Equal(t, float64(3400), input.Dimension, "unexpected value for dimension marshalled from form")
}