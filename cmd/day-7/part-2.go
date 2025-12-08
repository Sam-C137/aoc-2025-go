package day_7

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

func Part2() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	start, grid := parseInput(file)
	fmt.Println(goDownV2(start, grid, make(map[string]int)))
}

func goDownV2(point Point, grid Grid[string], seen map[string]int) int {
	h := hash(point)
	val, ok := seen[h]
	if ok {
		return val
	}
	if point.r >= len(grid) || point.r < 0 || point.c >= len(grid[0]) || point.c < 0 {
		seen[hash(point)] = 1
		return 1
	}

	if grid[point.r][point.c] != "^" {
		p := Point{point.r + 1, point.c}
		result := goDownV2(p, grid, seen)
		seen[hash(p)] = result
		return result
	}

	l := Point{point.r + 1, point.c - 1}
	r := Point{point.r + 1, point.c + 1}
	rl := goDownV2(l, grid, seen)
	seen[hash(l)] = rl
	rr := goDownV2(r, grid, seen)
	seen[hash(r)] = rr
	return rl + rr
}
