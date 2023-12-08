package lib

import (
	"log"
	"os"
)

func Part01(content string) int {
	log.Println("TODO")

	return 0
}

func Part02(content string) int {
	log.Println("TODO")

	return 0
}

func debugLog(items ...any) {
	if os.Getenv("DEBUG") == "1" {
		log.Println(items...)
	}
}
