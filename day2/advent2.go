package main

import (
	"bufio"
	"fmt"
	"os"
)

func printMap(m map[rune]int) {
	for key, value := range m {
		fmt.Printf("%c - %d\n", key, value)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	twoCounts := 0
	threeCounts := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		counts := make(map[rune]int)

		for _, character := range line {
			counts[character]++
		}

		printMap(counts)

		foundTwo := false
		foundThree := false
		for _, value := range counts {
			if value == 3 && !foundThree {
				threeCounts++
				foundThree = true
			}
			if value == 2 && !foundTwo {
				twoCounts++
				foundTwo = true
			}
		}
	}

	fmt.Printf("Checksum: %d", twoCounts*threeCounts)
}
