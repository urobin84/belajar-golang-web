package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request){
	err := myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
	PanicError(err)
}

func Upload(writer http.ResponseWriter, request *http.Request){
	err := request.ParseMultipartForm(100 << 20)
	PanicError(err)
	file, fileHeader, err := request.FormFile("file")
	PanicError(err)
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	PanicError(err)
	_, err = io.Copy(fileDestination, file)
	PanicError(err)

	name := request.PostFormValue("name")
	err = myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
	PanicError(err)
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:1700",
		Handler: mux,
	}

	err := server.ListenAndServe()
	PanicError(err)
}

//go:embed resources/1563799935075.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Muhammad Muqorrobin")
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.jpg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:1700/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
