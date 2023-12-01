package main

import (
	"io/ioutil"
	"log"

	"github.com/DistilledP/advent-of-code/src/lib"
)

func main() {
	content, err := ioutil.ReadFile("./input/input1")

	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}

	combinedValue := lib.Part02(string(content))

	log.Printf("Combined value is: %d", combinedValue)
}
