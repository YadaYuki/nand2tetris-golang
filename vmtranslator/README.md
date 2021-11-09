# Nand2tetris VMTranslator by Golang 😺

## Overview

- VMTranslator implementation by Golang ʕ◔ϖ◔ʔ.
- Assembler translates intermediate language to assembly.
- This package corresponds to chapter 7,8 of 「Building a Modern Computer from First Principles」

## Requirements

- Go==1.16

### Intermediate language

Intermediate code runs on virtual machine. 

Code below set the result of  2 + 3 to head of the stack.

```
push constant 2
push constant 3
add
```

## How to work

### Generate assembly from intermediate code




