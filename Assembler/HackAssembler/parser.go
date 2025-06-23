package main

import "strings"

type Instruction interface {
	Convert() (string, error)
}

type CInstruction struct {
	Dest string
	Comp string
	Jump string
}

type AInstruction struct {
	Value string
}

func Parse(lines []string) []Instruction {
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		var instruction Instruction;
		if line[0] == '(' {
			AddLabel(line, len(instructions)) //Note: not checking errors rn
			continue;
		} else if line[0] == '@' {
			instruction = parseAInstruction(line)			
		} else {
			instruction = parseCInstruction(line)
		}
		instructions = append(instructions, instruction)
	}
	
	return instructions
}

func parseAInstruction(line string) AInstruction {
	address, _ := strings.CutPrefix(line, "@")	
	return AInstruction{Value: address}
}

func parseCInstruction(line string) CInstruction {
	destMnemonic, remainingLine, foundDest := strings.Cut(line, "=")
	if foundDest == false {
		remainingLine = line
		destMnemonic = ""
	}
	compMnemonic, jumpMnemonic, _ := strings.Cut(remainingLine, ";")

	instruction := CInstruction{
		Dest: destMnemonic,
		Comp: compMnemonic,
		Jump: jumpMnemonic,
	}		
	return instruction
}
