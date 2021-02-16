package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		s := fileScanner.Text()
		fmt.Println(s)
	}
}
