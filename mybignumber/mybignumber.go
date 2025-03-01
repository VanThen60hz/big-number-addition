package mybignumber

import (
	log "github.com/sirupsen/logrus"
)

type MyBigNumber struct{}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

func (m *MyBigNumber) Sum(str1, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	len1, len2 := len(str1), len(str2)
	result := make([]byte, len1+1)
	currentPos := len(result) - 1

	carry := 0
	for i := 0; i < len1; i++ {
		digit1 := int(str1[len1-1-i] - '0')
		digit2 := 0
		if i < len2 {
			digit2 = int(str2[len2-1-i] - '0')
		}

		total := digit1 + digit2 + carry
		currentDigit := total % 10
		carry = total / 10

		log.WithFields(log.Fields{
			"step":          i + 1,
			"position":      len1 - i,
			"digit1":        digit1,
			"digit2":        digit2,
			"carry_in":      carry,
			"total":         total,
			"current_digit": currentDigit,
			"carry_out":     carry,
		}).Info("Processing step")

		result[currentPos] = byte(currentDigit + '0')
		currentPos--
	}

	if carry > 0 {
		log.WithField("final_carry", carry).Info("Adding final carry")
		result[currentPos] = byte(carry + '0')
	} else {
		result = result[1:]
	}

	finalResult := string(result)
	log.WithFields(log.Fields{
		"input1":       str1,
		"input2":       str2,
		"final_result": finalResult,
	}).Info("Calculation completed")

	return finalResult
}
