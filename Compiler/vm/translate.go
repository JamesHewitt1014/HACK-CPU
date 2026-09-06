package vm

import (
	"vm/assembly"
	"vm/parser"
)

func translate(input string, filename string) string {
	parser := parser.New(input)
	writer := assembly.NewWriter()
	writer.SetFileName(filename)

	for !parser.IsComplete() {
		cmd, _ := parser.NextCommand()
		writer.WriteCommand(cmd)
	}

	return writer.Output()
}
