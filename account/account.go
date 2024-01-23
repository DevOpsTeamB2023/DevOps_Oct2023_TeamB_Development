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
	router.Use(corsMiddleware)
	router.HandleFunc("/api/v1/accounts", createAccHandler).Methods("POST")
	router.HandleFunc("/api/v1/accounts", getAccHandler).Methods("GET")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Api-Key, X-Requested-With, Content-Type, Accept, Authorization")
		next.ServeHTTP(w, r)
	})
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

func getAccHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username == "" || password == "" {
		http.Error(w, "Username and Password parameters are required", http.StatusBadRequest)
		return
	}

	var acc Account
	err := db.QueryRow("SELECT AccID, Username, Password, AccType, AccStatus FROM Account WHERE Username = ? AND Password = ?", username, password).Scan(&acc.AccID, &acc.Username, &acc.Password, &acc.AccType, &acc.AccStatus)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Username or Password", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with user information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}
