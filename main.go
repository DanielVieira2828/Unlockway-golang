package main

import (
	"log"

	"github.com/DanielVieirass/Unlockway_golang/server"
)

func main() {

	if err := server.Run(); err != nil {
		log.Println(err.Error())
		return
	}
}
