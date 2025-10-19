package main

import (
	internal "go-template/internal"
	log "log"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Fatal(err)
	}
}
