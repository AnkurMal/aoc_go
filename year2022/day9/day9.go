package day9

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed day9.txt
var data string

type Coord struct {
	x, y int
}

func Part1() {
	var h, t Coord
	dirs := map[string]Coord{
		"R": {x: 1},
		"U": {y: 1},
		"L": {x: -1},
		"D": {y: -1},
	}

	set := make(map[Coord]struct{})

	for line := range strings.SplitSeq(data, "\r\n") {
		spl := strings.Split(line, " ")
		dir, _ := strconv.Atoi(spl[1])
		move := dirs[spl[0]]

		for range dir {
			h.x += move.x
			h.y += move.y

			if math.Abs(float64(h.x)-float64(t.x)) > 1 || math.Abs(float64(h.y)-float64(t.y)) > 1 {
				t.x = h.x - move.x
				t.y = h.y - move.y

				set[t] = struct{}{}
			}
		}
	}

	fmt.Println("Part 1:", len(set)+1)
}

func Part2() {
	var t [10]Coord
	dirs := map[string]Coord{
		"R": {x: 1},
		"U": {y: 1},
		"L": {x: -1},
		"D": {y: -1},
	}

	set := make(map[Coord]struct{})

	for line := range strings.SplitSeq(data, "\r\n") {
		spl := strings.Split(line, " ")
		dir, _ := strconv.Atoi(spl[1])
		move := dirs[spl[0]]

		for range dir {
			t[0].x += move.x
			t[0].y += move.y

			for i := 1; i < 10; i++ {
				dx := t[i-1].x - t[i].x
				dy := t[i-1].y - t[i].y

				if math.Abs(float64(dx)) > 1 || math.Abs(float64(dy)) > 1 {
					t[i].x += sign(dx)
					t[i].y += sign(dy)
				}
			}

			set[t[9]] = struct{}{}
		}
	}

	fmt.Println("Part 1:", len(set))
}

func sign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}
