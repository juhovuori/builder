package version

//go:generate sh -c "./version.sh >version.json"
//go:generate go-bindata -o version-data.go -pkg version version.json

import (
	"encoding/json"
	"log"
	"time"
)

const fn = "version.json"

// Info describes builder version
type Info struct {
	Commit      string `json:"commit"`
	BuildTime   string `json:"build-time"`
	StartupTime string `json:"startup-time"`
	Time        string `json:"time"`
	Uptime      int64  `json:"uptime"`
}

func init() {
	startup = time.Now()

	data, err := versionJsonBytes()
	if err != nil {
		log.Printf("Error reading bundled version data %v\n", err)
		return
	}
	err = json.Unmarshal(data, &globalInfo)
	if err != nil {
		log.Printf("Error decoding version JSON in %s: %s\n", fn, err.Error())
		return
	}
	globalInfo.StartupTime = startup.Format(time.RFC3339)
}

// Version returns current version info
func Version() Info {
	v := globalInfo
	now := time.Now()
	v.Time = now.Format(time.RFC3339)
	v.Uptime = now.Unix() - startup.Unix()
	return v
}

var globalInfo Info
var startup time.Time
