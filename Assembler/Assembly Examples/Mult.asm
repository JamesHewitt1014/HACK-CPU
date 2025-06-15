// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

//Set R2 to be zero
@R2
D=0
M=D

//Check if R1 is zero
@R1
D=M
@END
D;JEQ

(MULT)
    // Add R0 to R2
    @R0
    D=M
    @R2
    D=D+M
    M=D

    // Takes 1 off R1
    @R1
    D=M-1
    M=D

    // Repeat if R1 not yet zero
    @MULT
    D;JNE

(END)
@END
0;JMP

