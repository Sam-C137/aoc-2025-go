package day_6

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func Part2() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(1)
	}

	file, err := os.ReadFile(path.Join(path.Dir(filename), "input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(strings.Trim(string(file), "\n"), " ")
	for _, str := range s {
		fmt.Println(str)
		fmt.Println("================")
	}
}
