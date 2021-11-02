package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	DEBUG bool = false
	// FILE  string = "test.txt"
	FILE string = ""
)

func main() {
	var scanner *bufio.Scanner
	if len(FILE) > 0 {
		file, err := os.Open(FILE)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	scanner.Scan()
	countTest, _ := strconv.Atoi(scanner.Text())
	for testN := 1; testN <= countTest; testN++ {
		scanner.Scan()
		var g int64 = Atoi64(scanner.Text())
		counter := CountVariant(g)
		fmt.Printf("Case #%v: %v\n", testN, counter)
	}
}

func CountVariant(g int64) int64 {
	var counter int64 = 0
	var z int64
	var d int64 = 0
	for {
		d++
		z = int64(g - d*(d-1)/2)
		if z < 0 {
			break
		}
		if z%d == 0 && z/d > 0 {
			counter++
			if DEBUG {
				fmt.Println("K=", z/d, "D=", d)
			}
		}
	}
	return counter
}

func Atoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
