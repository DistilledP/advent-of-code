package main

import (
	"log"
	"os"

	"github.com/DistilledP/advent-of-code/src/lib"
)

func main() {
	content, err := os.ReadFile("./input/input")
	if err != nil {
		log.Fatalf("failed to read input: %s", err)
	}

	result := lib.Part02_batch(string(content))

	log.Printf("the result: %d\n", result)
}
