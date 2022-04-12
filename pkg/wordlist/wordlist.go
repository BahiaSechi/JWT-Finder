package wordlist

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"cmd/jwtfinder/pkg/domain"
)

func OpenFile(wordlist string) {
	file, err := os.Open(wordlist)

	if err != nil {
		log.Fatal("No wordlist provided")
	}

	readFile(file)
}

func readFile(file *os.File) {
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		urle := fileScanner.Text()
		fmt.Print(urle, "\n")
		domain.Navigate(urle)
	}
}
