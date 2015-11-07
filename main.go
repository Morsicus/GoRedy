package main

import (
	"fmt"
	//"os"
	"log"
	"net/http"
	"encoding/json"
	//"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"

)

type redisEntry struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := &redisEntry{
		Key: "Yolo",
		Value: "PORRROOO",
	}
	myJson, _ := json.Marshal(data)
	fmt.Println(string(myJson))
}

func CmdKeys(w http.ResponseWriter, r *http.Request) {
	data := &redisEntry{
		Key: "Yolo",
		Value: "PORRROOO",
	}
	myJson, _ := json.Marshal(data)
	fmt.Println(string(myJson))
}

func CmdGetKey(w http.ResponseWriter, r *http.Request) {
	data := &redisEntry{
		Key: "Yolo",
		Value: "PORRROOO",
	}
	myJson, _ := json.Marshal(data)
	fmt.Println(string(myJson))

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/cmd/keys/{regEx}", CmdKeys)
	router.HandleFunc("/cmd/get/{key}", CmdGetKey)

	log.Fatal(http.ListenAndServe(":4242", router))
}
