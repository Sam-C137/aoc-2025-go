package day_6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type Grid[T any] [][]T

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	nums, ops := processFile(file)

	var total int64
	for row := 0; row < len(nums[0]); row++ {
		var acc int64
		if ops[row] == "*" {
			acc = 1
		}
		for col := 0; col < len(nums); col++ {
			op := ops[row]
			num := nums[col][row]
			if op == "*" {
				acc *= num
			} else {
				acc += num
			}
		}
		total += acc
	}
	fmt.Println(total)
}

func processFile(file *os.File) (Grid[int64], []string) {
	nums := make(Grid[int64], 0)
	var ops []string

	scan := bufio.NewScanner(file)
	wre := regexp.MustCompile(`\s+`)
	nre := regexp.MustCompile(`\d+`)

	for scan.Scan() {
		t := scan.Text()
		row := nre.FindAll([]byte(t), -1)
		nr := make([]int64, 0)
		for _, bytes := range row {
			n, _ := strconv.ParseInt(string(bytes), 10, 64)
			nr = append(nr, n)
		}
		if strings.HasPrefix(t, "*") || strings.HasPrefix(t, "+") {
			opr := wre.Split(t, -1)
			ops = opr
		}
		if len(nr) > 0 {
			nums = append(nums, nr)
		}
	}

	return nums, ops
}
