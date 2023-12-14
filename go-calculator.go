package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func is_num(word string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(word)
}

func calc(str string) int {
	var i int
	for i = 0; i < len(str); i++ {
		if !is_num(string(str[i])) {
			break
		}
	}
	switch string(str[i]) {
	case "x":
		x, _ := strconv.Atoi(str[0:i])
		y, _ := strconv.Atoi(str[(i + 1):len(str)])
		return x * y
	case "*":
		x, _ := strconv.Atoi(str[0:i])
		y, _ := strconv.Atoi(str[(i + 1):len(str)])
		return x * y
	case "+":
		x, _ := strconv.Atoi(str[0:i])
		y, _ := strconv.Atoi(str[(i + 1):len(str)])
		return x + y
	case "-":
		x, _ := strconv.Atoi(str[0:i])
		y, _ := strconv.Atoi(str[(i + 1):len(str)])
		return x - y
	case "/":
		x, _ := strconv.Atoi(str[0:i])
		y, _ := strconv.Atoi(str[(i + 1):len(str)])
		return x / y
	default:
		return 0
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(strings.ReplaceAll(text, " ", ""), "\r\n", "")
	return text
}

func main() {

	value := getUserInput()
	fmt.Println("Evaluating expression:", value)
	fmt.Println(calc(value))
}
