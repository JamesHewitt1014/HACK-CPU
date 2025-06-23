package main

import (
	"errors"
	"strings"
)

var Symbols = map[string]int{
	"R0": 0,
	"R1": 1,
	"R2": 2,
	"R3": 3,
	"R4": 4,
	"R5": 5,
	"R6": 6,
	"R7": 7,
	"R8": 8,
	"R9": 9,
	"R10": 10,
	"R11": 11,
	"R12": 12,
	"R13": 13,
	"R14": 14,
	"R15": 15,
	"SCREEN": 16384,
	"KBD": 24576,
	"SP": 0,
	"LCL": 1,
	"ARG": 2,
	"THIS": 3,
	"THAT": 4,
}

var counter int = 16

func AddLabel(line string, lineNumber int) (error){	
	// Remove "(" and everything after ")"
	label, found := strings.CutSuffix(line[1:], ")")
	if !found {
		return errors.New("label syntax incorrect: " + label)
	}
	if _, found := Symbols[label]; found {	
		return errors.New("label already exists: " + label)	
	}	
	label = label
	Symbols[label] = lineNumber
	return nil
}

func AddSymbol(symbol string) (int) {
	Symbols[symbol] = counter 
	counter++
	return Symbols[symbol]
}
