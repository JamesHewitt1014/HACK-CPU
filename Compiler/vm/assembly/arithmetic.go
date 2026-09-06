package assembly

import (
	. "vm/types"
	"fmt"
)

var binaryOps = map[Operation]string{
	"add": "M=D+M",
	"sub": "M=M-D",
	"and": "M=D&M",
	"or":  "M=D|M",
	"neg": "M=-M",
	"not": "M=!M",
	"eq":  "JEQ",
	"gt":  "JLT",
	"lt":  "JGT",
}

func (w *asmWriter) writeArithmetic(operation Operation) {
	instruction := binaryOps[operation] // M = Operation(M, D)
	switch operation {
	// Note: Instead of popping the second value and pushing the result, we can just replace the second value with the result
	case ADD, SUB, AND, OR:
		w.popToD()
		w.add(
			"A=M-1",     // A = SP - 1 (assumes A is already the SP address following popToD, SP-1 is the value at the top of the stack)
			instruction, // RAM[SP-1] = Operation(D, RAM[SP-1])
		)
	case NOT, NEG:
		w.add(
			"@SP",       // A = SP Address
			"A=M-1",     // A = SP-1
			instruction, // RAM[SP-1] = Operation(RAM[SP-1])
		)
	case EQ, LT, GT:
		w.comparison(instruction)
	default:
		panic(fmt.Sprintf("invalid operation: %s", operation))
	}
}

func (w *asmWriter) comparison(jumpOp string) {
	labelTrue := w.nextLabel("TRUE")
	labelEnd := w.nextLabel("END")
	w.popToD() // D = y
	w.add(
		"A=M-1", // A = SP-1  (assumes A is already the SP address following popToD)
		"D=M-D", // D = x - y (where x is current value at SP-1)
		fmt.Sprintf("@%s", labelTrue),
		fmt.Sprintf("D;%s", jumpOp),
		"M=0", // RAM[SP-1] = 0 (false)
		fmt.Sprintf("(%s)", labelTrue),
		"M=-1", // RAM[SP-1] = -1 (true)
		fmt.Sprintf("(%s)", labelEnd),
	)
}

