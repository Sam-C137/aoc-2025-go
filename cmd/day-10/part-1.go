package day_10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
)

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		re := regexp.MustCompile(`\[([.#]+)\] ([()\d, ]+) \{([\d,]+)\}`)
		matches := re.FindAllStringSubmatch(scan.Text(), -1)

		for i, match := range matches {
			fmt.Printf("Match %d: %v\n", i, match)
			fmt.Printf("  Pattern: %s\n", match[1])
			fmt.Printf("  Groups: %s\n", match[2])
			fmt.Printf("  Numbers: %s\n", match[3])
		}
		fmt.Println("Number of matches:", len(matches))
	}
}
