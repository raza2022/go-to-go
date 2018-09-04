package main

import (
	"dao"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//import (
//	"encoding/json"
//	"log"
//	"net/http"
//
//	"gopkg.in/mgo.v2/bson"
//
//)

type Config struct {
	Server   string
	Database string
}

var configg = Config{"localhost", "movies_db"}
var daoo = dao.MoviesDAO{"localhost", "movies_db"}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := daoo.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	//config.Read()
	daoo.Server = configg.Server
	daoo.Database = configg.Database
	daoo.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
