package assembly

import (
	"fmt"
	. "vm/types"
)

func (w *asmWriter) writePop(segment Segment, index int) {
	switch segment {
	case "local", "argument", "this", "that":
		w.writePopToSegment(segment, index)
	case "temp":
		w.writePopToTemp(index)
	case "pointer":
		w.writePopToPointer(index)
	case "static":
		w.writePopToStatic(index)
	default:
		panic(fmt.Sprintf("invalid pop segment: %s", segment))
	}
}

/* Writes the assembly to pop register D off the top of the stack */
func (w *asmWriter) popToD() {
	w.add(
		"@SP",    // A    = Stack Pointer Address
		"AM=M-1", // A, M = Stack Pointer Value - 1 (Decrements @SP and loads SP-1 into Register A at the same time)
		"D=M",    // D    = RAM[SP-1] (Note that the top of the stack is empty, so SP-1 is the value at the top of the stack)
	)
}

/* Writes the assembly to store the address of label+index in Register 13 */
func (w *asmWriter) storeAddressInR13(label string, index int) {
	w.add(
		fmt.Sprintf("@%s", label),   // A = Address of Segment Pointer
		"D=M",                       // D = Segment pointer
		fmt.Sprintf("@%d", index),   // A = Index
		"D=D+A",                     // D = Pointer offset by index
		"@R13",                      // A = Address of R13
		"M=D",                       // R13 = pointer
	)
}

func (w *asmWriter) writePopToSegment(segment Segment, index int) {
	label := SegmentMap[segment]
	w.storeAddressInR13(label, index)
	w.popToD()
	w.add(
		"@R13", // A = Address of R13
		"A=M",  // A = R13 (the address of segment+index)
		"M=D",  // M = D (store popped value at target location)
	)
}

func (w *asmWriter) writePopToTemp(index int) {
	// TEMP address is always 5
	addr := 5 + index
	w.popToD()
	w.add(
		fmt.Sprintf("@%d", addr), // A = Address of TEMP
		"M=D", 					  // RAM[TEMP] = D
	)
}

func (w *asmWriter) writePopToPointer(index int) {
	addr := "THIS"
	if index == 1 {
		addr = "THAT"
	}
	w.popToD()
	w.add(
		fmt.Sprintf("@%s", addr), // A = Address
		"M=D",                    // RAM[A] = D
	)
}

func (w *asmWriter) writePopToStatic(index int) {
	w.popToD()
	w.add(
		fmt.Sprintf("@%s.%d", w.fileName, index), // A = {filename}.{index}
		"M=D",                                    // RAM[{filename}.{index}] = D
	)
}

