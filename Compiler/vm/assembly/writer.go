package assembly

import (
	. "vm/types"
	"fmt"
	"strings"
)

type asmWriter struct {
	output       strings.Builder
	fileName     string
	labelCounter int
}

/* Constructor for AssembleyWriter */
func NewWriter() *asmWriter {
	return &asmWriter{}
}

func (w *asmWriter) SetFileName(filename string) {
	w.fileName = filename	
}

func (w *asmWriter) Output() string {
	return  w.output.String()
}

func (w *asmWriter) WriteCommand(cmd Command) {
	w.writeComment(cmd)
	switch cmd.Type {
	case PUSH:
		w.writePush(cmd.Segment, cmd.Index)
	case POP:
		w.writePop(cmd.Segment, cmd.Index)
	case ARITHMETIC: 
		w.writeArithmetic(cmd.Operation)
	default:
		panic(fmt.Sprintf("invalid command %d", cmd.Type))
	}
}
	
/* Add instruction(s) to the output */
func (w *asmWriter) add(lines ...string) {
	for _, line := range lines {
		w.output.WriteString(line)
		w.output.WriteByte('\n')
	}
}

func (w *asmWriter) nextLabel(prefix string) string {
	label := fmt.Sprintf("%s.%d", prefix, w.labelCounter)
	w.labelCounter++
	return label
}

func (w *asmWriter) writeComment(cmd Command) {
	var comment string
	switch cmd.Type {
	case PUSH:
		comment = fmt.Sprintf("// PUSH %s %d", cmd.Segment, cmd.Index)
	case POP:
		comment = fmt.Sprintf("// POP %s %d", cmd.Segment, cmd.Index)
	case ARITHMETIC:
		comment = fmt.Sprintf("// OP %s", cmd.Operation)
	default:
		comment = ""
	}
	w.add(comment)
}

