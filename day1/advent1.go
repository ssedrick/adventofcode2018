package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func doOperation(operation byte, left int64, right int64) int64 {
	if operation == '-' {
		return left - right
	} else if operation == '+' {
		return left + right
	}
	return left
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	var frequency int64

	for err == nil {
		sign, err := reader.ReadByte()

		if err == io.EOF {
			break
		}

		numStr, err := reader.ReadString('\n')
		check(err)

		offset, err := strconv.ParseInt(numStr[:len(numStr)-1], 0, 0)
		check(err)

		frequency = doOperation(sign, frequency, offset)
	}

	fmt.Printf("frequency: %d\n", frequency)

}
