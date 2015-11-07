package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type redisEntry struct {
	Key    string
	Value  string
}

func Index(w http.ResponseWriter, r *http.Request) {
}

func CmdKeys(w http.ResponseWriter, r *http.Request) {
}

func CmdGetKey(w http.ResponseWriter, r *http.Request) {
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/cmd/keys/{regEx}", CmdKeys)
	router.HandleFunc("/cmd/get/{key}", CmdGetKey)

	log.Fatal(http.ListenAndServe(":4242", router))
}
