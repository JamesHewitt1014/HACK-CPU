// PUSH constant 10
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP local 0
@LCL
D=M
@0
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// PUSH constant 21
@21
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 22
@22
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP argument 2
@ARG
D=M
@2
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// POP argument 1
@ARG
D=M
@1
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// PUSH constant 36
@36
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP this 6
@THIS
D=M
@6
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// PUSH constant 42
@42
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 45
@45
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP that 5
@THAT
D=M
@5
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// POP that 2
@THAT
D=M
@2
D=D+A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// PUSH constant 510
@510
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP temp 6
@SP
AM=M-1
D=M
@11
M=D
// PUSH local 0
@LCL
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// PUSH that 5
@THAT
D=M
@5
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// OP add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// PUSH argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// OP sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
// PUSH this 6
@THIS
D=M
@6
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// PUSH this 6
@THIS
D=M
@6
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// OP add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// OP sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
// PUSH temp 6
@11
D=M
@SP
A=M
M=D
@SP
M=M+1
// OP add
@SP
AM=M-1
D=M
A=A-1
M=D+M
