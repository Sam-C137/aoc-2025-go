package day_1

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type Ctx struct {
	start int
	zeros int
}

func Part1() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.Open(filepath.Join(filepath.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(file)
	ctx := &Ctx{
		50, 0,
	}
	for scan.Scan() {
		line := scan.Text()
		step, _ := strconv.Atoi(line[1:])
		dir := line[:1]
		ctx.move(dir, step)
	}
	log.Println("Password: ", ctx.zeros)
}

func (ctx *Ctx) move(dir string, steps int) {
	if dir == "L" {
		steps = -steps
	}
	ctx.start = ((ctx.start+steps)%100 + 100) % 100
	ctx.track()
}

func (ctx *Ctx) track() {
	if ctx.start == 0 {
		ctx.zeros++
	}
}
