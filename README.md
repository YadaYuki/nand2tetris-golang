# Nand2Tetris :smiley_cat:

Implmentation of Nand2Tetris by Golang ʕ◔ϖ◔ʔ

# How To Work

## CPU & Logic Circuit
01 ~ 05 folders are implementation of Logic Circuit and CPU. 
You can emulate them by using [Hardware Simulator](https://www.nand2tetris.org/software).
![image](https://user-images.githubusercontent.com/57289763/128625790-8b70d3a0-bc7d-46cd-94f7-131c240742b3.png)

## Assembly

[Assembler](https://github.com/YadaYuki/nand2tetris/tree/main/assembler) convert assembly( \*.asm ) to  machine leannguage ( \*.hack)

For example, create a assembly file Add/add.asm
```
@2
D=A
@3
D=D+A
@0
D=M
```
Then, run command follow
```
go run main.go -asm add/Add.asm -hack add/Add.hack
```
After that, machine language file Add/add.hack will be generated.
```
0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1111110000010000
```



## VMTranslator

## JackCompiler

#  Reference

- [「Nand2Tetris Official Site」](https://www.nand2tetris.org/)

 - [「コンピュータシステムの理論と実装」](https://www.amazon.co.jp/%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E7%90%86%E8%AB%96%E3%81%A8%E5%AE%9F%E8%A3%85-%E2%80%95%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9-Noam-Nisan/dp/4873117127)

 - [「The Elements of Computing Systems: Building a Modern Computer from First Principles」](https://www.amazon.co.jp/Elements-Computing-Systems-Building-Principles/dp/0262640686)



 - [「Go言語でつくるインタプリタ」](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220)

 - [「Writing An Interpreter In Go 」](https://www.amazon.co.jp/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118/ref=pd_bxgy_img_1/358-0651022-5160614?pd_rd_w=NJ0lb&pf_rd_p=d8f6e0ab-48ef-4eca-99d5-60d97e927468&pf_rd_r=H5DDRH744DZQWEC8887N&pd_rd_r=92fb3969-78f9-42fe-9c0b-f605fd3b7bc8&pd_rd_wg=B98nq&pd_rd_i=3982016118&psc=1)


