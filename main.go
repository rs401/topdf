package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs401/topdf/converter"
)

var PORT string

const form = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Convert file to PDF</title>
    <style>
        body {
            padding-top: 4em;
            text-align: center;
        }
        form {
            display: inline-block;
            box-sizing: border-box;
        }
        input {
            margin: .5em;
            width: 100%;
        }
    </style>
</head>
<body>
    <div>
        <form action="/topdf" method="post" enctype="multipart/form-data">
          <label for="file">Select file to convert:</label>
            <input type="file" name="file" id="file">
            <input type="submit" value="Convert File" name="submit">
        </form>
      </div>
</body>
</html>`

func init() {
	PORT = config("PORT")
}

func config(key string) string {
	return os.Getenv(key)
}

func convertFile(w http.ResponseWriter, r *http.Request) {
	// 10MB
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, "Error Retrieving the File")
		log.Println(err)
		return
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", header.Filename)
	log.Printf("File Size: %+v\n", header.Size)
	log.Printf("MIME Header: %+v\n\n", header.Header)

	input, err := os.CreateTemp("./", "input")
	if err != nil {
		log.Println("Error Creating the Temp Input File")
		log.Println(err)
		fmt.Fprintln(w, "Error: Internal Server Error")
		return
	}
	defer input.Close()
	io.Copy(input, file)
	defer os.Remove(input.Name())

	// output
	output, err := os.CreateTemp("./", "output-*.pdf")
	if err != nil {
		log.Println("Error Creating the Temp Output File")
		log.Println(err)
		fmt.Fprintln(w, "Error: Internal Server Error")
		return
	}
	defer output.Close()
	defer os.Remove(output.Name())

	err = converter.Convtopdf(input.Name(), output.Name())
	if err != nil {
		log.Println("Error Creating PDF File")
		log.Printf("Error: %s\n", err.Error())
		fmt.Fprintln(w, "Error: Unable to convert file.\n")
		return
	}

	out, err := os.Open(output.Name())
	if err != nil {
		log.Println("Error Opening PDF File for transfer.")
		log.Println(err)
		fmt.Fprintln(w, "Error: Internal Server Error")
		return
	}
	defer out.Close()
	buffer := make([]byte, 10<<20)
	out.Read(buffer)
	contentType := http.DetectContentType(buffer)
	stat, err := out.Stat()
	if err != nil {
		log.Println("Error Stating PDF File.")
		log.Println(err)
		fmt.Fprintln(w, "Error: Internal Server Error")
		return
	}
	fileSize := strconv.FormatInt(stat.Size(), 10)
	w.Header().Set("Content-Type", contentType+";"+out.Name())
	w.Header().Set("Content-Length", fileSize)
	out.Seek(0, 0)
	io.Copy(w, out)
	log.Printf("Successfully Converted File\n\n")
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("form").Parse(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setupRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/topdf", convertFile).Methods("POST")
	r.HandleFunc("/topdf", uploadFile).Methods("GET")
	return r, nil
}

func main() {
	router, err := setupRoutes()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port: %s...\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
