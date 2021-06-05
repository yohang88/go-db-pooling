package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *sql.DB

func main() {
	dbConnection := os.Getenv("DB_CONNECTION")

	if dbConnection != "" {
		db = connect()
	}

	http.HandleFunc("/", health)
	http.HandleFunc("/db", hello)

	log.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hostname, _ := os.Hostname()

	m := make(map[string]interface{})
	m["server"] = hostname
    resp, _ := json.Marshal(m)
	fmt.Fprintf(w, string(resp))
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query(`SELECT SLEEP(2);`)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error.")
		return
	}

	defer rows.Close()

	hostname, _ := os.Hostname()
	m := make(map[string]interface{})
	m["server"] = hostname
	m["db"] = "OK"

	resp, _ := json.Marshal(m)
	fmt.Fprintf(w, string(resp))
}

func connect() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbMaxConnections := os.Getenv("DB_CONNECTION_MAX")
	maxConnection, _ := strconv.Atoi(dbMaxConnections)

	log.Printf("DB Host: %s\n", dbHost)
	log.Printf("DB Port: %s\n", dbPort)
	log.Printf("DB Database: %s\n", dbDatabase)
	log.Printf("DB Username: %s\n", dbUsername)
	log.Printf("DB Max Connection: %d\n", maxConnection)

	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase)
	db.SetMaxOpenConns(maxConnection)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected.")

	return db
}