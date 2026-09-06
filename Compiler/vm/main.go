package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// for val := range os.Args {
	// 	fmt.Sprintln("%s", val)
	// }

	// if len(os.Args) != 1 {
	// 	fmt.Println("Error: incorrect number of arguments")
	// 	return
	// }

	var inputFilePath string = os.Args[1]
	fileName, isVMFile := strings.CutSuffix(os.Args[1], ".vm")
	if !isVMFile {	
		fmt.Println("Error: not a VM file")
		return
	}

	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	stringInput := string(input)

	fmt.Print(input)

	output := translate(stringInput, fileName)
	fmt.Println(output)

	var outputFilePath string = fileName + ".asm"
	err = os.WriteFile(outputFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
}
