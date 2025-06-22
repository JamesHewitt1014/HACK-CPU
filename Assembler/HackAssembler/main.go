package main

import (
	"fmt"
	"os"
	"strings"
)

//TODO: 
// 1. Add a CLI
// 2. Testing

func main() {
	var inputFilePath string = "test.asm"
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var lines []string = GetLines(string(input))
	var instructions []Instruction = Parse(lines)
	output, err := convertInstructions(instructions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var outputFilePath string = "output.asm"
	err = os.WriteFile(outputFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("DONE")
}

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
