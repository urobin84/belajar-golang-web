package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:1700",
		Handler: handler,
	}

	err := server.ListenAndServe()
	PanicError(err)
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/salam", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Assalamualaikum")
	})

	server := http.Server{
		Addr: "localhost:1700",
		Handler: mux,
	}

	err := server.ListenAndServe()
	PanicError(err)
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:1700",
		Handler: handler,
	}

	err := server.ListenAndServe()
	PanicError(err)
}
