package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_3_2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	c, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	nextSet := func(set []string, p int, f func(int) bool) []string {
		if len(set) == 1 {
			return set
		}
		var a []string
		var b []string
		acc := 0
		for _, v := range set {
			if v[p] == '1' {
				a = append(a, v)
				acc += 1
			} else {
				b = append(b, v)
				acc -= 1
			}
		}
		if f(acc) {
			return a
		}
		return b
	}

	lines := strings.Split(string(c), "\n")

	l := len(lines[0])

	oxy := *(&lines)
	cab := *(&lines)

	for i := 0; i < l; i++ {
		oxy = nextSet(oxy, i, func(i int) bool {
			return i >= 0
		})
		cab = nextSet(cab, i, func(i int) bool {
			return i < 0
		})
		if len(cab) == 1 && len(oxy) == 1 {
			oxy_, _ := strconv.ParseInt(oxy[0], 2, 64)
			cab_, _ := strconv.ParseInt(cab[0], 2, 64)
			fmt.Println("Res: ", oxy_*cab_)
		}
	}
}
