# HACK Computer
This is an implementation of the [HACK Computer](https://en.wikipedia.org/wiki/Hack_computer) from the book **Elements of Computing Systems: Building a Modern Computer from First Principles** and the associated course **Nand2Tetris**.

### Basic Architecture Overview
**Computer**
![image](https://github.com/user-attachments/assets/598c3094-a34f-40b6-9207-681e80468f99)
Source: Wikipedia

**Summary**
* Instructions are 16-bit. There are Two types; "A" and "C"
    * A instructions - the first bit (bit 15) represents the instruction types and the remaining bits are a 15-bit address.
    * C instructions - remaining bits indicate what operations the cpu should perform. See table below.
* The HACK CPU has two registers Register A and Register D.
* The HACK Computer has basic memory-mapped I/O for keyboard and display.

Bit | Meaning
--- | ---
15  | opcode - determines instruction type (0 for A, 1 for C)
14  | n/a
13  | n/a
12  | use A register or Memory in ALU operation
11  | ALU Operation - zero the x input
10  | ALU Operation - negate the x input
9   | ALU Operation - zero the y input
8   | ALU Operation - negate the y input
7   | ALU Operation - compute addition or logical AND
6   | ALU Operation - negate the ALU output
5   | Destination bit - write ALU output to A register
4   | Destination bit - write ALU output to D register
3   | Destination bit - write ALU output to memory 
2   | Jump bit - jump if ALU output negative
1   | Jump bit - jump if ALU output zero
0   | Jump bit - jump if ALU output positive

**CPU**


**ALU**
![image](https://github.com/JamesHewitt1014/HACK-CPU/blob/main/Project%202%20-%20ALU/ALU-HACK-DIAGRAM.svg)

# Project
Note that the .hdl and .asm files in this project are exercises from the Book/Course that I completed. These files cover the implementation of multiple aspects of the HACK Cpu, including; 
* the construction of basic logic gates starting from NAND gates,
* the adder and arithmetic logic unit (ALU),
* the memory,
* the cpu and computer as a whole.

# Resources
* [HACK Hardware Description Language (HDL)](https://drive.google.com/file/d/1dPj4XNby9iuAs-47U9k3xtYy9hJ-ET0T/view)
* [HACK ISA & Machine Language](https://en.wikipedia.org/wiki/Hack_computer#Instruction_set_architecture_(ISA)_and_machine_language)
* [HACK Assembly Language](https://en.wikipedia.org/wiki/Hack_computer#Assembly_language)
