package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	var grid = make([][]int, n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		grid[i] = make([]int, m)
		for j, char := range s {
			grid[i][j] = int(char - '0')
		}
	}
	rowSummary := make([]int, m)
	for i := 0; i < n; i++ {
		rowSummary[i] = smallest(grid[i]...)
	}
	colSummary := make([]int, n)
	for i := 0; i < n; i++ {
		colSummary[i] = 5
		for j := 0; j < m; j++ {
			colSummary[j] = smallest(colSummary[j], grid[i][j])
		}
	}
	//worst := gridSmallest(rowSummary, colSummary)

}

func smallest(nums ...int) int {
	rez := nums[0]
	for _, num := range nums {
		if num < rez {
			rez = num
		}
	}
	return rez
}

func smallestWithoutI(summary []int, index int) int {
	rez := 5
	for i, num := range summary {
		if i != index && num < rez {
			rez = num
		}
	}
	return rez
}

func gridSmallest(rowSummary, colSummary []int) int {
	rowSmallest := smallest(rowSummary...)
	colSmallest := smallest(colSummary...)
	rez := smallest(rowSmallest, colSmallest)
	return rez
}

func gridSmallestWithoutIJ(rowSummary, colSummary []int, i, j int) int {
	rowSmallest := smallestWithoutI(rowSummary, j)
	colSmallest := smallestWithoutI(colSummary, i)
	rez := smallest(rowSmallest, colSmallest)
	return rez
}
