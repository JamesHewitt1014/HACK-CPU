package assembly

import (
	"fmt"
	. "vm/types"
)

func (w *asmWriter) writePush(segment Segment, index int) {
	switch segment {
	case "local", "argument", "this", "that":
		w.writePushSegment(segment, index)
	case "constant":
		w.writePushConstant(index)
	case "temp":
		w.writePushTemp(index)
	case "pointer":
		w.writePushPointer(index)
	case "static":
		w.writePushStatic(index)
	default:
		panic(fmt.Sprintf("invalid push segment: %s", segment))
	}
}

/* Writes the assembly to push register D onto the top of the stack */
func (w *asmWriter) pushD() {
	w.add(
		"@SP",   // A       = Stack Pointer Address
		"A=M",   // A       = Stack Pointer Value
		"M=D",   // RAM[SP] = Register D value
		"@SP",   // A       = Stack Pointer Address
		"M=M+1", // RAM[SP] = SP++
	)
}

func (w *asmWriter) writePushSegment(segment Segment, index int) {
	label := SegmentMap[segment]
	w.add(
		fmt.Sprintf("@%s", label), // A = Address of Segment Pointer
		"D=M",                     // D = Segment pointer
		fmt.Sprintf("@%d", index), // A = Index
		"A=D+A",                   // A = Pointer offset by index)
		"D=M",                     // D = SEGMENT[index]
	)
	w.pushD()
}

func (w *asmWriter) writePushConstant(index int) {
	w.add(
		fmt.Sprintf("@%d", index), // A = Constant number
		"D=A",                     // D = Constant number
	)
	w.pushD()
}

func (w *asmWriter) writePushTemp(index int) {
	// Base address for TEMP is always 5
	addr := 5 + index
	w.add(
		fmt.Sprintf("@%d", addr), // A = Temp Address
		"D=M",                    // D = Value @ Temp Address
	)
	w.pushD()
}

func (w *asmWriter) writePushPointer(index int) {
	addr := "THIS"
	if index == 1 {
		addr = "THAT"
	}
	w.add(
		fmt.Sprintf("@%s", addr), // A = Address
		"D=M",                    // D = Value @ address
	)
	w.pushD()
}

func (w *asmWriter) writePushStatic(index int) {
	w.add(
		fmt.Sprintf("@%s.%d", w.fileName, index), // A = {filename}.{index}
		"D=M",                                    // D = RAM[{filename}.{index}]
	)
	w.pushD()
}
