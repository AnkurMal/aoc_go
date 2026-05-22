package day11

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

func Part1() {
	fmt.Println("Part 1:", getAns(false))
}

func Part2() {
	fmt.Println("Part 2:", getAns(true))
}

func getAns(part2 bool) int {
	items, ops, direct, test := parse()
	inspect := make([]int, len(items))
	prod, rng := 1, 20

	if part2 {
		for _, num := range test {
			prod *= num
		}
		rng = 10000
	}

	for range rng {
		for i := range len(items) {
			for len(items[i]) > 0 {
				level := items[i][0]
				items[i] = items[i][1:]

				level = int(math.Pow(float64(level*ops[i][0]+ops[i][1]), float64(ops[i][2])))
				if part2 {
					level %= prod
				} else {
					level /= 3
				}

				if level%test[i] != 0 {
					items[direct[i][1]] = append(items[direct[i][1]], level)
				} else {
					items[direct[i][0]] = append(items[direct[i][0]], level)
				}

				inspect[i]++
			}
		}
	}

	slices.Sort(inspect)
	slices.Reverse(inspect)
	return inspect[0] * inspect[1]
}

func parse() ([][]int, [][]int, [][]int, []int) {
	var items, ops, direct [][]int
	var test []int

	for line := range strings.SplitSeq(data, "\r\n\r\n") {
		grp := strings.Split(line, "\r\n")

		var item []int
		for num := range strings.SplitSeq(grp[1][18:], ", ") {
			n, _ := strconv.Atoi(num)
			item = append(item, n)
		}
		items = append(items, item)

		syms := strings.Split(grp[2][23:], " ")
		if num, err := strconv.Atoi(syms[1]); err != nil {
			ops = append(ops, []int{1, 0, 2})
		} else {
			if syms[0] == "*" {
				ops = append(ops, []int{num, 0, 1})
			} else {
				ops = append(ops, []int{1, num, 1})
			}
		}

		spl := strings.Split(grp[3], " ")
		num, _ := strconv.Atoi(spl[len(spl)-1])
		test = append(test, num)

		direct = append(direct, []int{int(grp[4][len(grp[4])-1] - '0'), int(grp[5][len(grp[5])-1] - '0')})
	}

	return items, ops, direct, test
}
