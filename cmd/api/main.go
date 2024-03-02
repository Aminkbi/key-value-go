package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello net/http!\n"))
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", helloGoHandler)

	srv := http.Server{
		Handler:      mux,
		Addr:         ":4000",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Printf("starting %s server on port %d", "develop", 4000)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
