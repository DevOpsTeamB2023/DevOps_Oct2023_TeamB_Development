package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Account struct {
	AccID     int    `json:"accId"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	AccType   string `json:"accType"`
	AccStatus string `json:"accStatus"`
}

var (
	db  *sql.DB
	err error
)

func dB() {
	db, err = sql.Open("mysql", "record_system:dopasgpwd@tcp(127.0.0.1:3306)/record_db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func main() {
	dB()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/accounts", createAccHandler).Methods("POST")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}

func createAccHandler(w http.ResponseWriter, r *http.Request) {
	var newAcc Account
	err := json.NewDecoder(r.Body).Decode(&newAcc)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert the new account into the database
	stmt, err := db.Prepare("INSERT INTO Account (Username, Password, AccType, AccStatus) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newAcc.Username, newAcc.Password, newAcc.AccType, newAcc.AccStatus)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account created successfully")
}
