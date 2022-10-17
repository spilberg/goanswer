package main

import (
	"fmt"
)

func main() {
	fmt.Println("Добрий день! Я програма 'лікарь'")
	fmt.Printf("Як вас звати? ")
	var name string
	_, _ = fmt.Scan(&name)
	fmt.Printf("Я радий вас вітати %s \n", name)

	fmt.Printf("Скільки вам років? ")
	var age int
	_, _ = fmt.Scan(&age)

	

	fmt.Printf("На що скаржитесь? (голова, ноги, спина)")
	var scared string
	_, _ = fmt.Scan(&scared)


	if (age > 10 && age < 25) {
		fmt.Printf("В такому віці вам скоріше за все лікар не потрібен :-) Насолоджуйтесь життям!")
	}else{
		fmt.Printf("Саме час потуруватися про здоров'я\n")

		switch scared {
	case "голова":
		fmt.Printf("Вам потрібно звернутися до невтропатолога ")
	case "ноги":
		fmt.Printf("Цією проблемою займається фліболог")
	case "спина":
		fmt.Printf("Ага болить спина, лікувальний масаж вам допоможе ")
	default: 
		fmt.Printf("З такими скаргами ми нажаль не лікуємо, взерніться в іншу поліклініку :-)")
	}

	}

	

}