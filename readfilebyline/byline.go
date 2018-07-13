package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("./readfilebyline/file.txt")
	fmt.Println(os.Getwd())
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}
