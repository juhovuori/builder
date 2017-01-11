package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

const fn = "version.json"

type versionInfo struct {
	Commit    string `json:"commit"`
	BuildTime string `json:"build-time"`
	BootTime  string `json:"boot-time"`
	Uptime    int    `json:"uptime"`
}

func version() versionInfo {
	once.Do(loadVersion)
	v := globalVersionInfo
	return v
}

var once sync.Once
var globalVersionInfo versionInfo

func loadVersion() {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Printf("Error reading %s: %s\n", fn, err.Error())
		return
	}
	err = json.Unmarshal(data, &globalVersionInfo)
	if err != nil {
		log.Printf("Error decoding version JSON in %s: %s\n", fn, err.Error())
		return
	}
}
