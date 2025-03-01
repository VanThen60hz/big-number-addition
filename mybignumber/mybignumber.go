package mybignumber

import (
	"log"
	"strings"
)

type MyBigNumber struct{}

func (m *MyBigNumber) Sum(str1, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	s1 := reverse(str1)
	s2 := reverse(str2)

	var result strings.Builder
	carry := 0

	for i := 0; i < len(s1); i++ {
		digit1 := int(s1[i] - '0')
		digit2 := 0
		if i < len(s2) {
			digit2 = int(s2[i] - '0')
		}

		total := digit1 + digit2 + carry
		currentDigit := total % 10
		carry = total / 10

		log.Printf("Step %d: %d + %d + carry(%d) = %d, write %d, carry %d",
			i+1, digit1, digit2, carry, total, currentDigit, carry)

		result.WriteByte(byte(currentDigit + '0'))
	}

	if carry > 0 {
		log.Printf("Final carry: %d", carry)
		result.WriteByte(byte(carry + '0'))
	}

	finalResult := reverse(result.String())
	log.Printf("Final result: %s", finalResult)
	return finalResult
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
