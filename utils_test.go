package main

import (
	"log"
	"bytes"
	"net/http/httptest"
	"os"
	"testing"
  "text/template"
	"github.com/stretchr/testify/assert"
)

func TestCanGetDefaultPort(t *testing.T) {
	expected := ":8080"
	actual := getPort()
	assert.Equal(t, expected, actual, "unexpected port returned")
}

func TestCanGetPortFromEnvironmentVariable(t *testing.T) {
	os.Setenv("PORT", "1234")
	actual := getPort()
	assert.Equal(t, ":1234", actual, "unexpected port returned")
}

func TestCanRenderPageTemplates(t *testing.T) {
	expectedBody := `<!DOCTYPE html>
<html>
<head>
<title>test title</title>
</head>
<body>
</body>
</html>`

	templatePath := "testdata/validTemplate.html"
	pageVariables := pageVariables{PageTitle: "test title"}
	responseRecorder := httptest.NewRecorder()

	renderPage(responseRecorder, templatePath, pageVariables)
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, expectedBody, responseRecorder.Body.String())
}

func TestParseTemplateLogsFailure(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	_ = parseTemplate("testdata/invalidTemplate.html")
	logMessage := buffer.String()
	assert.Contains(t, logMessage, "template parsing error: template: invalidTemplate.html:4: unexpected \"}\" in operand\n", "unexpected log message")
}

func TestParseTemplateLogsSuccess(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	_ = parseTemplate("testdata/validTemplate.html")
	logMessage := buffer.String()
	assert.Contains(t, logMessage, "successfully parsed template", "unexpected log message")
}
func TestExecuteTemplateLogsFailure(t *testing.T){
	templ := template.New("test")
	vars := pageVariables{PageTitle: "test title"}
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	responseRecorder := httptest.NewRecorder()
	executeTemplate(templ, responseRecorder, vars)
	logMessage := buffer.String()
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Contains(t, logMessage, "template executing error template: test: \"test\" is an incomplete or empty template", "unexpected log message")
}

func TestExecuteTemplateLogsSuccess(t *testing.T){
	templ := parseTemplate("testdata/validTemplate.html")
	vars := pageVariables{PageTitle: "test title"}
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	responseRecorder := httptest.NewRecorder()
	executeTemplate(templ, responseRecorder, vars)
	logMessage := buffer.String()
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Contains(t, logMessage, "successfully executed template", "unexpected log message")
	
}