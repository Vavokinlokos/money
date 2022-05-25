package main

import (
	"gitlab.com/valeriia_sokolnikova/money_words/custom_error"
	"gitlab.com/valeriia_sokolnikova/money_words/interpreter"
	"testing"
)

type testPair struct {
	input  string
	output string
}

type testError struct {
	input string
	err   error
}

var testsStrings = []testPair{
	{"123.43", "сто двадцать три гривны сорок три копейки"},
	{"888888,88", "восемьсот восемьдесят восемь тысяч восемьсот восемьдесят восемь гривен восемьдесят восемь копеек"},
	{"0.0", "ноль гривен "},
	{"42", "сорок две гривны "},
	{"0", "ноль гривен "},
	{"1000000", "один миллион гривен"},
	{"45.568", ""},
	{"444444444445.568", ""},
}

var testErrors = []testError{
	{"45.65585", custom_error.BigCoinsNumberError},
	{"555555555", custom_error.BigNumberError},
	{"*", custom_error.IntConvertError},
}

func TestInterpret(t *testing.T) {
	for _, pair := range testsStrings {
		v, _ := interpreter.Interpret(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, pair := range testErrors {
		_, err := interpreter.Interpret(pair.input)
		if err != pair.err {
			t.Error(
				"For", pair.input,
				"expected", pair.err,
				"got", err,
			)
		}
	}
}
