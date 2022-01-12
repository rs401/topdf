package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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
	log.Printf("MIME Header: %+v\n", header.Header)

	f, err := os.OpenFile("./downloaded", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error Creating the Temp File")
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	err = convtopdf(f.Name(), "./out.pdf")
	if err != nil {
		log.Println("Error Creating PDF File")
		log.Printf("Error: %s\n", err.Error())
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}

	out, err := os.Open("./out.pdf")
	if err != nil {
		log.Println("Error Creating PDF File")
		log.Println(err)
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}
	defer out.Close()
	io.Copy(w, out)
	log.Printf("Successfully Converted File\n")
}

func setupRoutes() error {
	http.HandleFunc("/topdf", convertFile)
	return http.ListenAndServe(":8888", nil)
}

func main() {
	log.Fatal(setupRoutes())
}
