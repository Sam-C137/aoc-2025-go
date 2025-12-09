package day_8

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"runtime"
	"slices"
	"sort"
	"strconv"
	"strings"

	day7 "github.com/Sam-C137/aoc-2025-go/cmd/day-7"
	"github.com/idsulik/go-collections/v3/disjointset"
)

const (
	NumPairs = 1000
	TOP      = 3
)

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	grid := parseInput(file)
	e := edges(grid)
	sort.Slice(e, func(i, j int) bool {
		return dist3D(grid, e[i]) < dist3D(grid, e[j])
	})
	ds := disjointset.New[int]()
	for i := range grid {
		ds.MakeSet(i)
	}

	for _, edge := range e[:NumPairs] {
		a, b := edge[0], edge[1]
		ds.Union(a, b)
	}

	circuits := make([]int, 0)
	for _, s := range ds.GetSets() {
		circuits = append(circuits, len(s))
	}

	slices.Sort(circuits)
	var total = 1
	for _, n := range circuits[len(circuits)-TOP:] {
		total *= n
	}

	fmt.Println(total)
}

func dist3D(grid day7.Grid[int], edge []int) float64 {
	p1, p2 := grid[edge[0]], grid[edge[1]]
	dx := float64(p1[0] - p2[0])
	dy := float64(p1[1] - p2[1])
	dz := float64(p1[2] - p2[2])
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func parseInput(file *os.File) day7.Grid[int] {
	grid := make(day7.Grid[int], 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		rs := strings.Split(scan.Text(), ",")
		ri := make([]int, len(rs))

		for i, r := range rs {
			n, _ := strconv.Atoi(r)
			ri[i] = n
		}
		grid = append(grid, ri)
	}

	return grid
}

func edges(in day7.Grid[int]) day7.Grid[int] {
	grid := make(day7.Grid[int], 0)

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			grid = append(grid, []int{i, j})
		}
	}

	return grid
}
