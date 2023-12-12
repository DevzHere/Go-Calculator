package main

import (
	"fmt"
)

func main() {

	var opr string
	var a, b int
    fmt.Println("Go Calculator")
	fmt.Println("Value 1: ")
	fmt.Scan(&a)
	fmt.Println("Value 2: ")
	fmt.Scan(&b)
	fmt.Println("Operator: ")
	fmt.Scan(&opr)

	fmt.Println("Evaluating Expression: ", a, opr, b)

	switch string(opr) {
	case "+":
		fmt.Println(a+b)
	case "*", "x":
		fmt.Println(a*b)
	case "-":
		fmt.Println(a-b)
	case "/":
		fmt.Println(a/b)
	default:
		fmt.Println("Press correct operator")
	}
}


// package main

// import (
//     "fmt"
//     "regexp"
//     "strconv"
// )

// func is_num(word string) bool {
//     return regexp.MustCompile(`^[0-9]*$`).MatchString(word)
// }
// func calc(str string) int {
//     var i int
//     for i = 0; i < len(str); i++ {
//         if !is_num(string(str[i])) {
//             break
//         }
//     }
//     switch string(str[i]) {
//     case "x":
//         x, _  := strconv.Atoi(str[0:i])
//         y, _ := strconv.Atoi(str[(i + 1):len(str)])
//         return x * y
//     case "*":
//         x, _ := strconv.Atoi(str[0:i])
//         y, _ := strconv.Atoi(str[(i + 1):len(str)])
//         return x * y
//     case "+":
//         x, _ := strconv.Atoi(str[0:i])
//         y, _ := strconv.Atoi(str[(i + 1):len(str)])
//         return x + y
//     case "-":
//         x, _ := strconv.Atoi(str[0:i])
//         y, _ := strconv.Atoi(str[(i + 1):len(str)])
//         return x - y
//     case "/":
//         x, _ := strconv.Atoi(str[0:i])
//         y, _ := strconv.Atoi(str[(i + 1):len(str)])
//         return x / y
//     default:
//         return 0
//     }
// }

// func main() {
//     fmt.Println(calc("3_3333"))
// }