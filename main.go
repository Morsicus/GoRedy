package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"strconv"
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
	w.Write(myJson)
}

func CmdKeys(w http.ResponseWriter, r *http.Request) {
	data := &redisEntry{
		Key: "Yolo",
		Value: "PORRROOO",
	}
	myJson, _ := json.Marshal(data)
	w.Write(myJson)	
}

func CmdGetKey(w http.ResponseWriter, r *http.Request) {
	conf := RedisConfiguration{}
	conf.New()

	vars := mux.Vars(r)
	key := vars["key"]

	c, err := redis.DialTimeout("tcp", conf.IpAddr + ":" + strconv.Itoa(conf.Port), time.Duration(10)*time.Second)
	errHandler(err)
	defer c.Close()

	res, err := c.Cmd("get", key).Str()
	errHandler(err)

	data := &redisEntry{
		Key: key,
		Value: res,
	}
	res_json, _ := json.Marshal(data)
	w.Write(res_json)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/cmd/keys/{regEx}", CmdKeys)
	router.HandleFunc("/cmd/get/{key}", CmdGetKey)

	log.Fatal(http.ListenAndServe(":4242", router))
}
