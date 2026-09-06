package main

import (
	. "HackAssembler/assembler"
	"fmt"
)

func WriteAssembly(input string) (string, error) {	
	var lines []string = GetLines(string(input))
	var instructions []Instruction = Parse(lines)
	output, err := ConvertInstructions(instructions)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return output, nil
}
