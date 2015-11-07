package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"github.com/fzzy/radix/redis"
	"github.com/gorilla/mux"
)

type redisEntry struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func errHandler(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
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
	vars := mux.Vars(r)
	key := vars["key"]

	c, err := redis.DialTimeout("tcp", "172.17.0.2:6379", time.Duration(10)*time.Second)
	errHandler(err)
	defer c.Close()

	res, err := c.Cmd("get", key).Str()
	errHandler(err)

	fmt.Println(res)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/cmd/keys/{regEx}", CmdKeys)
	router.HandleFunc("/cmd/get/{key}", CmdGetKey)

	log.Fatal(http.ListenAndServe(":4242", router))
}
