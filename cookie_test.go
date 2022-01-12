package belajar_golang_web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request)  {
	cookie := new(http.Cookie)
	cookie.Name = "X-MQR-NAME"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")

}

func GetCookie(w http.ResponseWriter, r *http.Request)  {
	cookie, err := r.Cookie("X-MQR-NAME")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:1600",
		Handler: mux,
	}

	err:= server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700/?name=Robin", nil)
	recorder := httptest.NewRecorder()
	
	SetCookie(recorder, request)
	
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1700/?name=Robin", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}
