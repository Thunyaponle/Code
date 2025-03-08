package main

import (
	"log"
)

func main() {
	router := SetupRouter()

	log.Println("Starting Coffee Shop API server on :8080")
	router.Run(":8080")
}
