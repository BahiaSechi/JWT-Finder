package wordlist

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"cmd/jwtfinder/pkg/domain"
)

// Open the provided wordlist
func OpenFile(wordlist string) {
	file, err := os.Open(wordlist)

	if err != nil {
		log.Fatal("No wordlist provided")
	}

	readFile(file)
}

// Read every line in the wordlist
func readFile(file *os.File) {
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		urle := fileScanner.Text()
		fmt.Print(urle, "\n")
		//Verify cookies for every domain in wordlist
		domain.Navigate(urle)
	}
}
