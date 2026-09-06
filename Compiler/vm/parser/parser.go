package parser

import (
	"strconv"
	"strings"
	. "vm/types"
)

type VMParser struct {
	lines       []string
	index       int
}

func New(input string) *VMParser {
	cleanLines := func(str string) []string {
		// Remove Comments
		// Split Lines
		return strings.Split(str, "\n")
	}

	return &VMParser{
		lines:    cleanLines(input),
	}
}

func (p *VMParser) IsComplete() bool {
	return p.index >= len(p.lines)
}

func (p *VMParser) nextLine() (string, bool) {
	isComplete := p.IsComplete()
	if !isComplete {
		p.index++
	}
	line := p.lines[p.index]
	return line, isComplete
}

func (p *VMParser) NextCommand() (Command, error) {
	line, isComplete := p.nextLine()
	if (isComplete) {
		return Command{}, nil //TODO: ADD ERROR HANDLING
	}

	tokens := strings.Fields(line)

	cmdType, found := getCommandType[tokens[0]]
	if !found {
		return Command{}, nil // TODO: ADD ERROR HANDLING
	}

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
	case RETURN:
		return Command{}, nil
	default:
	 	return Command{}, nil
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

