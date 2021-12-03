package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_2_1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x, y := 0, 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n := strings.Split(scanner.Text(), " ")
		dir := n[0]

		change, err := strconv.Atoi(n[1])
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "forward":
			x += change
		case "down":
			y += change
		case "up":
			y -= change
		default:
			log.Fatal("invalid dir")
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(x * y)
}
