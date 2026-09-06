package main

import (
	"fmt"
	"vm/assembly"
	"vm/parser"
)

func translate(input string, filename string) string {
	parser := parser.New(input)
	writer := assembly.NewWriter()
	writer.SetFileName(filename)
	for !parser.IsComplete() {
		cmd, err := parser.NextCommand()
		if (err == nil) {
			fmt.Printf("CMD: %d", cmd.Type)
			writer.WriteCommand(cmd)
		}
	}

	return writer.Output()
}
