package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/nxshock/promodj/api"
)

func init() {
	log.SetFlags(0)

	rand.Seed(time.Now().Unix())

	if len(os.Args) == 2 {
		if err := initConfig(os.Args[1]); err != nil {
			log.Fatalln("config error:", err)
		}
	} else {
		if err := initConfig(configFilePath); err != nil {
			log.Fatalln("config error:", err)
		}
	}

	Log(LogLevelDebug, "Updating genre list...")
	err := api.UpdateGenres()
	if err != nil {
		log.Fatalln(err)
	}

	Log(LogLevelDebug, "Initialization completed.")
}

func main() {
	err := http.ListenAndServe(config.ListenAddr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
