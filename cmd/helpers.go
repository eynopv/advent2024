package helpers

import "os"

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
