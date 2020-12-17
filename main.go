package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
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
		initConfig("")
	}
}

func main() {
	err := http.ListenAndServe(config.ListenAddr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
