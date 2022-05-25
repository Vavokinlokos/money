package reverse_interpreter

import (
	"fmt"
	"gitlab.com/valeriia_sokolnikova/money_words/custom_constants"
	"strings"
)

func ReverseInterpret(input string) (string, error) {
	splittedInput := strings.Split(input, " ")
	dozensFlag := false
	result := ""
	for i, v := range splittedInput {
		if dozensFlag {
			dozensFlag = false
			continue
		}
		if i != len(splittedInput)-1 {
			resultPart, flag, err := analise(v, splittedInput[i+1])
			if err != nil {
				return "", err
			}
			dozensFlag = flag
			result += resultPart
		}
	}
	return result, nil
}

func analise(value, next string) (string, bool, error) {
	if value == "ноль" {
		return "0", false, nil
	}
	for _, v := range custom_constants.Hryvnia {
		if v[:len(v)-1] == value {
			return ".", false, nil
		}
	}
	for _, v := range custom_constants.Thousands {
		if v[:len(v)-1] == value {
			return "", false, nil
		}
	}
	for i, v := range custom_constants.Hundreds {
		if v[:len(v)-1] == value {
			h := i + 1
			for i, v := range custom_constants.Dozens2 {
				if v[:len(v)-1] == next {
					return fmt.Sprintf("%d%d", h, i+1), true, nil
				}
			}
			for i, v := range custom_constants.Dozens {
				if v[:len(v)-1] == next {
					return fmt.Sprintf("%d%d", h, i+1), true, nil

				}
			}
			return fmt.Sprintf("%d0", h), false, nil
		}
	}
	for i, v := range custom_constants.Dozens {
		if v[:len(v)-1] == value {
			return fmt.Sprintf("%d", i+1), false, nil
		}
	}
	for i, v := range custom_constants.Dozens2 {
		if v[:len(v)-1] == value {
			return fmt.Sprintf("%d", i+1), false, nil
		}
	}
	for i, v := range custom_constants.Units {
		if v[:len(v)-1] == value {
			return fmt.Sprintf("%d", i+1), false, nil
		}
	}
	return "", false, custom_constants.GrammarError
}
