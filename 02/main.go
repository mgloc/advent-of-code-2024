package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() [][]int {
	var reports [][]int

	file, _ := os.Open("i.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var report []int
		for _, v := range fields {
			digit, _ := strconv.Atoi(v)
			report = append(report, digit)
		}
		reports = append(reports, report)
	}
	return reports
}

func abs(x int) int {
	return max(x, -x)
}

func is_safe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	original_direction := report[0]-report[1] < 0

	for i := 0; i < len(report)-1; i++ {
		difference := report[i] - report[i+1]
		abs_difference := abs(difference)
		if abs_difference < 1 || abs_difference > 3 {
			return false
		}
		if difference < 0 != original_direction {
			return false
		}
	}
	return true

}

func part1() {
	reports := get_input()
	sum := 0
	for _, report := range reports {
		if is_safe(report) {
			sum++
		}
	}
	fmt.Println(sum)
}

func duplicate_and_remove[T any](index int, slice []T) []T {
	result := make([]T, 0, len(slice))
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func is_safe_2(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	original_direction := report[0]-report[1] < 0

	left := 0
	right := 1
	n := len(report)

	for right < n {
		difference := report[left] - report[right]
		abs_difference := abs(difference)

		if abs_difference < 1 || abs_difference > 3 || difference < 0 != original_direction {
			without_left := duplicate_and_remove(left, report)
			without_right := duplicate_and_remove(right, report)
			return is_safe(without_left) || is_safe(without_right) || is_safe(without_right) || is_safe(report[1:])
		}

		left = right
		right++
	}
	return true

}

func part2() {
	reports := get_input()
	sum := 0
	for _, report := range reports {
		if is_safe_2(report) {
			fmt.Println(report)
			sum++
		}
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}
