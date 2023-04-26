package main

import (
	"log"

	"github.com/Brigant/PetPorject/app/transport/rest"
)

func main() {
	if err := rest.SetupAndRun(); err != nil {
		log.Fatalf("error while SetupAndRun server: %s", err.Error())
	}
}
