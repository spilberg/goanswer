package main

import (
	"fmt"
)

func calc(a int, b int, o string) int {
	var ret int

	switch o {
	case "*":
		ret = a * b //fmt.Printf("*")
	case "/":
		ret = a / b //fmt.Printf("*")
	case "+":
		ret = a + b //fmt.Printf("*")
	case "-":
		ret = a - b //fmt.Printf("*")
	default:

	}
	return ret
}

func main() {
	fmt.Println("Привіт! Я маленький калькулятор, можу виконувати прості дії з операціями + - / *")

	fmt.Printf("Введіть перший операнд: ")
	var opA int
	_, _ = fmt.Scan(&opA)

	fmt.Printf("Введіть другий операнд: ")
	var opB int
	_, _ = fmt.Scan(&opB)

	fmt.Printf("Оберіть операцію [+ - / *]: ")
	var operation string
	_, _ = fmt.Scan(&operation)

	fmt.Printf("Результат: %d \n", calc(opA, opB, operation))

}
