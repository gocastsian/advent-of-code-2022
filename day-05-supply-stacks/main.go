package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []rune

func (s stack) Push(r rune) stack {
	return append(s, r)
}

func (s stack) Pop() (stack, rune) {
	// panic if stack is empty
	if len(s) == 0 {
		panic("stack is empty")
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) PopN(n int) (stack, []rune) {
	if len(s) < n {
		panic(fmt.Sprintf("stack length is smaller than %d", n))
	}

	l := len(s)
	return s[:l-n], s[l-n : l]
}

type move struct {
	From  int
	To    int
	Count int
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	stacks := initializeStacks()
	moves := readMoves("./moves")

	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			var lastItem rune
			stacks[m.From-1], lastItem = stacks[m.From-1].Pop()
			stacks[m.To-1] = stacks[m.To-1].Push(lastItem)
		}
	}

	for i := 0; i < 9; i++ {
		_, v := stacks[i].Pop()
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

func partTwo() {
	stacks := initializeStacks()
	moves := readMoves("./moves")

	for _, m := range moves {
		lastItems := make([]rune, m.Count)
		stacks[m.From-1], lastItems = stacks[m.From-1].PopN(m.Count)
		for _, l := range lastItems {
			stacks[m.To-1] = stacks[m.To-1].Push(l)
		}
	}

	for i := 0; i < 9; i++ {
		_, v := stacks[i].Pop()
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

func initializeStacks() []stack {
	stacks := make([]stack, 9)
	input := make(map[int][]rune)

	/*
		    [C]             [L]         [T]
		    [V] [R] [M]     [T]         [B]
		    [F] [G] [H] [Q] [Q]         [H]
		    [W] [L] [P] [V] [M] [V]     [F]
		    [P] [C] [W] [S] [Z] [B] [S] [P]
		[G] [R] [M] [B] [F] [J] [S] [Z] [D]
		[J] [L] [P] [F] [C] [H] [F] [J] [C]
		[Z] [Q] [F] [L] [G] [W] [H] [F] [M]
		 1   2   3   4   5   6   7   8   9
	*/

	input[0] = []rune{'Z', 'J', 'G'}
	input[1] = []rune{'Q', 'L', 'R', 'P', 'W', 'F', 'V', 'C'}
	input[2] = []rune{'F', 'P', 'M', 'C', 'L', 'G', 'R'}
	input[3] = []rune{'L', 'F', 'B', 'W', 'P', 'H', 'M'}
	input[4] = []rune{'G', 'C', 'F', 'S', 'V', 'Q'}
	input[5] = []rune{'W', 'H', 'J', 'Z', 'M', 'Q', 'T', 'L'}
	input[6] = []rune{'H', 'F', 'S', 'B', 'V'}
	input[7] = []rune{'F', 'J', 'Z', 'S'}
	input[8] = []rune{'M', 'C', 'D', 'P', 'F', 'H', 'B', 'T'}

	for i := 0; i < 9; i++ {
		for _, item := range input[i] {
			stacks[i] = stacks[i].Push(item)
		}
	}

	return stacks
}

func readMoves(path string) []move {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var moves []move
	for fileScanner.Scan() {
		l := strings.Split(fileScanner.Text(), " ")
		if len(l) != 6 {
			panic("bad move input")
		}

		f, t, c := getMove(l)
		moves = append(moves, move{
			From:  f,
			To:    t,
			Count: c,
		})
	}

	return moves
}

func getMove(s []string) (int, int, int) {
	f, err := strconv.Atoi(s[3])
	if err != nil {
		panic(err)
	}

	t, err := strconv.Atoi(s[5])
	if err != nil {
		panic(err)
	}

	c, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}

	return f, t, c
}
