package main

import (
	"fmt"
	"jack_compiler/compilationengine"
	"jack_compiler/tokenizer"
)

func main() {
	jt := tokenizer.New(`
	class Square {

		field int x, y; 
		field int size; 
 
		constructor Square new(int Ax, int Ay, int Asize) {
			 let x = Ax;
			 let y = Ay;
			 let size = Asize;
			 do draw();
			 return x;
		}
 
		method void dispose() {
			 do Memory.deAlloc(this);
			 return;
		}
 
		method void draw() {
			 do Screen.setColor(x);
			 do Screen.drawRectangle(x, y, x, y);
			 return;
		}
 
		method void erase() {
			 do Screen.setColor(x);
			 do Screen.drawRectangle(x, y, x, y);
			 return;
		}
 
		method void incSize() {
			 if (x) {
					do erase();
					let size = size;
					do draw();
			 }
			 return;
		}
 
		method void decSize() {
			 if (size) {
					do erase();
					let size = size;
					do draw();
			 }
			 return;
		}
 
		method void moveUp() {
			 if (y) {
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
					let y = y;
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
			 }
			 return;
		}
 
		method void moveDown() {
			 if (y) {
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
					let y = y;
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
			 }
			 return;
		}
 
		method void moveLeft() {
			 if (x) {
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
					let x = x;
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
			 }
			 return;
		}
 
		method void moveRight() {
			 if (x) {
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
					let x = x;
					do Screen.setColor(x);
					do Screen.drawRectangle(x, y, x, y);
			 }
			 return;
		}
 }  
 

`)
	ce := compilationengine.New(jt)
	fmt.Println(ce.ParseProgram().Xml())
}
