package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	x  int64
	y  int64
	dx int64
	dy int64
	id string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateFabric() [][]int16 {
	base := make([][]int16, 1000)

	for row := range base {
		base[row] = make([]int16, 1000)
	}
	return base
}

func parseClaim(claim string) Claim {
	parts := strings.Split(claim, " ")
	start := strings.Split(parts[2], ",")
	x, err := strconv.ParseInt(start[0], 10, 64)
	check(err)
	y, err := strconv.ParseInt(strings.TrimRight(start[1], ":"), 10, 64)
	check(err)
	distance := strings.Split(parts[3], "x")
	dx, err := strconv.ParseInt(distance[0], 10, 64)
	check(err)
	dy, err := strconv.ParseInt(distance[1], 10, 64)
	check(err)
	return Claim{x, y, dx, dy, parts[0]}
}

func isClaimWhole(fabric [][]int16, claim Claim) bool {
	var yOffset, xOffset int64
	for yOffset = 0; yOffset < claim.dy; yOffset++ {
		for xOffset = 0; xOffset < claim.dx; xOffset++ {
			if fabric[claim.y+yOffset][claim.x+xOffset] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fabric := generateFabric()
	var claims []Claim

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		claim := parseClaim(line)
		claims = append(claims, claim)

		var yOffset int64
		var xOffset int64

		for yOffset = 0; yOffset < claim.dy; yOffset++ {
			for xOffset = 0; xOffset < claim.dx; xOffset++ {
				fabric[claim.y+yOffset][claim.x+xOffset]++
			}
		}
	}

	for _, claim := range claims {
		if isClaimWhole(fabric, claim) {
			fmt.Printf("Claim %s is whole", claim.id)
		}
	}
}
