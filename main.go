package main

import (
	"log"

	"github.com/shamane/backend/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8080"))
}
