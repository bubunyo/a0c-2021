package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day_3_1/input_2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	l := 0

	scanner := bufio.NewScanner(file)
	var mon []int

	for scanner.Scan() {
		if l == 0 {
			l = len(scanner.Text())
			mon = make([]int, l)
		}
		for k, v := range scanner.Text() {
			if v == '1' {
				mon[k] += 1
			} else {
				mon[k] -= 1
			}
		}
	}

	gamma := 0
	for i, v := range mon {
		if v > 0 {
			gamma += 1
		}
		if i != len(mon)-1 {
			gamma <<= 1
		}
	}

	fmt.Print("Res: ", gamma*(gamma^((1<<l)-1)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
