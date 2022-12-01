package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	rawData := readRawData("./input")

	elves := countCalories(rawData)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})

	fmt.Printf("the number %d Elf has most calories: %d\n", elves[0].Index, elves[0].Calories)

	fmt.Printf("total calories of the top three Elves: %d\n", elves[0].Calories+elves[1].Calories+elves[2].Calories)
}

type Elf struct {
	Index    int
	Items    int
	Calories int
}

func countCalories(data []string) []Elf {
	var elves = make([]Elf, 0)

	elfIndex := 0
	elves = append(elves, Elf{
		Index:    elfIndex,
		Items:    0,
		Calories: 0,
	})

	for _, d := range data {
		if d == "" {
			elfIndex++

			elves = append(elves, Elf{
				Index:    elfIndex,
				Items:    0,
				Calories: 0,
			})

			continue
		}

		calories, err := strconv.Atoi(d)
		if err != nil {
			log.Fatal(err)
		}

		elves[elfIndex].Calories += calories
		elves[elfIndex].Items++
	}

	fmt.Printf("total elf count: %d\n", elfIndex)

	return elves
}

func readRawData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var data []string
	for fileScanner.Scan() {
		data = append(data, fileScanner.Text())
	}

	return data
}
