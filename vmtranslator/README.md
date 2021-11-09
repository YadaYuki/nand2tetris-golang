# Nand2tetris VMTranslator by Golang 😺

## Overview

- VMTranslator implementation by Golang ʕ◔ϖ◔ʔ.
- Assembler translates intermediate language to assembly.
- This package corresponds to chapter 7,8 of 「Building a Modern Computer from First Principles」

## Requirements

- Go==1.16

## Intermediate language

Intermediate code runs on virtual machine. 

Code below set the result of  2 + 3 to head of the stack.

```
push constant 2
push constant 3
add
```

## How to work

### Generate assembly from intermediate code

You can generate assembly from intermediate code by running:

```
$ go run main.go {path to vm dir}
```

Executing this command, assembly(.hack) program will be generated in the directory passed by argument

**※You must pass path to directory which has vm files not path to vm file.(even if the target vm file is single)**


## Reference

- [「Nand2Tetris Official Site」](https://www.nand2tetris.org/)

- [「コンピュータシステムの理論と実装」](https://www.amazon.co.jp/%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E7%90%86%E8%AB%96%E3%81%A8%E5%AE%9F%E8%A3%85-%E2%80%95%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9-Noam-Nisan/dp/4873117127)

- [「The Elements of Computing Systems: Building a Modern Computer from First Principles」](https://www.amazon.co.jp/Elements-Computing-Systems-Building-Principles/dp/0262640686)

- [「Go 言語でつくるインタプリタ」](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220)

- [「Writing An Interpreter In Go 」](https://www.amazon.co.jp/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118/ref=pd_bxgy_img_1/358-0651022-5160614?pd_rd_w=NJ0lb&pf_rd_p=d8f6e0ab-48ef-4eca-99d5-60d97e927468&pf_rd_r=H5DDRH744DZQWEC8887N&pd_rd_r=92fb3969-78f9-42fe-9c0b-f605fd3b7bc8&pd_rd_wg=B98nq&pd_rd_i=3982016118&psc=1)




