package day_3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
)

func Part2() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(0)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	total := 0
	for scan.Scan() {
		line := scan.Text()
		s := toIntSlice(line)
		jtg := largestJoltageV2(s, 11, 0)
		t, _ := strconv.Atoi(jtg)
		total += t
	}
	fmt.Println(total)
}

func largestJoltageV2(bank []int, n, mx int) string {
	if n < 0 {
		return ""
	}

	for i := mx; i < len(bank)-n; i++ {
		if bank[i] > bank[mx] {
			mx = i
		}
	}

	return strconv.Itoa(bank[mx]) + largestJoltageV2(bank, n-1, mx+1)
}
