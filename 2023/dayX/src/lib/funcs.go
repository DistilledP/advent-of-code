package lib

import (
	"log"
	"os"
)

func Part01() {
	log.Println("TODO")
}

func Part02() {
	log.Println("TODO")
}

func debugLog(items ...any) {
	if os.Getenv("DEBUG") == "1" {
		log.Println(items...)
	}
}
