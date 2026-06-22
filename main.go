package main

import (
	"go-starter-kit/app"
	"log"
)

func main() {
	a := app.New()
	if err := a.Start(":9000"); err != nil {
		log.Fatal(err)
	}
	log.Println("Server started on :9000")
}
