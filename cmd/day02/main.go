package main

import (
	"fmt"
	"math"
	"strings"

	helpers "github.com/eynopv/advent2024/cmd"
)

func main() {
	data := helpers.ReadFileInput("input.txt")
	levels := parseData(data)

	part_1 := 0
	part_2 := 0
	for _, level := range levels {
		is_safe := checkSafe(level)

		if is_safe {
			part_1 += 1
			part_2 += 1
		} else {
			for i := 0; i < len(level); i++ {
				sublevel := removeIndex(level, i)
				is_safe = checkSafe(sublevel)

				if is_safe {
					part_2 += 1
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", part_1)
	fmt.Println("Part 2:", part_2)
}

func checkSafe(level []int) bool {
	is_safe := true
	previous_delta := level[1] - level[0]

	for i := 0; i < len(level)-1; i++ {
		current := level[i]
		next := level[i+1]
		delta := next - current

		if previous_delta < 0 && delta > 0 {
			is_safe = false
			break
		}

		if previous_delta > 0 && delta < 0 {
			is_safe = false
			break
		}

		previous_delta = delta

		delta_abs := int(math.Abs(float64(delta)))

		if delta_abs < 1 || delta_abs > 3 {
			is_safe = false
			break
		}
	}

	return is_safe
}

func parseData(d string) [][]int {
	levels := [][]int{}
	lines := strings.Split(d, "\n")

	for _, line := range lines {
		if line != "" {
			level := []int{}
			numbers := strings.Split(line, " ")
			for _, n := range numbers {
				number := helpers.ParseInt(n)
				level = append(level, number)
			}
			levels = append(levels, level)
		}
	}

	return levels
}

func removeIndex(s []int, i int) []int {
	return append(append([]int{}, s[:i]...), s[i+1:]...)
}
