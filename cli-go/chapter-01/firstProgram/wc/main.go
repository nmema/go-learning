package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wc := 0

	// For every words scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Defining a boolean flag -b to count bytes
	bytes := flag.Bool("b", false, "Count bytes")

	// Defining a boolean flag -file to read data from file
	file := flag.String("file", "", "Read data from file")

	// Parsing the flags provided by the user
	flag.Parse()

	var input *os.File
	var err error
	if *file != "" {
		input, err = os.Open(*file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		input = os.Stdin
	}

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(count(input, *lines, *bytes))
}
