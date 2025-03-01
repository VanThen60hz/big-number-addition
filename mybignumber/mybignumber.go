package mybignumber

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type MyBigNumber struct{}

func init() {
	// Configure logrus
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

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

		log.WithFields(log.Fields{
			"step":          i + 1,
			"digit1":        digit1,
			"digit2":        digit2,
			"carry_in":      carry,
			"total":         total,
			"current_digit": currentDigit,
			"carry_out":     carry,
		}).Info("Processing step")

		result.WriteByte(byte(currentDigit + '0'))
	}

	if carry > 0 {
		log.WithField("final_carry", carry).Info("Adding final carry")
		result.WriteByte(byte(carry + '0'))
	}

	finalResult := reverse(result.String())
	log.WithFields(log.Fields{
		"input1":       str1,
		"input2":       str2,
		"final_result": finalResult,
	}).Info("Calculation completed")

	return finalResult
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
