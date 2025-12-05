package day_4

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

type Point struct {
	r, c int
}

func Part2() {
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

	for {
		curr, points := didAccess(grid)
		if curr == 0 {
			break
		}

		for _, point := range points {
			grid[point.r][point.c] = "."
		}
		total += curr
	}

	fmt.Println(total)
}

func didAccess(grid Grid) (int, []Point) {
	total := 0
	points := make([]Point, 0)
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == "@" {
				if canAccess(grid, r, c) {
					total++
					points = append(points, Point{r, c})
				}
			}
		}
	}

	return total, points
}
