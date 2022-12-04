package main

import (
	"advent2022/util"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type RPS int

const (
	Rock RPS = iota
	Paper
	Scissor
)

type GameResult int

const (
	Win GameResult = iota
	Loss
	Draw
)

type RPSGame struct {
	opponent RPS
	player   RPS
	score    int
	result   GameResult
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
	var games []RPSGame

	for _, line := range strings.Split(input, "\n") {
		game := parseGame(line)
		game.parse()
		games = append(games, game)
	}

	return util.Reduce(games, func(acc int, current RPSGame) int {
		//fmt.Printf("%+v\n", current)
		return acc + current.score
	}, 0)

}

func part2(input string) int {
	var games []RPSGame
	for _, line := range strings.Split(input, "\n") {
		game := parseGame2(line)
		game.parse()
		games = append(games, game)
	}
	return util.Reduce(games, func(acc int, current RPSGame) int {
		//fmt.Printf("%+v\n", current)
		return acc + current.score
	}, 0)
}

func parseGame(g string) RPSGame {
	fields := strings.Fields(g)
	var game RPSGame
	switch fields[0] {
	case "A":
		game.opponent = Rock
	case "B":
		game.opponent = Paper
	case "C":
		game.opponent = Scissor
	}
	switch fields[1] {
	case "X":
		game.player = Rock
	case "Y":
		game.player = Paper
	case "Z":
		game.player = Scissor
	}

	return game
}

func parseGame2(g string) RPSGame {
	fields := strings.Fields(g)
	var game RPSGame
	switch fields[0] {
	case "A":
		game.opponent = Rock
	case "B":
		game.opponent = Paper
	case "C":
		game.opponent = Scissor
	}

	game.player = playResult(game.opponent, fields[1])

	return game
}

func playResult(opponent RPS, instruction string) RPS {
	var result RPS

	if instruction == "Y" {
		result = opponent
	} else {
		switch opponent {
		case Rock:
			if instruction == "X" {
				result = Scissor
			} else {
				result = Paper
			}
		case Paper:
			if instruction == "X" {
				result = Rock
			} else {
				result = Scissor
			}

		case Scissor:
			if instruction == "X" {
				result = Paper
			} else {
				result = Rock
			}
		}
	}
	return result
}

func (r GameResult) score() int {
	switch r {
	case Win:
		return 6
	case Loss:
		return 0
	case Draw:
		return 3
	}
	return 0
}

func (r RPS) score() int {
	switch r {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	}
	return 0
}

func (g *RPSGame) parse() {
	if g.opponent == g.player {
		g.result = Draw
		g.score = g.result.score() + g.player.score()
	} else {
		match := [2]RPS{g.opponent, g.player}
		switch match {
		case [2]RPS{Rock, Paper}:
			g.result = Win
		case [2]RPS{Rock, Scissor}:
			g.result = Loss
		case [2]RPS{Paper, Rock}:
			g.result = Loss
		case [2]RPS{Paper, Scissor}:
			g.result = Win
		case [2]RPS{Scissor, Rock}:
			g.result = Win
		case [2]RPS{Scissor, Paper}:
			g.result = Loss
		}
	}
	g.score = g.result.score() + g.player.score()
}
