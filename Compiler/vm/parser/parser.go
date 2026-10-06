package parser

import (
	. "vm/types"
	"strconv"
	"strings"
	"fmt"
	"io"
)

type VMParser struct {
	lines       []string
	index       int
}

func New(input string) *VMParser {
	var lines []string
	// Clean lines; remove comments, trim whitespace on ends, remove empty lines
	for _, line := range strings.Split(input, "\n") {
		if i := strings.Index(line, "//"); i >= 0 {
       		line = line[:i]
		}
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}

	return &VMParser{ lines: lines }
}

func (p *VMParser) IsComplete() bool {
	return p.index >= len(p.lines)
}

func (p *VMParser) NextCommand() (Command, error) {
	// Next Line
	if p.IsComplete() {
		return Command{}, io.EOF
	}
	line := p.lines[p.index]
	p.index++

	// Convert line into tokens
	tokens := strings.Fields(line)
	cmdType, found := getCommandType[tokens[0]]
	if !found {
		return Command{}, fmt.Errorf("Not Found, %v", tokens)
	}

	expectedArgCount := numberOfArguments[cmdType];
	if len(tokens) < expectedArgCount {
		return Command{}, fmt.Errorf("Too few arguments: %s", line)
	}


	// Return Command
	switch cmdType {
	case ARITHMETIC:
		op := Operation(tokens[0])
		return Command{ Type: cmdType, Operation: op }, nil
	case PUSH, POP:
		segment := Segment(tokens[1])
		index, _ := strconv.Atoi(tokens[2])
	 	return Command{ Type: cmdType, Segment: segment, Index: index}, nil
	case FUNCTION, CALL:
		index, _ := strconv.Atoi(tokens[2])
		return Command{ Type: cmdType, Arg1: tokens[1], Arg2: index }, nil
	default: // RETURN
	return Command{Type: cmdType}, nil
	}
}

var getCommandType = map[string]CommandType{
	"push":   PUSH,
	"pop":    POP,
	"label":  LABEL,
	"goto":   GOTO,
	"return": RETURN,
	"call":   CALL,
	"add":    ARITHMETIC,
	"sub":    ARITHMETIC,
	"neg":    ARITHMETIC,
	"eq":     ARITHMETIC,
	"gt":     ARITHMETIC,
	"lt":     ARITHMETIC,
	"and":    ARITHMETIC,
	"or":     ARITHMETIC,
	"not":    ARITHMETIC,
}

// The number of arguments a certain command type takes
var numberOfArguments = map[CommandType]int {
	ARITHMETIC: 1,
	PUSH: 3,
	POP: 3,
	FUNCTION: 3,
	CALL: 3,
	RETURN: 1,
	LABEL: 2,
	GOTO: 2,
	IF: 2,
}

