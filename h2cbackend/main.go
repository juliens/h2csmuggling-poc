package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	h2s := &http2.Server{}

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    log.Println("/ called")
		fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	})

	handler.HandleFunc("/flag", func(w http.ResponseWriter, r *http.Request) {
        log.Println("/flag called")
		fmt.Fprintf(w, "You got the flag!")
    })

	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: h2c.NewHandler(handler, h2s),
	}

	fmt.Printf("Listening [0.0.0.0:80]...\n")
	checkErr(server.ListenAndServe(), "while listening")
}
