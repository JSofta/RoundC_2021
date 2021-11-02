package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	MOD10 int  = 1000000007
	N_MAX int  = 100002
	DEBUG bool = false
	// FILE  string = "ts2_input.txt"
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
	amountTests, _ := strconv.Atoi(scanner.Text())
	for testN := 1; testN <= amountTests; testN++ {
		timeBegin := time.Now()
		scanner.Scan()
		valuesNandK := strings.Fields(scanner.Text())
		nValue, _ := strconv.Atoi(valuesNandK[0])
		kValue, _ := strconv.Atoi(valuesNandK[1])
		scanner.Scan()
		stringValue := scanner.Text()

		var flagsIsPalindromOrNot [N_MAX]int
		SetFlagPalindromOrNotInEveryPositions(&stringValue, &flagsIsPalindromOrNot)

		var mAmountsPalindromsKN [N_MAX]int
		SetAmountPalindromNlenghtFromKsymbols(kValue, nValue, &mAmountsPalindromsKN)

		amountPalindroms := CountPalindrom(&stringValue, nValue, 0, kValue, &flagsIsPalindromOrNot, &mAmountsPalindromsKN) % MOD10
		fmt.Printf("Case #%v: %v", testN, amountPalindroms)
		timeEnd := time.Now()
		if DEBUG {
			diff := timeEnd.Sub(timeBegin)
			diff = diff.Round(time.Millisecond)
			fmt.Printf("\tTime of execution: %s", diff)
		}
		fmt.Printf("\n")
	}
}

func SetAmountPalindromNlenghtFromKsymbols(k int, n int, amountsPalindroms *[N_MAX]int) {
	amountsPalindroms[1] = k
	amountsPalindroms[2] = k
	for i := 3; i <= n; i++ {
		amountsPalindroms[i] = (amountsPalindroms[i-2] * k) % MOD10
	}
}

func SetFlagPalindromOrNotInEveryPositions(s *string, flagsIsPalindromOrNot *[N_MAX]int) {
	strLength := len(*s)
	strCenter := strLength / 2
	flagsIsPalindromOrNot[strCenter+1] = 1
	for i := strCenter; i >= 0; i-- {
		if (*s)[i] != (*s)[strLength-1-i] {
			break
		}
		flagsIsPalindromOrNot[i] = 1
	}
}

func CountPalindrom(s *string, sLength int, leftPos int, amountSymbols int, flagsIsPalindromOrNot *[N_MAX]int, mCountPalindromsKN *[N_MAX]int) int {
	rightPos := sLength - 1 - leftPos

	// 0 symbols in string
	if leftPos > rightPos {
		return 0
	}

	counter := 0
	leftSymbol := byte((*s)[leftPos])

	// 1 OR 2 symbols in string
	if rightPos-leftPos <= 1 {
		counter = int(leftSymbol - byte('a'))
		if (*s)[leftPos] < (*s)[rightPos] {
			counter++
		}
		return counter
	}

	strLength := sLength - (leftPos+1)*2
	countPalindrom := mCountPalindromsKN[strLength]

	for c := byte('a'); c < leftSymbol; c++ {
		counter = (counter + countPalindrom)
	}
	amountPalindroms := CountPalindrom(s, sLength, leftPos+1, amountSymbols, flagsIsPalindromOrNot, mCountPalindromsKN)
	counter = counter + amountPalindroms
	if (*s)[leftPos] < (*s)[rightPos] {
		counter = counter + flagsIsPalindromOrNot[leftPos+1]
	}

	return counter
}
