package types

type Command struct {
	Type CommandType

	// Temp
	Arg1 string
	Arg2 int

	// Arithmetic
	Operation Operation

	// Push & Pop
	Segment Segment
	Index   int
}

type CommandType int
const (
	ARITHMETIC CommandType = iota
	PUSH
	POP
	LABEL
	FUNCTION
	GOTO
	IF
	RETURN
	CALL
)


type Segment string
const (
	Local    Segment = "local"
	Argument Segment = "argument"
	Static   Segment = "static"
	Constant Segment = "constant"
	This     Segment = "this"
	That     Segment = "that"
	Pointer  Segment = "pointer"
	Temp     Segment = "temp"
)

var SegmentMap = map[Segment]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
}

type Operation string
const (
	NEG Operation = "neg"
	ADD Operation = "add"
	SUB Operation = "sub"

	NOT Operation = "not"
	AND Operation = "and"
	OR  Operation = "or"

	EQ Operation = "eq"
	LT Operation = "lt"
	GT Operation = "gt"
)
