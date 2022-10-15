package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	if err != nil {
		log.Fatalf("failed to ping %s\n", err)
	}

	fmt.Println("Successfully connected!")

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
	user := os.Getenv("PSQUSER")
	password := os.Getenv("PSQLPASSWORD")
	dbname := os.Getenv("DBNAME")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
