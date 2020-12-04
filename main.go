package main

import (
	"log"

	"github.com/bluesky6529/golang_LinuxScript/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
