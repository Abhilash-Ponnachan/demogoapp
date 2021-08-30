package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

const configFile = "./config.json"

type configData struct {
	Port         string
	DefaultPage  string
	AssetsDir    string
	DbConnection string
}

var once sync.Once
var cf *configData

func config() *configData {
	if cf == nil {
		once.Do(
			func() {
				cf = &configData{}
				cf.load()
			})
	}
	return cf
}

func (cf *configData) load() {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(bytes, cf)
	// <TO DO> chang unmarshall to map[string]string
	// iterate and check each key is loaded to not empty
	// assign to 'cf' fields
	log.Printf("cf = %v\n", cf)
	if err != nil {
		panic(err.Error())
	}
	// for 'Port' override with os.env value!!
}
