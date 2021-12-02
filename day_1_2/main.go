package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1_2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depth, i, count, prevV := 3, 0, 0, 0
	skip := true
	prev, curr := newQ(depth), newQ(depth)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if skip {
			prevV = n
			skip = false
			i++
			continue
		}
		prev.push(prevV)
		prevV = n
		curr.push(n)
		if i >= depth && prev.sum() < curr.sum() {
			count++
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

type q struct {
	i    int
	data []int
}

func newQ(limit int) *q {
	return &q{
		data: make([]int, limit),
	}
}

func (q *q) push(d int) {
	q.data[q.i] = d
	q.i = (q.i + 1) % len(q.data)
}

func (q q) sum() int {
	acc := 0
	for _, v := range q.data {
		acc += v
	}
	return acc
}
