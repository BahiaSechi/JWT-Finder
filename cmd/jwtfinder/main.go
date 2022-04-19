package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"cmd/jwtfinder/pkg/domain"
	"cmd/jwtfinder/pkg/wordlist"
)

// Configuration - Flags
var config struct {
	domain     string
	wordlist   string
	outputfile string
	// TODO add flag print mode or only output in file
}

const (
	usage = `usage: %s
Find JWT (JSON Web Token) in a web page

Options:
`
)

func main() {
	// Flags available
	flag.StringVar(&config.domain, "d", config.domain, "Specify the domain to scan")
	flag.StringVar(&config.wordlist, "w", config.wordlist, "Specify a path to the wordlist")
	flag.StringVar(&config.outputfile, "o", config.outputfile, "Output file to save results")
	// Display flags help
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Only one option is authorized
	if config.domain != "" {
		domain.Navigate(config.domain)
	} else if config.wordlist != "" {
		wordlist.OpenFile(config.wordlist)
	} else {
		log.Fatal("You must choose between the wordlist or the domain option")
	}
}
