


Instruction:
opcode      - bit 15 (technically bits 13-14 are also available for the opcode)
memory      - bits 12 (use A register or memory value)
computation - bits 7 - 11
destination - bits 3-5
    bit 6 - A Register
    bit 5 - D Register
    bit 4 - Write to memory
jump        - bits 0-2
    // 000 No  - false
    // 001 JGT - NOT zr AND NOT ng
    // 010 JEQ - zr 
    // 011 JGE - zr or NOT ng
    // 100 JLT - ng
    // 101 JNE - ng or NOT ng
    // 110 JLE - ng or zr
    // 111 JMP - true
