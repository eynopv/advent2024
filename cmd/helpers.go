package helpers

import (
	"os"
	"strconv"
)

func ReadFileInput(f string) string {
	data, err := os.ReadFile(f)
	NilOrPanic(err)
	return string(data)
}

func NilOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return number
}
