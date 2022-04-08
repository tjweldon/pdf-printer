package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"tjweldon/pdf-printer/url2pdf"
)

const tempPath = "/tmp/tmp.pdf"

type Controller func(w http.ResponseWriter, req *http.Request)
type Middleware func(c Controller) Controller

type MiddlewareStack struct {
	middlewares []Middleware
}

func DeclareMiddleware(mStack ...Middleware) MiddlewareStack {
	return MiddlewareStack{middlewares: mStack}
}

func (m MiddlewareStack) Decorate(c Controller) (wrapped Controller) {
	wrapped = c
	for _, mid := range m.middlewares {
		wrapped = mid(wrapped)
	}

	return wrapped
}

var middlewares = DeclareMiddleware(
	wrapLogs,
)

func index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		log.Printf("Invalid request method: %v, Header: %v", req.Method, req.Header)
		w.WriteHeader(422)
	case "POST":
		defer func() {
			if err := os.Remove(tempPath); err != nil {
				log.Printf("Error removing temp pdf: %s", err)
			}
		}()
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(500)
			log.Printf("Error parsing formdata: %s", err)
			return
		}

		target := req.FormValue("url")
		log.Printf("Request for print of page at %s to pdf received", target)
		if err := url2pdf.Url2PDF(target, tempPath); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			_, _ = fmt.Fprintf(w, "Internal Server Error: %s", err)
			return
		}
		http.ServeFile(w, req, tempPath)
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

	http.HandleFunc("/print", middlewares.Decorate(index))

	log.Fatal(http.ListenAndServe(":"+initPort(), nil))
}

func initPort() (port string) {
	port = "80"
	if len(os.Args) > 1 {
		userPort := os.Args[1]
		if _, err := strconv.Atoi(userPort); err == nil {
			port = userPort
		}
	}

	return port
}

func initLogfile() (logfile *os.File, err error) {
	logfile = os.Stdout
	if logfilePath, ok := os.LookupEnv("PDFP_LOGFILE"); ok {
		logfile, err = os.OpenFile(logfilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
	}
	log.SetOutput(logfile)

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
