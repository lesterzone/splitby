package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

var logger = log.New(os.Stdout, "", 0)

func main() {
	spliter := flag.String("spliter", "\n", "which expression we are looking for?")
	separator := flag.String("separator", "\n", "separate lines by?")
	joiner := flag.String("joiner", "\n", "join lines until next occurrence by?")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	storage := []string{}

	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, *spliter) {
			// append text to storage so we can print it later when we have
			// the next occurence
			storage = append(storage, text)
		}

		if strings.Contains(text, *spliter) {
			// let's print line only if we have content in the storage
			if len(storage) > 0 {
				logger.Println(strings.Join(storage, *joiner) + *separator)
				// reset storage with found text
				storage = []string{text}
			}
		}
	}
}
