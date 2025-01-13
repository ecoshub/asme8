package utils

import "fmt"

func PrintHexArray(arr []uint8) {
	fmt.Print("[")
	for i, b := range arr {
		fmt.Printf("%02x", b)
		if i != len(arr)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Println("]")
}
