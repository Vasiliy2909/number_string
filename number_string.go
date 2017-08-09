// number_string.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	file_path := flag.String("path", "test.txt", "the file path")
	flag.Parse()

	file, _ := os.Open(*file_path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	number_string := 0
	for scanner.Scan() {
		number_string++
	}
	fmt.Println(number_string)
}
