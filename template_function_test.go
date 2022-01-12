package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name +", My Name is " + myPage.Name
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Robin",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request)  {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"sayHello" : func(value string) string{
			return "Hello " + value
		},
		"upper" : func(value string) string{
			return strings.ToUpper(value)
		},
	})
	t  = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Robin",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
