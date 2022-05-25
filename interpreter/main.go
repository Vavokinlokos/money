package interpreter

import (
	"gitlab.com/valeriia_sokolnikova/money_words/custom_constants"
	"strconv"
	"strings"
)

func Interpret(input string) (string, error) {
	inputSlice, err := split(input)
	if err != nil {
		return "", err
	}
	integer := inputSlice[0]
	fractional := inputSlice[1]
	if integer == 1000000 {
		return "один миллион гривен", nil
	}
	if integer > 1000000 {
		return "", custom_constants.BigNumberError
	}
	if fractional > 99 {
		return "", custom_constants.BigCoinsNumberError
	}
	triads := getTriads(integer)
	integerRes := make([]string, 0)
	if len(triads) == 1 {
		integerRes = convertToWords(triads[0])
	}
	if len(triads) == 2 {
		integerRes = convertToWords(triads[0])
		integerRes = append(integerRes, custom_constants.Thousands[declension(triads[0]%10)])
		for _, word := range convertToWords(triads[1]) {
			integerRes = append(integerRes, word)
		}
	}
	var coin string
	var hr string
	if fractional/10 > 0 || (fractional < 20 && fractional >= 10) {
		coin = sliceToStr(convertToWords(fractional)) + custom_constants.Coins[declension((fractional%10)-1)]
	} else {
		if fractional == 0 {
			coin = ""
		} else {
			coin = custom_constants.Dozens[fractional-1] + custom_constants.Coins[2]
		}
	}

	if integer%100 < 20 && integer%100 >= 10 {
		hr = custom_constants.Hryvnia[2]
	} else {
		hr = custom_constants.Hryvnia[declension(triads[0]%10)]
	}
	if integerRes == nil {
		integerRes = append(integerRes, "ноль ")
	}
	return sliceToStr(integerRes) + hr + coin, nil

}

func convertPartsToInt(splittedInput []string) ([]int, error) {
	res := make([]int, 0, 2)
	integer, err := strconv.Atoi(splittedInput[0])
	if err != nil {
		return nil, custom_constants.IntConvertError
	}
	fractional, err := strconv.Atoi(splittedInput[1])
	if err != nil {
		return nil, custom_constants.FractConvertError
	}
	res = append(res, integer)
	res = append(res, fractional)
	return res, err
}

func sliceToStr(slice []string) string {
	return strings.Join(slice[:], "")
}

func split(input string) ([]int, error) {
	res := make([]int, 0)
	var err error
	comaSeparator := strings.Index(input, ",")
	splittedInput := make([]string, 0)
	switch {
	case comaSeparator < 0:
		pointSeparator := strings.Index(input, ".")
		switch {
		case pointSeparator < 0:
			splittedInput = append(splittedInput, input)
			splittedInput = append(splittedInput, "0")
			res, err = convertPartsToInt(splittedInput)
			if err != nil {
				return nil, err
			}
		default:
			splittedInput = strings.Split(input, ".")
			res, err = convertPartsToInt(splittedInput)
			if err != nil {
				return nil, err
			}
		}
	default:
		splittedInput = strings.Split(input, ",")
		res, err = convertPartsToInt(splittedInput)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func getTriads(num int) []int {
	res := make([]int, 0)
	if num == 0 {
		res = append(res, 0)
	}
	for num > 0 {
		if num < 1000 {
			res = append(res, num)
			return res
		}
		res = append(res, num/1000)
		num -= num / 1000 * 1000
	}
	return res
}

func convertToWords(num int) []string {
	var res []string
	for num > 0 {
		switch {
		case num < 9:
			unit := num % 10
			res = append(res, custom_constants.Units[unit-1])
			num = num - unit
		case num < 20:
			i := num % 10
			res = append(res, custom_constants.Dozens2[i-1])
			num = 0
		case num < 99:
			i := num / 10
			res = append(res, custom_constants.Dozens[i-1])
			num = num - i*10
		default:
			i := num / 100
			res = append(res, custom_constants.Hundreds[i-1])
			num = num - i*100
		}
	}
	return res
}

func declension(num int) int {
	switch num {
	case 1:
		return 0
	case 2, 3, 4:
		return 1
	default:
		return 2
	}
}
