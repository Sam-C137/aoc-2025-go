package day_9

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type Point struct {
	r, c int
}

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}

	points := parseInput(file)
	var largest float64

	for _, point := range points {
		for _, opposite := range points {
			if point.r == opposite.r && point.c == opposite.c {
				continue
			}
			if point.r == opposite.r || point.c == opposite.c {
				continue
			}
			w := math.Abs(float64(point.c-opposite.c)) + 1
			l := math.Abs(float64(point.r-opposite.r)) + 1
			largest = max(largest, w*l)
		}
	}

	fmt.Println(int(largest))
}

func parseInput(file *os.File) []Point {
	points := make([]Point, 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		p := strings.Split(scan.Text(), ",")
		c, _ := strconv.Atoi(p[0])
		r, _ := strconv.Atoi(p[1])
		points = append(points, Point{r, c})
	}

	return points
}
