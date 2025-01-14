package utils

import "fmt"

func PrintHexArray(arr []uint8) {
	fmt.Println(ToHexArray(arr, false))
}

func ToHexArray(arr []uint8, noParenthesis ...bool) string {
	out := ""
	if len(noParenthesis) == 0 {
		out += "["
	}
	for i, b := range arr {
		out += fmt.Sprintf("%02x", b)
		if i != len(arr)-1 {
			out += " "
		}
	}
	if len(noParenthesis) == 0 {
		out += "]"
	}
	return out
}

func ToHexArray_2byte(arr []uint16, noParenthesis ...bool) string {
	arr8 := make([]byte, len(arr)*2)
	offset := 0
	for i := 0; i < len(arr); i++ {
		arr8[offset] = uint8(arr[i])
		offset++
		arr8[offset] = uint8(arr[i] >> 8)
		offset++
	}
	return ToHexArray(arr8, noParenthesis...)
}

func ToHexArray_1byte(arr []uint16, noParenthesis ...bool) string {
	arr8 := make([]byte, len(arr))
	for i := range arr8 {
		arr8[i] = uint8(arr[i])
	}
	return ToHexArray(arr8, noParenthesis...)
}
