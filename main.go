package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var host string
var content string
var contentFile string
var contentType string
var statusCode int

func main() {
	flag.StringVar(&host, "a", ":8080", "server address, format: '<hostname|ip-address>:<port>'")
	flag.StringVar(&content, "c", "", "content to return to the client; given as string value on the command line")
	flag.StringVar(&contentFile, "cf", "", "filename of the content to return to the client")
	flag.StringVar(&contentType, "ct", "text/plain", "content type of the response")
	flag.IntVar(&statusCode, "s", 200, "status code of the response")

	flag.Parse()

	if content != "" && contentFile != "" {
		panic("Must use only one of 'c' or 'cf'")
	}

	fmt.Printf(">> Listening on %s\n", host)

	if contentFile != "" {
		fmt.Printf(">> Serving content from file: %s, Content-Type: %s\n", contentFile, contentType)
		fileContent, err := os.ReadFile(contentFile)
		if err != nil {
			panic("Cannot load file contents")
		}
		content = string(fileContent)
	} else {
		fmt.Printf(">> Serving content: %s, Content-Type: %s\n", content, contentType)
	}

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">> Request URL: %s | Timestamp: %s | Method: %s\n%s\n",
		r.URL,
		time.Now().Format("2006-01-02 15:04:05"),
		r.Method,
		string(dump),
	)

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", contentType)
	w.Write([]byte(content))
}
