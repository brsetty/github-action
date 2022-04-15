// Launch microservice server- main.go
package main

import (
	"github.com/brsetty/github-action/microservice"
	"log"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
