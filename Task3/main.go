package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_DAYS int  = 200
	ROUNDS   int  = 60
	DEBUG    bool = false
	// FILE  string = "ts2_input.txt"
	FILE string = ""
	// FILE string = "test.txt"
)

type Point struct {
	X int
	Y int
	Z int
}

func main() {
	var amountDays int
	var averagePayoff int
	var dayWin [MAX_DAYS + 1]int
	var dayEqual [MAX_DAYS + 1]int

	Initialize(&amountDays, &averagePayoff, &dayWin, &dayEqual)
	for nDay := 1; nDay <= amountDays; nDay++ {
		Solve(nDay, dayWin[nDay], dayEqual[nDay])
	}
}

func Solve(nDay int, w int, e int) {
	var v [ROUNDS + 1][ROUNDS + 1][ROUNDS + 1]float32 // Summa
	var s [ROUNDS + 1][ROUNDS + 1][ROUNDS + 1]int     // Number of Point

	startValue := float32(w)/3 + float32(e)/3
	// fill cube border = FIRST ROUND
	for i := 1; i <= ROUNDS; i++ {
		v[i][0][0] = startValue
		v[0][i][0] = startValue
		v[0][0][i] = startValue
	}

	var valR, valP, valS float32
	for round := 2; round <= ROUNDS; round++ {
		for i := 0; i <= ROUNDS; i++ {
			for j := 0; j <= ROUNDS; j++ {
				for k := 0; k <= ROUNDS; k++ {
					if i+j+k == round {
						rPrev := round - 1
						if i > 0 {
							valR = v[i-1][j][k] + float32(w*j/rPrev) + float32(e*k/rPrev)
						} else {
							valR = float32(w*j/rPrev) + float32(e*k/rPrev)
						}
						if j > 0 {
							valP = v[i][j-1][k] + float32(w*k/rPrev) + float32(e*i/rPrev)
						} else {
							valP = float32(w*k/rPrev) + float32(e*i/rPrev)
						}
						if k > 0 {
							valS = v[i][j][k-1] + float32(w*i/rPrev) + float32(e*j/rPrev)
						} else {
							valS = float32(w*i/rPrev) + float32(e*j/rPrev)
						}
						valMax, pos := Max3(valR, valP, valS)
						v[i][j][k] = valMax
						s[i][j][k] = pos
					}
					if (i + j + k) >= round {
						break
					}
				}
				if (i + j) >= round {
					break
				}
			}
		}
	}
	maxPoint := findMaxPoint(&v)
	path := findPath(maxPoint, &s)

	path = "R" + path
	fmt.Printf("Case #%d: ", nDay)
	fmt.Println(path)

	return
}

func findPath(maxPoint Point, s *[ROUNDS + 1][ROUNDS + 1][ROUNDS + 1]int) string {
	// 1 - R
	// 2 - P
	// 3 - S
	path := ""
	thisX := maxPoint.X
	thisY := maxPoint.Y
	thisZ := maxPoint.Z

	for i := ROUNDS; i > 1; i-- {
		//		fmt.Println("path step", i, thisX, thisY, thisZ, s[thisX][thisY][thisZ])
		if s[thisX][thisY][thisZ] == 1 {
			path = "R" + path
			thisX--
		} else if s[thisX][thisY][thisZ] == 2 {
			path = "P" + path
			thisY--
		} else if s[thisX][thisY][thisZ] == 3 {
			path = "S" + path
			thisZ--
		}
	}
	return path
}

func findMaxPoint(v *[ROUNDS + 1][ROUNDS + 1][ROUNDS + 1]float32) Point {
	maxPoint := Point{}
	maxValue := v[1][0][0]
	for i := 1; i <= ROUNDS; i++ {
		for j := 1; j <= ROUNDS; j++ {
			for k := 1; k <= ROUNDS; k++ {
				if i+j+k == ROUNDS {
					if v[i][j][k] > maxValue {
						maxValue = v[i][j][k]
						maxPoint.X = i
						maxPoint.Y = j
						maxPoint.Z = k
					}
				}
				if (i + j + k) >= ROUNDS {
					break
				}
			}
			if (i + j) >= ROUNDS {
				break
			}
		}
	}
	//	fmt.Println("maxPoint", maxPoint)
	return maxPoint
}

func Max3(v1 float32, v2 float32, v3 float32) (float32, int) {
	maxValue := v1
	nPoint := 1
	if v2 > maxValue {
		maxValue = v2
		nPoint = 2
	}
	if v3 > maxValue {
		maxValue = v3
		nPoint = 3
	}
	return maxValue, nPoint
}

func Initialize(amountDays *int, averagePayoff *int, dayWin *[MAX_DAYS + 1]int, dayEqual *[MAX_DAYS + 1]int) {
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
	*amountDays, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	*averagePayoff, _ = strconv.Atoi(scanner.Text())

	for i := 1; i <= *amountDays; i++ {
		scanner.Scan()
		valuesString := strings.Fields(scanner.Text())
		dayWin[i], _ = strconv.Atoi(valuesString[0])
		dayEqual[i], _ = strconv.Atoi(valuesString[1])
	}
}
