package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	file, err := os.Open("day_1_1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int

	scanner := bufio.NewScanner(file)
	previous := -1

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if previous == -1 {
			previous = i
			continue
		}
		if i > previous {
			count++
		}
		previous = i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

}
