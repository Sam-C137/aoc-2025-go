package day_2

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Ctx struct {
	total int
}

func Part1() {
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
		ctx.processRange(start, end)
	}
	log.Println("Invalid id sum: ", ctx.total)
}

func (ctx *Ctx) processRange(start, end int) {
	for i := start; i <= end; i++ {
		if isInvalidId(strconv.Itoa(i)) {
			ctx.total += i
		}
	}
}

func isInvalidId(s string) bool {
	return s[0:len(s)/2] == s[len(s)/2:]
}

func ScanRanges(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0

	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if string(r) == "," {
			return i + width, data[start:i], nil
		}
	}

	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	return start, nil, nil
}
