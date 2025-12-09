package day_8

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sort"

	"github.com/idsulik/go-collections/v3/disjointset"
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
	grid := parseInput(file)
	e := edges(grid)
	sort.Slice(e, func(i, j int) bool {
		return dist3D(grid, e[i]) < dist3D(grid, e[j])
	})
	ds := disjointset.New[int]()
	for i := range grid {
		ds.MakeSet(i)
	}

	circuits := len(grid)
	for _, edge := range e {
		a, b := edge[0], edge[1]
		if ds.Find(a) == ds.Find(b) {
			continue
		}
		ds.Union(a, b)
		circuits--
		if circuits == 1 {
			fmt.Println(grid[a][0] * grid[b][0])
			break
		}
	}
}
