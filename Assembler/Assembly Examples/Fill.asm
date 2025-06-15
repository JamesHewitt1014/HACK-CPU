// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

// Set the max address of the screen memory map
@24576
D=A
@max_screen
M=D

(CHECK)
    // Get initial pixel row of the screen map
    @SCREEN
    D=A
    @row
    M=D

    // Get keyboard input
    @KBD
    D=M
    // Jump to fill if key is pressed
    @FILL
    D;JGT

(CLEAR)
    // Set row to be clear (all zeros)
    @row
    A=M
    D=0
    M=D

    // Set pointer to next row of pixels
    @row
    D=M+1
    M=D
    // Loop CLEAR function if not reached end of screen memory map
    D=M
    @max_screen
    D=M-D
    @CLEAR
    D;JGT
    // Check once screen clearer
    @CHECK
    0;JEQ

(FILL)
    // Set row of pixels to be filled (all ones)
    @row
    A=M
    D=0
    D=!D
    M=D

    // Set pointer to next row of pixels
    @row
    D=M+1
    M=D
    // Loop FILL function if not reached end of screen memory map
    D=M
    @max_screen
    D=M-D
    @FILL
    D;JGT

@CHECK
0;JMP


