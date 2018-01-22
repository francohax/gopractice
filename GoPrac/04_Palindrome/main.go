package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give me a number")
	fmt.Println("---------------------")

	for {
		fmt.Println("->")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		test, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("ERROR ", err)
		}

		isPalindrome(test)
	}
}

func isPalindrome(x int) {
	p := strconv.Itoa(x)
	p = reverse(p)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
