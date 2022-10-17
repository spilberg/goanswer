package main

import (
	"fmt"
	"flag"
)

type country string
type color string
type brand string

const (
	USA  country = "u"
	Ukraine country = "k"
	Germany country = "g"

	Green color = `g`
	Yellow color = `y`
	Red   color = `r`

	BMW brand = `b`
	Chevrolet brand = `c`
	Zaz brand = `z`
)


func main() {
	countr := flag.String("country",``,"Country")
	col := flag.String("color",``,"Color")
    br := flag.String("brand",``,"Brand")
	help := flag.Bool("help", false, "Help")
	flag.Parse()

	if *help {
		fmt.Print("Запустіть програму з параметрами country та color\n" +
			"де country може мати значення [u|k|g] \n" +
			"та color [g|y|r]\n" +
			"Приклад: go run main.go --country a --color")
	}else{
		switch *countr {
		case string(USA):
			fmt.Println("Країна виробництва: Сполучені Штати\n")
		case string(Ukraine):
			fmt.Println("Країна виробництва: Україна\n")
		case string(Germany):
			fmt.Println("Країна виробництва: Німетчина\n")
		default:
			fmt.Println("Країна виробництва: невідома\n")
		}

		switch *col {
		case string(Green):
			fmt.Println("Колір: Зелений\n")
		case string(Yellow):
			fmt.Println("Колір: Жовтий\n")
		case string(Red):
			fmt.Println("Колір: Червоний\n")
		default:
			fmt.Println("Колір: якийсь незрозумілий колір\n")
		}

		switch *br {
		case string(BMW):
			fmt.Println("Модель: БМВ\n")
		case string(Chevrolet):
			fmt.Println("Модель: Шевролет\n")
		case string(Zaz):
			fmt.Println("Модель: Запорожець\n")
		default:
			fmt.Println("Модель: невідома модель\n")
		}

	}

}