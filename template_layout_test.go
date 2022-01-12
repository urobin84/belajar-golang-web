package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))
	t.ExecuteTemplate(writer, "layout.gohtml", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Robin",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//input 5, 4, 8, 9, 7, 3, 2
//result 5 > 4 < 9 > 7 < 8 > 2 < 3 > 1 < 6
func TestKuisZigZag(t *testing.T) {
	angka := []int{5, 4, 7, 9, 8, 3, 2, 1, 6}
	start := 0
	end := len(angka) - 1

	for i := 0; i < len(angka); i++ {
		fmt.Print(angka[i], ",")
	}
	fmt.Println(" ")

	for i := 0; i < len(angka); i++ {
		if i%2 == 0 {
			if i == start {
				if angka[i] < angka[1] {
					change(i, 1, angka)
				}
			} else if i == end {
			} else {
				if angka[i] < angka[i+1] {
					change(i, i+1, angka)
				}
			}
		} else {
			if i == start {
				if angka[i] > angka[1] {
					change(i, 1, angka)
				}
			} else if i == end {
			} else {
				if angka[i] > angka[i+1] {
					change(i, i+1, angka)
				}
			}
		}

	}

	for i := 0; i < len(angka); i++ {
		if i == end {
			fmt.Print(angka[i])
		} else {
			if i%2 == 0 {
				fmt.Print(angka[i], " > ")
			} else {
				fmt.Print(angka[i], " < ")
			}
		}
	}
}

func change(a int, b int, data []int) {
	pv := data[a]
	data[a] = data[b]
	data[b] = pv
}
