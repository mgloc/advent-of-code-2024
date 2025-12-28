package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const input_file string = "i.txt"

func part1() {
	program, _ := os.ReadFile(input_file)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(string(program), -1)

	sum := 0
	for _, v := range matches {
		x, _ := strconv.Atoi(v[1])
		y, _ := strconv.Atoi(v[2])
		sum += x * y
	}
	fmt.Println(sum)
}

func part2() {
	program, _ := os.ReadFile(input_file)
	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	matches := re.FindAllString(string(program), -1)

	sum := 0
	activated := true
	for _, v := range matches {
		if v == "don't()" {
			activated = false
		} else if v == "do()" {
			activated = true
		} else {
			if !activated {
				continue
			}
			var a, b int
			fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
			sum += a * b
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
