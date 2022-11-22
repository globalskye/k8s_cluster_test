package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello 2"))
	})
	server := http.Server{Addr: ":1111", Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
