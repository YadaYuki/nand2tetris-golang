# Nand2tetris Assembler by Golang 😺

## Overview

- Assembler implementation by Golang ʕ◔ϖ◔ʔ.
- Assembler translates assembly to machine language.
- This package corresponds to chapter 6 of 「Building a Modern Computer from First Principles」
       

## Requirements

- Go==1.16

## Assembly and Machine Language.

### Assembly

Assembly is a format in which machine language are represented by symbols for easy human understanding. There is one machine language instruction corresponding to one assembly command.

Code below is assembly program [Add.asm](https://github.com/YadaYuki/nand2tetris/blob/main/assembler/asm/add/Add.asm) which stores the result of 2 + 3 to 0th register(R0) 

```
@2
D=A
@3
D=D+A
@0
M=D
```

### Machine Language

Machine language is a format of instructions that can be directly understood by the CPU.

Code below is machine language program which translate [Add.asm](https://github.com/YadaYuki/nand2tetris/blob/main/assembler/asm/add/Add.asm) to machine language .

```
0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000
```

## How to work

### Generate machine language from assembly

You can generate machine language from assembly by running:

```
$ go run main.go {path to asm file}
```

Executing this command, machine language program (.hack) will be generated in same dir as assembly file(.asm)

For example, to translate `asm/add/Add.asm` program which stores the result of 2 + 3 to 0th register(R0)  to machine language, execute below:

```
$ go run main.go asm/add/Add.asm
```

Executing this command, you can confirm that machine language program file(`Add.hack`) is generated in same dir of `Add.asm`(`asm/add/`)  

### Run machine language program on CPU Emulator

You can emulate machine language program by CPU Emulator provided by [nand2tetris official site](https://www.nand2tetris.org/software)

The gif animation below is pong/Pong.asm converted to machine language by this assembler and executed by CPU Emulator.

![cpu](https://user-images.githubusercontent.com/57289763/140737209-5759bd5c-e476-471f-bfb0-8cb00b0610a1.gif)


## Reference

- [「Nand2Tetris Official Site」](https://www.nand2tetris.org/)

- [「コンピュータシステムの理論と実装」](https://www.amazon.co.jp/%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E7%90%86%E8%AB%96%E3%81%A8%E5%AE%9F%E8%A3%85-%E2%80%95%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9-Noam-Nisan/dp/4873117127)

- [「The Elements of Computing Systems: Building a Modern Computer from First Principles」](https://www.amazon.co.jp/Elements-Computing-Systems-Building-Principles/dp/0262640686)

- [「Go 言語でつくるインタプリタ」](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220)

- [「Writing An Interpreter In Go 」](https://www.amazon.co.jp/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118/ref=pd_bxgy_img_1/358-0651022-5160614?pd_rd_w=NJ0lb&pf_rd_p=d8f6e0ab-48ef-4eca-99d5-60d97e927468&pf_rd_r=H5DDRH744DZQWEC8887N&pd_rd_r=92fb3969-78f9-42fe-9c0b-f605fd3b7bc8&pd_rd_wg=B98nq&pd_rd_i=3982016118&psc=1)
