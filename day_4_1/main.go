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
	file, err := os.Open("day_4_1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	draws := strings.Split(scanner.Text(), ",")

	scanner.Scan() // scanning again to avoid next entry line

	type d struct{ r, c int }

	board, row := 0, 0

	// this watches the score for each board
	scoreWatch := map[int]map[string]map[int]int{board: {"r": {}, "c": {}}}
	boardNos := map[int][]string{}

	//              map[entry]map[board]d{r, c}
	entries := make(map[string][]map[int]d)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			row = 0
			board += 1
			scoreWatch[board] = map[string]map[int]int{"r": {}, "c": {}}
			continue
		}

		for col, v := range strings.Split(strings.TrimSpace(strings.ReplaceAll(scanner.Text(), "  ", " ")), " ") {
			boards := entries[v]
			boards = append(boards, map[int]d{board: {row, col}})
			entries[v] = boards
			boardNos[board] = append(boardNos[board], v)
		}
		row += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	calledNo := map[string]struct{}{}

	called := func(v string, b int) {
		n, _ := strconv.Atoi(v)
		acc := 0
		for _, ns := range boardNos[b] {
			if _, ok := calledNo[ns]; ok {
				continue
			}
			vv, _ := strconv.Atoi(ns)
			acc += vv
		}
		fmt.Println("Res: ", n*acc)
	}

	for _, v := range draws {
		for _, m := range entries[v] {
			for b, e := range m {
				calledNo[v] = struct{}{}

				scoreWatch[b]["r"][e.r] += 1
				scoreWatch[b]["c"][e.c] += 1

				if scoreWatch[b]["r"][e.r] == 5 || scoreWatch[b]["c"][e.c] == 5 {
					called(v, b)
					os.Exit(0)
				}
			}
		}
	}

}
