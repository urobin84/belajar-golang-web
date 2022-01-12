package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request)  {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title" : "Template Auto Escape",
		"Body" : "<p>Ini Adalah Body</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1700",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	PanicError(err)
}

func TemplateAutoEscapeDisables(writer http.ResponseWriter, request *http.Request)  {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title" : "Template Auto Escape",
		"Body" : template.HTML("<p>Ini Adalah Body</p>"),
	})
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1700",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisables),
	}

	err := server.ListenAndServe()
	PanicError(err)
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request)  {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title" : "Template Auto Escape",
		"Body" : template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1700",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	PanicError(err)
}
