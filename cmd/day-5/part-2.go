package day_5

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func Part2() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.ReadFile(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`\n\s+`)
	str := re.Split(string(file), -1)

	ranges := getRanges(strings.Split(str[0], "\n"))
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var last []int
	var total int
	for _, pair := range ranges {
		if last == nil {
			last = pair
		} else if last[1] < pair[0] {
			total += last[1] - last[0] + 1
			last = pair
		} else {
			last = []int{last[0], max(last[1], pair[1])}
		}
	}
	total += last[1] - last[0] + 1

	fmt.Println(total)
}
