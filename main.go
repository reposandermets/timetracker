package main

import (
	"log"

	sw "github.com/reposandermets/timetracker/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}
