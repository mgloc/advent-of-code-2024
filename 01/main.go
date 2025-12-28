package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	return max(x, -x)
}

func get_input() ([]int, []int) {
	var left_list []int
	var right_list []int

	file, _ := os.Open("i.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])
		left_list = append(left_list, a)
		right_list = append(right_list, b)
	}
	return left_list, right_list
}

func part1() {
	left_list, right_list := get_input()

	sort.Ints(left_list)
	sort.Ints(right_list)

	sum := 0
	for i := 0; i < len(right_list); i++ {
		distance := abs(left_list[i] - right_list[i])
		sum += distance
	}

	fmt.Println("Part 1:")
	fmt.Println(sum)
}

func part2() {
	left_list, right_list := get_input()

	sort.Ints(left_list)
	sort.Ints(right_list)

	index_left := 0
	index_right := 0
	sum := 0
	n := len(left_list)

	for index_left < n && index_right < n {
		if right_list[index_right] < left_list[index_left] {
			index_right++
			continue
		}

		if right_list[index_right] > left_list[index_left] {
			index_left++
			continue
		}

		left_occ := 0
		right_occ := 0
		value := right_list[index_right]

		for index_right < n && right_list[index_right] == value {
			index_right++
			right_occ++
		}

		for index_left < n && left_list[index_left] == value {
			index_left++
			left_occ++
		}

		sum += left_occ * right_occ * value
	}

	fmt.Println("Part 2:")
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
