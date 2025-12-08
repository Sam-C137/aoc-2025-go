package day_7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/idsulik/go-collections/v3/queue"
	"github.com/idsulik/go-collections/v3/set"
)

type Grid[T any] [][]T
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
	start, grid := parseInput(file)
	fmt.Println(goDown(start, grid))
}

func parseInput(file *os.File) (Point, Grid[string]) {
	grid := make(Grid[string], 0)
	var point Point

	scan := bufio.NewScanner(file)
	var r int
	for scan.Scan() {
		t := scan.Text()
		grid = append(grid, strings.Split(t, ""))
		for c, char := range t {
			if string(char) == "S" {
				point = Point{
					r, c,
				}
			}
		}
		r++
	}

	return point, grid
}

func goDown(start Point, grid Grid[string]) int {
	q := queue.New[Point](10)
	q.Enqueue(start)
	s := set.New[string]()
	var total int

	for q.Len() > 0 {
		node, ok := q.Dequeue()
		if !ok {
			break
		}
		if node.c >= len(grid[0]) || node.c < 0 || node.r >= len(grid) || node.r < 0 {
			continue
		}
		if s.Has(hash(node)) {
			continue
		}
		s.Add(hash(node))
		if grid[node.r][node.c] == "^" {
			q.Enqueue(Point{node.r + 1, node.c - 1})
			q.Enqueue(Point{node.r + 1, node.c + 1})
			total++
		} else {
			q.Enqueue(Point{node.r + 1, node.c})
		}
	}

	return total
}

func hash(point Point) string {
	return strconv.Itoa(point.r) + ":" + strconv.Itoa(point.c)
}
