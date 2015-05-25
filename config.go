package main

import (
	"flag"
	"fmt"
)

type redisConfig struct {
	Host string
}

var REDIS_CONFIG_COLLECTION = map[string]redisConfig{
	"development": redisConfig{Host: "127.0.0.1:6379"},
	"test":        redisConfig{Host: "127.0.0.1:6379"},
	"integration": redisConfig{Host: "127.0.0.1:6379"},
	"production":  redisConfig{Host: "redis_producion.com:6379"},
}

var RESOURCES_COLLECTION = map[string]string{
	"image_process_channel": "image_process_message#*",
}

var REDIS_CONFIG redisConfig
var RESOURCES map[string]string
var MAXWORKER = 50

func config() {
	fmt.Println("in function config")
	envPtr := flag.String("env", "integration", "env")
	flag.Parse()

	env := *envPtr

	println("env:", env)

	REDIS_CONFIG = REDIS_CONFIG_COLLECTION[env]
	println("redisconfig:", REDIS_CONFIG.Host)
	RESOURCES = RESOURCES_COLLECTION
}
func init(){
	config()
}