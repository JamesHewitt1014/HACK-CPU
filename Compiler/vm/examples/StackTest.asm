// PUSH constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP eq
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.0
D;JEQ
@SP
A=M-1
M=0
@END.0
0;JMP
(TRUE.0)
@SP
A=M-1
M=-1
(END.0)
// PUSH constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP eq
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.1
D;JEQ
@SP
A=M-1
M=0
@END.1
0;JMP
(TRUE.1)
@SP
A=M-1
M=-1
(END.1)
// PUSH constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP eq
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.2
D;JEQ
@SP
A=M-1
M=0
@END.2
0;JMP
(TRUE.2)
@SP
A=M-1
M=-1
(END.2)
// PUSH constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP lt
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.3
D;JGT
@SP
A=M-1
M=0
@END.3
0;JMP
(TRUE.3)
@SP
A=M-1
M=-1
(END.3)
// PUSH constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP lt
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.4
D;JGT
@SP
A=M-1
M=0
@END.4
0;JMP
(TRUE.4)
@SP
A=M-1
M=-1
(END.4)
// PUSH constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP lt
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.5
D;JGT
@SP
A=M-1
M=0
@END.5
0;JMP
(TRUE.5)
@SP
A=M-1
M=-1
(END.5)
// PUSH constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP gt
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.6
D;JLT
@SP
A=M-1
M=0
@END.6
0;JMP
(TRUE.6)
@SP
A=M-1
M=-1
(END.6)
// PUSH constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP gt
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.7
D;JLT
@SP
A=M-1
M=0
@END.7
0;JMP
(TRUE.7)
@SP
A=M-1
M=-1
(END.7)
// PUSH constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP gt
@SP
AM=M-1
D=M
A=A-1
D=D-M
@TRUE.8
D;JLT
@SP
A=M-1
M=0
@END.8
0;JMP
(TRUE.8)
@SP
A=M-1
M=-1
(END.8)
// PUSH constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
// PUSH constant 53
@53
D=A
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
// PUSH constant 112
@112
D=A
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
// OP neg
@SP
A=M-1
M=-M
// OP and
@SP
AM=M-1
D=M
A=A-1
M=D&M
// PUSH constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
// OP or
@SP
AM=M-1
D=M
A=A-1
M=D|M
// OP not
@SP
A=M-1
M=!M
