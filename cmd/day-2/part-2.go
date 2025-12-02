package day_2

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Part2() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(filepath.Join(filepath.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(ScanRanges)

	ctx := &Ctx{0}
	for scan.Scan() {
		text := scan.Text()
		s := strings.Split(text, "-")
		start, _ := strconv.Atoi(s[0])
		end, _ := strconv.Atoi(s[1])
		ctx.processRangeV2(start, end)
	}
	log.Println("Invalid id sum: ", ctx.total)
}

func (ctx *Ctx) processRangeV2(start, end int) {
	for i := start; i <= end; i++ {
		if hasRepeatedSubstr(strconv.Itoa(i)) {
			ctx.total += i
		}
	}
}

func hasRepeatedSubstr(s string) bool {
	str := s + s
	return strings.Contains(str[1:len(str)-1], s)
}
