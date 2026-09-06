package assembler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/* Converts a decimal integer to a string of containing its 15-bit binary representation */
func DecimalAsBinary(decimal int) (string, error) {
	if decimal < 0 || decimal > 32767 {
		return "", errors.New("Decimal " + strconv.Itoa(decimal) + " out of range")
	}
	return fmt.Sprintf("%015b", decimal), nil
}

/* Checks if a string is a decimal number (no signs) */
func isNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) == false {
			return false
		}
	}
	return true
}

/* Splits a string into a series of lines and removes whitespace */
func GetLines(file string) []string {
	lines := strings.Split(file, "\n")
	cleanedLines := make([]string, 0, len(lines))
	for _, line := range lines {
		beforeComments, _, _ := strings.Cut(line, "//")
		removeSpaces := strings.ReplaceAll(beforeComments, " ", "")
		removeTabs := strings.ReplaceAll(removeSpaces, "\t", "")
		if removeTabs != "" && removeTabs != "\n" {
			cleanedLines = append(cleanedLines, removeTabs)
		}
	}
	return cleanedLines
}
