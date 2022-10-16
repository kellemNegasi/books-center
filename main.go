package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	HOST = "postgres"
	PORT = 5432
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
	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to DB %s \n", err)
	}
	defer db.Close()
	err = db.Ping()
	start := time.Now()
	for db.Ping() != nil {
		if start.After(start.Add(20 * time.Second)) {
			fmt.Printf("Failed to connect after 10 seconds %s\n", err.Error())
			break
		}
	}

	fmt.Println("Successfully connected!", db.Ping() == nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	server := &http.Server{Addr: ":" + port, Handler: mux}
	fmt.Printf("started listening on %s", port)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("server failed %s", err.Error())
	}
}

func connectDB() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable", HOST, PORT, user, password, dbname)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
