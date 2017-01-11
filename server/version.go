package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

const fn = "version.json"

type versionInfo struct {
	Commit      string `json:"commit"`
	BuildTime   string `json:"build-time"`
	StartupTime string `json:"startup-time"`
	Time        string `json:"time"`
	Uptime      int64  `json:"uptime"`
}

func version() versionInfo {
	v := globalVersionInfo
	now := time.Now()
	v.Time = now.Format(time.RFC3339)
	v.Uptime = now.Unix() - startup.Unix()
	return v
}

var globalVersionInfo versionInfo
var startup time.Time

func setupVersion() {
	startup = time.Now()
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
	globalVersionInfo.StartupTime = startup.Format(time.RFC3339)
}
