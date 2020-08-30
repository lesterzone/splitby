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
		// only append to the storage if we have our first line.
		// which represents the first time we find the spliter
		if !strings.Contains(text, *spliter) && len(storage) >= 1 {
			// append text to storage so we can print it later when we have
			// the next occurence
			storage = append(storage, text)
		}

		if strings.Contains(text, *spliter) {
			// let's print line, we have found the next occurence
			logger.Println(strings.Join(storage, *joiner) + *separator)
			// let's reset storage
			storage = []string{text}
		}
	}
}
