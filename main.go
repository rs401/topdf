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

type MsgData struct {
	Messages []string
}

type GetHandler struct {
	Msgs MsgData
}

var (
	PORT string
	GH   *GetHandler
)

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
        ul {
            list-style-type: none;
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
		<ul>
			{{ range .Messages }}
			<li>💥 {{.}} 🔥</li>
			{{ end }}
		</ul>
    </div>
</body>
</html>`

func init() {
	PORT = config("PORT")
	GH = &GetHandler{Msgs: MsgData{Messages: []string{}}}
}

func config(key string) string {
	return os.Getenv(key)
}

func (gh *GetHandler) convertFile(w http.ResponseWriter, r *http.Request) {
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

	if header.Size > (10 << 20) {
		// Post notification about size
		log.Println("File too large.")
		GH.Msgs.Messages = append(GH.Msgs.Messages, "File too large. Must be less then 10MB.")
		// http.Redirect(w, r, "/topdf", http.StatusTemporaryRedirect)
		t, err := template.New("form").Parse(form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusRequestEntityTooLarge)

		if err := t.Execute(w, gh.Msgs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

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

	pdfFile, err := converter.Convtopdf(input.Name())
	if err != nil {
		log.Println("Error Creating PDF File")
		log.Printf("Error: %s\n", err.Error())
		fmt.Fprintln(w, "Error: Unable to convert file.")
		return
	}

	out, err := os.Open(pdfFile)
	if err != nil {
		log.Println("Error Opening PDF File for transfer.")
		log.Println(err)
		fmt.Fprintln(w, "Error: Internal Server Error")
		return
	}
	defer out.Close()
	defer os.Remove(out.Name())

	buffer := make([]byte, 512)
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

func (gh *GetHandler) uploadFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("form").Parse(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, gh.Msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setupRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/topdf", GH.convertFile).Methods("POST")
	r.HandleFunc("/topdf", GH.uploadFile).Methods("GET")
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
