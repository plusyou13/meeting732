package main

import (
	"os"

	"github.com/plusyou13/comm-go/log"
	"github.com/plusyou13/meeting732/config"
	"github.com/plusyou13/meeting732/server"
	"github.com/plusyou13/meeting732/store"
)


func setupLog() {
	logpath := config.GetConfig().Log.LogPath
	if len(logpath) > 0 {
		os.MkdirAll(logpath, os.ModePerm)
	}

	log.NewLogger(logpath, config.GetConfig().Log.LogLevel)
}

func setupDB() {
	c := config.GetConfig().MongoDB
	store.SetupMgo(c.Host, c.UserName, c.Password)
}

func main() {
	setupLog()
	setupDB()
	server.RunService(config.GetConfig().Port)
}
