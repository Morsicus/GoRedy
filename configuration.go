package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type RedisConfiguration struct {
	IpAddr string
	Port   int
}

func (conf *RedisConfiguration) New() *RedisConfiguration {
	file, err := os.Open("redis_conf.json")
	if err != nil {
		fmt.Println("ERROR opening redis configuration : ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("ERROR decoding JSON conf file : ", err)
	}

	return conf
}
