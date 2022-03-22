# Nand2tetris Jack Compiler by Golang 😺

## Overview

- **Compiler implementation by Golang ʕ◔ϖ◔ʔ**.
- Jack Compiler compiles **Jack(Object oriented language like Java,C#)** to intermediate language code.
- This package corresponds to chapter 10,11 of 「Building a Modern Computer from First Principles」

## Requirements

- Go==1.16+

## Jack ( High-level Programming Language )

- Jack is object-oriented programming language provided by nand2tetris project
- Jack's syntax is similar with Java,C#.
      
Code below stores the value given by the standard input in each element of an array, calculates the average value of the array values, and outputs the result.([jack/Average/Main.jack](https://github.com/YadaYuki/nand2tetris/blob/main/jackcompiler/jack/Average/Main.jack))

```
class Main {
   function void main() {
     var Array a; 
     var int length;
     var int i, sum;

     let length = Keyboard.readInt("How many numbers? ");
     let a = Array.new(length); 
     
     let i = 0;
     while (i < length) {
        let a[i] = Keyboard.readInt("Enter a number: ");
        let sum = sum + a[i];
        let i = i + 1;
     }
     
     do Output.printString("The average is ");
     do Output.printInt(sum / length);
     return;
   }
}
```

In addition, Jack includes the following language features

- class field
- class method
- class static variable
- if - else statement

...etc

## How to work

### Compile Jack Program to intermediate language code.

You can compile jack program by running:

```
$ go run main.go {path to Jack File or Dir} 
```

Executing this command,  single jack file or jack files in the dir passed by argument will be translated to intermediate code(.vm). Intermediate code(.vm) is generated in vm/program directory

For example, to compile `jack/HelloWorld/Main.jack` program which display "Hello,world" to screen, execute below:

```
$ go run main.go jack/HelloWorld
```

or 

```
$ go run main.go jack/HelloWorld/Main.jack
```

Executing this command, you can confirm that intermediate code file(`Main.vm`) is generated in dir `vm/program`  

In jack/ directory, there are serveral jack program. If you are interested, let's compile them by this jack compiler.

### Run intermediate code on VM Emulator

You can emulate intermediate code by VM Emulator provided by [nand2tetris official site](https://www.nand2tetris.org/software)

The gif animation below is jack/Square(`go run main.go jack/Square`) converted to intermediate code by this compiler and executed by VM Emulator.

![square_vm](https://user-images.githubusercontent.com/57289763/141052135-bcc77289-34e3-4d3a-bb7e-f3796cda87c4.gif)


The gif animation below is jack/Pong(`go run main.go jack/Pong`) converted to intermediate code by this compiler and executed by VM Emulator.

![pong_vm](https://user-images.githubusercontent.com/57289763/141052284-674d2401-afdd-4998-8614-8805685302ab.gif)

## Reference

- [「Nand2Tetris Official Site」](https://www.nand2tetris.org/)

- [「コンピュータシステムの理論と実装」](https://www.amazon.co.jp/%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E7%90%86%E8%AB%96%E3%81%A8%E5%AE%9F%E8%A3%85-%E2%80%95%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9-Noam-Nisan/dp/4873117127)

- [「The Elements of Computing Systems: Building a Modern Computer from First Principles」](https://www.amazon.co.jp/Elements-Computing-Systems-Building-Principles/dp/0262640686)

- [「Go 言語でつくるインタプリタ」](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220)

- [「Writing An Interpreter In Go 」](https://www.amazon.co.jp/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118/ref=pd_bxgy_img_1/358-0651022-5160614?pd_rd_w=NJ0lb&pf_rd_p=d8f6e0ab-48ef-4eca-99d5-60d97e927468&pf_rd_r=H5DDRH744DZQWEC8887N&pd_rd_r=92fb3969-78f9-42fe-9c0b-f605fd3b7bc8&pd_rd_wg=B98nq&pd_rd_i=3982016118&psc=1)
