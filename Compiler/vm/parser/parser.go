package parser

import (
	"strconv"
	"strings"
	. "vm/types"
	"errors"
	"fmt"
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
	return p.index >= len(p.lines) - 1
}

func (p *VMParser) nextLine() (string, bool) {
	isComplete := p.IsComplete()
	if !isComplete {
		line := p.lines[p.index]
		p.index++
		return line, p.IsComplete()
	}
	return "", isComplete
}

// TODO: IDEA
// What if instead of (Command, error)
// It was (Command, bool)
// And it returns if complete
// Errors can be printed or maybe its an additional return type idk

func (p *VMParser) NextCommand() (Command, error) {
	line, isComplete := p.nextLine()
	if (isComplete) {
		return Command{}, errors.New("IsComplete") //TODO: ADD ERROR HANDLING
	}

	tokens := strings.Fields(line)

	cmdType, found := getCommandType[tokens[0]]
	if !found {
		return Command{}, errors.New(fmt.Sprintf("Not Found, %v", tokens)) // TODO: ADD ERROR HANDLING
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

