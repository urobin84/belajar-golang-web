package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1700",
	}

	err := server.ListenAndServe()
	PanicError(err)
}

func PanicError(err error)  {
	if err != nil {
		panic(err)
	}
}