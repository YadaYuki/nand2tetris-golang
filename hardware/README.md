# Nand2tetris Hardware 😺

## Overview

- Hardware implementation (CPU,ALU,Memory...etc)
- This package corresponds to chapter 1 ~ 5 of 「Building a Modern Computer from First Principles」

## Directories

What is implemented in each directory is as follows:

- **bool_gate/** ... Implementation of several bool gates (Or-gate,And-gate,Not-gate,Multiplexer ...etc) from Nand Gate. (chapter 1)

- **alu/** ... Implementation of ALU and the circuits that make it up (chapter 2)

- **sequential_circuit/** ... Implementation of several sequantial circuits (Register, Program Counter ...etc) from D Flip-Flop. (chapter 3)

- **computer/** ... Implementation of computer and the circuits(CPU,Memory ...etc) that make it up (chapter 5)

## How to work

### Emulate hardware on Hardware simulator

You can emulate these hardware by Hardware simulator provided by [nand2tetris official site](https://www.nand2tetris.org/software)

The gif animation below shows the execution of computer/ComputerRect-external.tst on the Hardware Simulator.

![hardware](https://user-images.githubusercontent.com/57289763/141056794-0eeefb5c-a7ee-458c-b8e6-1ab105b9a735.gif)



## Reference 

- [「Nand2Tetris Official Site」](https://www.nand2tetris.org/)

- [「コンピュータシステムの理論と実装」](https://www.amazon.co.jp/%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E7%90%86%E8%AB%96%E3%81%A8%E5%AE%9F%E8%A3%85-%E2%80%95%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%81%AE%E4%BD%9C%E3%82%8A%E6%96%B9-Noam-Nisan/dp/4873117127)


