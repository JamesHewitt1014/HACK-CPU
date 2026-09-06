The **Hack assembly language** is outlined in the book  **The Elements of Computing Systems: Building a Modern Computer From First Principles**. This program is my implementation of an Assembler for the Hack Assembly Language.

# 1. Registers
* `D` **register** - useful for temporary manipulation and calculations 
* `A` **register** - Stores raw data values or acts as an address index pointing to RAM or ROM
* `M` **memory** - pseudo-register, refers to the RAM address currently loaded into `A`

# 2. The Instruction Set
**A-Instructions (Addressing)**
The A-Instruction loads a 15-bit constant into the `A` register.

**Syntax** - `@value` (where `value` is a number from `0` to `32767` or a symbol)

**Example** - `@5` loads the value 5 into the `A` register

**C-Instructions (Compute & Jump)**
The C-Instruction executes computations using the ALU and manages program control flow.

**Syntax** - `destination = computation ; jump` (where both destination and jump fields are optional)

*Destinations* 
Summary: The location the output will be saved to
Options: `A`, `M`, `D`, or any combination of the three (i.e. `AMD`)

*Computations*
Summary: The operation to be performed or value to be returned
Options: `0`, `1`, `-1`, `D`, `A`, `M`, `!D`, `!A`, `!M`, `-D`, `-A`, `-M` `D+1`, `A+1`, `M+1`, `D-1`, `A-1`, `M-1`, `D+A`, `D+M`, `D-A`, `D-M`, `A-D`, `M-D`, `D&A`, `D&M`, `D|A`, `D|M`

*Jumps*
Summary: Compares the computation result against 0, to redirect execution to the address in the `A` register
Options: `JGT` (Greater Than), `JEQ` (Equal to), `JGE` (Greater or equal), `JLT` (Less than), `JNE` (if negative / less than), `JLE` (Less or equal), `JMP` (Always Jump)

# 3. Binary Representation
**A-Instruction** 
* First bit (bit-15) is `0`
* bit-0 to bit-14 is a 15-bit number

**C-Instruction**
* First bit (bit-15) is `1`
* See `C-Instruction.png`
