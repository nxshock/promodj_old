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

	if len(os.Args) > 1 {
		err := initConfig(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		initConfig(configFilePath)
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
