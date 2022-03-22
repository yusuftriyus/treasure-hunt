package main

import (
	"fmt"
	"treasure-hunt/treasure"
)

func main() {
	//########
	//#......#
	//#.###..#
	//#...#.##
	//#X#....#
	//########

	arr := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "X", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}

	fmt.Println("BEFORE")
	for _, row := range arr {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}

	i := 4
	j := 1

	treasure.MarkUp(arr, i, j)

	fmt.Println()
	fmt.Println("AFTER")
	for _, row := range arr {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}
