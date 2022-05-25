package custom_constants

import (
	"errors"
)

var BigNumberError = errors.New("НЕПРАВИЛЬНЫЙ ВВОД. Программа предназначена для обработки значений до 1 000 000 гривен")
var BigCoinsNumberError = errors.New("НЕПРАВИЛЬНЫЙ ВВОД. Значение копеек не может превышать 99")
var IntConvertError = errors.New("failed converting integer part to int")
var FractConvertError = errors.New("failed converting fractional part to int")
var GrammarError = errors.New("НЕПРАВИЛЬНЫЙ ВВОД. Проверьте рамматику")
