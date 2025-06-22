package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var dest = map[string]string{
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"ADM": "111",
}

var jump = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

// NOTE: 
// First bit indicates use of Register A or Data Memory input for ALU
// The other bits are the control bits for the ALU
var comp = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"!M":  "1110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"-M":  "1110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

func convertInstructions(instructions []Instruction) (string, error){
	var sb strings.Builder
	for _, instruction := range instructions {
		machineCode, err := instruction.Convert()
		if err != nil {
			return "", err
		}
		sb.WriteString(machineCode + "\n")	
	}
	return sb.String(), nil
}

func convertDest(mnemonic string) (string, error) {
	if mnemonic == ""{
		return "000", nil
	}
	bits, found := dest[mnemonic]
	if found == false {
		return "", errors.New("invalid input destination: " + mnemonic)
	}
	return bits, nil
}

func convertComp(mnemonic string) (string, error) {
	bits, found := comp[mnemonic]
	if found == false {
		return "", errors.New("invalid input computation: " + mnemonic)
	}
	return bits, nil
}

func convertJump(mnemonic string) (string, error){
	if mnemonic == "" {
		return "000", nil
	}
	bits, found := jump[mnemonic]
	if found == false {
		return "", errors.New("invalid input jump: " + mnemonic)
	}
	return bits, nil
}

func (i CInstruction) Convert() (string, error) {
	destbits, err := convertDest(i.Dest)
	if err != nil {
		return "", err
	}
	compbits, err := convertComp(i.Comp)
	if err != nil {
		return "", err
	}
	jumpbits, err := convertJump(i.Jump)
	if err != nil {
		return "", err
	}
	return "111" + compbits + destbits + jumpbits, nil
}

func (i AInstruction) Convert() (string, error) {
	if isNumber(i.Value){
		decimal, _ := strconv.Atoi(i.Value)	
		binary, err := DecimalAsBinary(decimal)
		if err != nil {
			return "", err
		}
		return "1" + binary, nil
	}
	
	if address, found := Symbols[i.Value]; found {
		binary, err := DecimalAsBinary(address)
		if err != nil {
			return "", err
		}
		return "1" + binary, nil
	}

	address := AddSymbol(i.Value)
	binary, err := DecimalAsBinary(address)
	if err != nil {
		return "", err
	}
	return "1" + binary, nil
}

// Converts a decimal integer to a string of containing its 15-bit binary representation
func DecimalAsBinary(decimal int) (string, error){
	if decimal < 0 || decimal > 32767 {
		return "", errors.New("Decimal " + strconv.Itoa(decimal) +" out of range")
	}
	return fmt.Sprintf("%015b", decimal), nil
}

// Checks if a string is a decimal number (no signs)
func isNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) == false{
			return false
		}
	}
	return true
}

























