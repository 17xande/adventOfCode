package main

import (
	"fmt"
	"os"
)

func main() {
	buf := input()
	startOfPacket := search(buf, 4)
	fmt.Printf("Start of Packet: %d\n", startOfPacket)
	startOfMessage := search(buf, 14)
	fmt.Printf("Start of Message: %d\n", startOfMessage)
}

func search(buf string, length int) int {
	length--
	for i := length; i < len(buf); i++ {
		s := buf[i-length : i+1]
		if containsDupes(s) {
			continue
		}
		return i + 1
	}

	return -1
}

func containsDupes(s string) bool {
	for _, r := range []rune(s) {
		count := 0
		for _, ir := range []rune(s) {
			if ir == r {
				count++
			}
		}
		if count > 1 {
			return true
		}
	}

	return false
}

func input() string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
