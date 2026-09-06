package main

import (
	"fmt"
	"os"
	"strings"
	. "HackAssembler/assembler"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: incorrect number of arguments")
		return
	}

	var inputFilePath string = os.Args[1]
	fileName, isAsmFile := strings.CutSuffix(os.Args[1], ".asm")
	if !isAsmFile {	
		fmt.Println("Error: not an assembly file")
		return
	}

	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var lines []string = GetLines(string(input))
	var instructions []Instruction = Parse(lines)
	output, err := ConvertInstructions(instructions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var outputFilePath string = fileName + ".hack"
	err = os.WriteFile(outputFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("DONE")
}

