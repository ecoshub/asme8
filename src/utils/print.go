package utils

import "fmt"

func PrintHexArray(arr []uint8) {
	fmt.Println(ToHexArray(arr))
}

func ToHexArray(arr []uint8) string {
	out := "["
	for i, b := range arr {
		out += fmt.Sprintf("%02x", b)
		if i != len(arr)-1 {
			out += " "
		}
	}
	out += "]"
	return out
}
