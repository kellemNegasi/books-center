package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("page not found!"))
		return
	}
	if r.Method == "GET" {
		fmt.Fprint(w, "welcome to the index of this page.")
	} else {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5060"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	server := &http.Server{Addr: ":" + port, Handler: mux}
	fmt.Printf("started listening on %s", port)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("server failed %s", err.Error())
	}
}
