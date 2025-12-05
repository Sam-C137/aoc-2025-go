package day_5

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.ReadFile(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`\n\s+`)
	s := re.Split(string(file), -1)
	ranges := getRanges(strings.Split(s[0], "\n"))

	fresh := 0
	for _, idStr := range strings.Split(s[1], "\n") {
		idInt, _ := strconv.Atoi(idStr)
		for _, pair := range ranges {
			if idInt >= pair[0] && idInt <= pair[1] {
				fresh++
				break
			}
		}
	}
	fmt.Println(fresh)
}

func getRanges(raw []string) [][]int {
	out := make([][]int, 0)

	for _, pair := range raw {
		lr := strings.Split(pair, "-")
		l, _ := strconv.Atoi(lr[0])
		r, _ := strconv.Atoi(lr[1])
		out = append(out, []int{l, r})
	}
	return out
}
