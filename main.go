package main

import (
	"fmt"
	"log"

	"github.com/juliotorresmoreno/server-api/app"
)

func main() {
	server := app.NewServer()
	fmt.Println("Serving API on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
