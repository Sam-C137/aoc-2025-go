package day_1

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func Part2() {
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
		ctx.moveV2(dir, step)
	}
	log.Println("Password: ", ctx.zeros)
}

func (ctx *Ctx) moveV2(dir string, steps int) {
	if dir == "L" {
		steps = -steps
		div, mod := divmod(steps, -100)
		ctx.zeros += div
		if ctx.start != 0 && ctx.start+mod <= 0 {
			ctx.zeros++
		}
	} else {
		div, mod := divmod(steps, 100)
		ctx.zeros += div
		if ctx.start != 100 && ctx.start+mod >= 100 {
			ctx.zeros++
		}
	}
	ctx.start = ((ctx.start+steps)%100 + 100) % 100
}

func divmod(num, denum int) (quot, rem int) {
	quot = num / denum
	rem = num % denum
	return
}
