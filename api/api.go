package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql" // import MySQL database
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type errorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type statusResponse struct {
	Status string `json:"status"`
}

var password string
var message string

var db *sql.DB

//Initalize initalizes the API
func Initalize(r *httprouter.Router) {
	var err error

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")

	messageRaw, err := ioutil.ReadFile("./message.txt")
	if err != nil {
		panic(err)
	}

	message = string(messageRaw)

	db, err = sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbDatabase+"?charset=utf8mb4&collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	password = os.Getenv("PASSWORD")

	fmt.Printf("PASSWORD: %s\n", password)

	r.POST("/login", loginRoute)
	r.GET("/contestants", contestantsRoute)
	r.POST("/vote", voteRoute)
	r.GET("/photos/:name", photosRoute)
}
func writeJSON(w http.ResponseWriter, status int, thing interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(thing)
}
