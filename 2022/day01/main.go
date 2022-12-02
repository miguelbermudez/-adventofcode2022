package main

import (
	"advent2022/cast"
	"advent2022/util"
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	elves := parseInput(input)
	elfCalorieSums := util.Map(elves, func(elfCalories []int) int {
		return util.Reduce(elfCalories, func(acc int, current int) int {
			return acc + current
		}, 0)
	})
	sort.Ints(elfCalorieSums)
	maxSum, _ := util.Last(elfCalorieSums)
	return maxSum
	// or if you wanted to reverse sort
	//sort.Sort(sort.Reverse(sort.IntSlice(elfCalorieSums)))
	//return elfCalorieSums[0]
}

func part2(input string) int {
	elves := parseInput(input)
	elfCalorieSums := util.Map(elves, func(elfCalories []int) int {
		return util.Reduce(elfCalories, func(acc int, current int) int {
			return acc + current
		}, 0)
	})
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalorieSums)))
	topThreeElves := elfCalorieSums[:3]
	return util.SumSlice(topThreeElves)

}

func parseInput(input string) (ans [][]int) {
	grps := strings.Split(input, "\n\n")
	for _, grp := range grps {
		var row []int
		for _, line := range strings.Split(grp, "\n") {
			row = append(row, cast.ToInt(line))
		}
		ans = append(ans, row)
	}
	return ans
}
