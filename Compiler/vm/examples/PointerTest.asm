// PUSH constant 3030
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP pointer 0
@SP
AM=M-1
D=M
@THIS
M=D
// PUSH constant 3040
@3040
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP pointer 1
@SP
AM=M-1
D=M
@THAT
M=D
// PUSH constant 32
@32
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP this 2
@THIS
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
// PUSH constant 46
@46
D=A
@SP
A=M
M=D
@SP
M=M+1
// POP that 6
@THAT
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
// PUSH pointer 0
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
// PUSH pointer 1
@THAT
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
// PUSH this 2
@THIS
D=M
@2
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
// PUSH that 6
@THAT
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
