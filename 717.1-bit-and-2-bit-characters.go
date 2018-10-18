package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	i := 0
	for {
		if i == len(bits)-1 {
			return bits[i] == 0
		} else if i >= len(bits) {
			return false
		} else if bits[i] == 0 {
			i += 1
		} else if bits[i] == 1 {
			i += 2
		}
	}
}

func main() {
	bits1 := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(bits1))

	bits2 := []int{1, 1, 1, 0}
	fmt.Println(isOneBitCharacter(bits2))
}
