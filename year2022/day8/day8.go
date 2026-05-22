package day8

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed data.txt
var data string

func Part1() {
	var vec [][]int

	for val := range strings.SplitSeq(data, "\r\n") {
		var internal []int
		for _, ch := range val {
			internal = append(internal, int(ch-'0'))
		}
		vec = append(vec, internal)
	}

	n := len(vec)
	count := 4*n - 4

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			if checkTopDown(vec, i, j) || checkLeftRight(vec, i, j) {
				count += 1
			}
		}
	}

	fmt.Println("Part 1:", count)
}

func Part2() {
	var vec [][]int

	for val := range strings.SplitSeq(data, "\r\n") {
		var internal []int
		for _, ch := range val {
			internal = append(internal, int(ch-'0'))
		}
		vec = append(vec, internal)
	}

	n := len(vec)
	var max int

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			target := vec[i][j]
			var a, b, c, d int

			for _, val := range slices.Backward(vec[i][:j]) {
				a++
				if val >= target {
					break
				}
			}

			for _, val := range vec[i][j+1:] {
				b++
				if val >= target {
					break
				}
			}

			for x := i - 1; x >= 0; x-- {
				c++
				if vec[x][j] >= target {
					break
				}
			}

			for x := i + 1; x < n; x++ {
				d++
				if vec[x][j] >= target {
					break
				}
			}

			res := a * b * c * d
			if max < res {
				max = res
			}
		}
	}

	fmt.Println("Part 2:", max)
}

func checkLeftRight(vec [][]int, i, j int) bool {
	target := vec[i][j]

	visibleLeft := true
	for _, val := range vec[i][:j] {
		if val >= target {
			visibleLeft = false
			break
		}
	}

	visibleRight := true
	for _, val := range vec[i][j+1:] {
		if val >= target {
			visibleRight = false
			break
		}
	}

	return visibleLeft || visibleRight
}

func checkTopDown(vec [][]int, i, j int) bool {
	target := vec[i][j]
	n := len(vec)

	visibleTop := true
	for x := range i {
		if vec[x][j] >= target {
			visibleTop = false
			break
		}
	}

	visibleBottom := true
	for x := i + 1; x < n; x++ {
		if vec[x][j] >= target {
			visibleBottom = false
			break
		}
	}

	return visibleTop || visibleBottom
}
