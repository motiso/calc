package calc

import "github.com/motiso/calc/strings"

func GetStringLength(str string) int {
	return strings.StringLength(str)
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
