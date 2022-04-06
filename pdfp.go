package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"tjweldon/pdf-printer/url2pdf"
)

type Controller func(w http.ResponseWriter, req *http.Request)

func index(w http.ResponseWriter, req *http.Request) {

	if strings.Count(req.URL.String(), "/") != 3 || !strings.HasSuffix(req.URL.String(), "/") {
		w.WriteHeader(404)
		return
	}
	switch req.Method {
	case "GET":
		log.Printf("Invalid request method: %v, Header: %v", req.Method, req.Header)
		w.WriteHeader(422)
	case "POST":
		defer func() {
			if err := os.Remove("./tmp.pdf"); err != nil {
				log.Print(err)
			}
		}()
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		var (
			path string
			err  error
		)
		if path, err = url2pdf.Url2PDF(req.FormValue("url")); err != nil {
			w.WriteHeader(500)
			log.Println(err)
		}
		http.ServeFile(w, req, path)
	}
}

func main() {
	var (
		logfile *os.File
		err     error
	)
	if logfile, err = initLogfile(); err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(logfile)

	http.HandleFunc("/", wrapLogs(index))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initLogfile() (logfile *os.File, err error) {
	logfile = os.Stdout
	if _, ok := os.LookupEnv("STDOUT_LOG"); !ok {
		logfile, err = os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		log.SetOutput(logfile)
	}

	return logfile, nil
}

func wrapLogs(c Controller) Controller {
	wrapped := func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received: %v", req.Header)
		c(w, req)
		log.Printf("Response Headers: %v", w.Header())
	}

	return wrapped
}
