package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestParseFloatLogsSuccess(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	handleFloatParsingErrors("1.2")
	logMessage := buffer.String()
	assert.Contains(t, logMessage, "successfully parsed float", "unexpected log message")
}

func TestParseFloatLogsSFailure(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	handleFloatParsingErrors("cat")
	logMessage := buffer.String()
	assert.Contains(t, logMessage, "error parsing floatstrconv.ParseFloat: parsing \"cat\": invalid syntax", "unexpected log message")
}
