package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	rStr := []rune(str)
	var fmtStr strings.Builder
	for i := 0; i < len(rStr); i++ {
		if unicode.IsDigit(rStr[i]) { // Если число, то строка корректная
			return "", ErrInvalidString
		}

		if rStr[i] == '\\' { // Проверка на экранирование
			if (i + 1) >= len(rStr) { // Следующий символ должен быть
				return "", ErrInvalidString
			}
			i++
			fmtStr.WriteRune(rStr[i])
			continue
		}

		if (i+1) < len(rStr) && unicode.IsDigit(rStr[i+1]) { // Проверка на повторения
			i++
			fmtStr.WriteString(strings.Repeat(string(rStr[i-1]), int(rStr[i]-'0')))
		} else {
			fmtStr.WriteRune(rStr[i])
		}
	}
	return fmtStr.String(), nil
}
