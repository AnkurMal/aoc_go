package day10

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string
var cycle, next int

func Part1() {
	x, res := 1, 0

	for line := range strings.SplitSeq(data, "\r\n") {
		spl := strings.Split(line, " ")

		if len(spl) == 1 {
			res += checkCycle(x)
		} else {
			val, _ := strconv.Atoi(spl[1])

			res += checkCycle(x)
			res += checkCycle(x)
			x += val
		}
	}

	fmt.Println("Part 1:", res)
}

func Part2() {
	fmt.Printf("Part 2:\n\n")
	x := 1

	for line := range strings.SplitSeq(data, "\r\n") {
		spl := strings.Split(line, " ")

		if len(spl) == 1 {
			draw(x)
		} else {
			val, _ := strconv.Atoi(spl[1])

			draw(x)
			draw(x)
			x += val
		}
	}
}

func checkCycle(x int) int {
	res := 0
	cycle++
	for i := 20; i < 221; i += 40 {
		if cycle == i {
			res += cycle * x
			break
		}
	}

	return res
}

func draw(x int) {
	dec := cycle % 40
	if x-1 <= dec && x+1 >= dec {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}

	cycle++
	if cycle/40 > next {
		next = cycle / 40
		fmt.Println()
	}
}
