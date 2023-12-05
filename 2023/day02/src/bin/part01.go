package main

import (
	"log"
	"os"

	"github.com/DistilledP/advent-of-code/src/lib"
)

func main() {
	content, err := os.ReadFile("./input/input")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	result := lib.Part01(string(content))

	log.Printf("the result: %d\n", result)
}
