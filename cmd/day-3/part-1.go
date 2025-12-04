package day_3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func Part1() {
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
		jtg, _ := goodLargestJoltage(toIntSlice(line))
		total += jtg
	}
	fmt.Println(total)
}

func toIntSlice(str string) []int {
	vs := strings.Split(str, "")
	vi := make([]int, len(vs))

	for i, s := range vs {
		vi[i], _ = strconv.Atoi(s)
	}
	return vi
}

func badLargestJoltage(bank []int) int {
	jtg := 0

	for i, num := range bank {
		for j := i + 1; j < len(bank); j++ {
			sum, _ := strconv.Atoi(strconv.Itoa(num) + strconv.Itoa(bank[j]))
			jtg = max(jtg, sum)
		}
	}

	return jtg
}

func goodLargestJoltage(bank []int) (int, error) {
	maxL := 0

	for i := 0; i < len(bank)-1; i++ {
		if bank[maxL] < bank[i] {
			maxL = i
		}
	}

	maxR := maxL + 1
	for i := maxL + 1; i < len(bank); i++ {
		if bank[maxR] < bank[i] {
			maxR = i
		}
	}

	return strconv.Atoi(strconv.Itoa(bank[maxL]) + strconv.Itoa(bank[maxR]))
}
