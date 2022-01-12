package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1700",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name != "" {
		fmt.Fprintf(writer, resourceOk)
	} else {
		fmt.Fprintf(writer, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1700",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

