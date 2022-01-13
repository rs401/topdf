package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs401/topdf/converter"
)

var PORT string

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
		return
	}
	defer output.Close()
	defer os.Remove(output.Name())

	err = converter.Convtopdf(input.Name(), output.Name())
	if err != nil {
		log.Println("Error Creating PDF File")
		log.Printf("Error: %s\n", err.Error())
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}

	out, err := os.Open(output.Name())
	if err != nil {
		log.Println("Error Opening PDF File for transfer.")
		log.Println(err)
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}
	defer out.Close()
	io.Copy(w, out)
	log.Printf("Successfully Converted File\n")
}

func setupRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/topdf", convertFile).Methods("POST")
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
