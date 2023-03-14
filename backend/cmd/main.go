package main

import (
	"log"

	"github.com/Brigant/GoPetPorject/backend/app/transport/rest"
)

func main() {
	if err := rest.SetupAndRun(); err != nil {
		log.Fatalf("error while SetupAndRun server: %s", err.Error())
	}
}
