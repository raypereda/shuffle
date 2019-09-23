// This command reads an input text file, and writes shuffled lines.
// Output is set to standard output.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

var seed int64

func init() {
	flag.Int64Var(&seed, "seed", 1, "Optional random number seed")
}

func main() {
	usage := func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "    shuffle [-seed N] <file>\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    shuffle [-h] [-help]\n")

		fmt.Fprintf(flag.CommandLine.Output(), "<file> is a required input file.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "This command reads an input text file, and writes shuffled lines.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Output is set to standard output.\n")
		flag.PrintDefaults()
	}
	flag.Usage = usage
	flag.Parse()
	rand.Seed(seed)

	if len(os.Args) < 2 {
		fmt.Fprintf(flag.CommandLine.Output(), "Input file is missing from command-line.\n")
		usage()
		os.Exit(1)
	}

	inputfile := os.Args[len(os.Args)-1]
	lines, err := readLines(inputfile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })

	for _, line := range lines {
		fmt.Println(line)
	}
}
