package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := readFileInput()
	list_1, list_2 := parseData(data)

	sort.Ints(list_1)
	sort.Ints(list_2)

	fmt.Println(len(list_1), len(list_2))

	sum := 0
	for i := range list_1 {
		a := float64(list_1[i])
		b := float64(list_2[i])
		sum += int(math.Abs(a - b))
	}

	fmt.Println("Part 1: ", sum)

	occurences := make([]int, list_2[len(list_2)-1]+1)
	for _, n := range list_2 {
		occurences[n] += 1
	}

	similarity := 0

	for _, n := range list_1 {
		similarity += n * occurences[n]
	}

	fmt.Println("Part 2: ", similarity)
}

func readFileInput() string {
	data, err := os.ReadFile("./input.txt")
	nilOrPanic(err)
	return string(data)
}

func parseData(d string) ([]int, []int) {
	list_1 := []int{}
	list_2 := []int{}

	lines := strings.Split(d, "\n")
	for _, line := range lines {
		if line != "" {
			re := regexp.MustCompile(`\d+`)
			numbers := re.FindAllString(line, -1)
			if len(numbers) != 2 {
				panic(fmt.Errorf("Expected two numbers"))
			}

			left, err := strconv.Atoi(numbers[0])
			nilOrPanic(err)
			right, err := strconv.Atoi(numbers[1])
			nilOrPanic(err)

			list_1 = append(list_1, left)
			list_2 = append(list_2, right)
		}
	}

	return list_1, list_2
}

func nilOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}