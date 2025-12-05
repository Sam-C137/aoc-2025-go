package day_4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

type Grid [][]string

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	grid := makeGrid(file)

	total := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == "@" {
				if canAccess(grid, r, c) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func makeGrid(file *os.File) Grid {
	out := make(Grid, 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		out = append(out, strings.Split(scan.Text(), ""))
	}

	return out
}

var Dir = [][]int{
	[]int{-1, -1},
	[]int{-1, 0},
	[]int{-1, 1},
	[]int{0, -1},
	[]int{0, 1},
	[]int{1, -1},
	[]int{1, 0},
	[]int{1, 1},
}

func canAccess(grid Grid, r, c int) bool {
	count := 0

	for _, dir := range Dir {
		dr := dir[0]
		dc := dir[1]
		if dr+r < 0 || dr+r >= len(grid) || dc+c < 0 || dc+c >= len(grid[r]) {
			continue
		}
		if grid[dr+r][dc+c] == "@" {
			count++
		}
	}

	return count < 4
}
