package main

import (
	"advent2022/util"
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	var lowerPat = regexp.MustCompile(`[a-z]`)
	var upperPat = regexp.MustCompile(`[A-Z]`)
	lowerArr := "_abcdefghijklmnopqrstuvwxyz"
	upperArr := "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	ans := 0

	for _, line := range strings.Split(input, "\n") {
		setA := map[string]struct{}{}
		setB := map[string]struct{}{}
		inter := make(map[string]bool)
		hash := make(map[string]bool)

		pt := len(line) / 2
		firstHalf := line[:pt]
		secondHalf := line[pt:]

		for _, char := range firstHalf {
			setA[string(char)] = struct{}{}
		}
		for _, char := range secondHalf {
			setB[string(char)] = struct{}{}
		}

		for str := range setA {
			hash[str] = true
		}
		for str := range setB {
			if _, ok := inter[str]; hash[str] && !ok {
				inter[str] = true
				hash[str] = false
			}
		}

		for s := range inter {
			switch {
			case lowerPat.MatchString(s):
				index := bytes.IndexByte([]byte(lowerArr), []byte(s)[0])
				if index > 0 {
					ans += index
				}

			case upperPat.MatchString(s):
				index := bytes.IndexByte([]byte(upperArr), []byte(s)[0])
				if index > 0 {
					ans += index + 26
				}
			}
		}
	}

	return ans
}

func part2(input string) int {
	return 0
}
