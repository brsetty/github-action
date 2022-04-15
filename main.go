// Launch microservice server- main.go
package main

import (
	"https://github.com/brsetty/github-action/tree/main/microservice"
	"log"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
