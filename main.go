package main

import (
	"fmt"
	"gitlab.com/valeriia_sokolnikova/money_words/interpreter"
	"gitlab.com/valeriia_sokolnikova/money_words/reverse_interpreter"
)

func main() {
	var input string
	fmt.Println("Введите сумму, которую необходимо вывести словами >> ")
	fmt.Scan(&input)
	result, err := interpreter.Interpret(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println("Введите словами сумму, которую необходимо вывести цифрой >> ")
	fmt.Scanln(&input)
	input = "четыреста пятьдесят три гривны пятьдесят копеек"
	reverseResult, err := reverse_interpreter.ReverseInterpret(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reverseResult)

	return
}
