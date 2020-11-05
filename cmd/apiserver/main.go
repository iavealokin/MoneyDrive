package main

import (
	"log"

	"../../internal/app/apiserver/apiserver.go"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
