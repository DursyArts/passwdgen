package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	utfUppercase, utfLowercase, utfNumbers := initUtf()
	utfTable := append(utfUppercase, utfLowercase...)
	utfTable = append(utfTable, utfNumbers...)
	fmt.Printf("uppercase chars: %c\n", utfUppercase[:])
	fmt.Printf("lowercase chars: %c\n", utfLowercase[:])
	fmt.Printf("number chars: %c\n", utfNumbers[:])
	fmt.Printf("lookup table: %c\n", utfTable[:])

	reader := rand.Reader
	var passLen int64 = 24
	var minLen int64 = 0
	var password string
	fmt.Printf("Generating %v char. long password: ", passLen)
	for i := minLen; i < passLen; i++ {
		genBoundary, err := rand.Int(reader, big.NewInt(int64(len(utfTable))))
		if err != nil {
			fmt.Printf("error generating random index: %v\n", err)
		}
		fmt.Printf("%c", utfTable[int(genBoundary.Int64())])
		password += string(utfTable[int(genBoundary.Int64())])
		time.Sleep(time.Millisecond * 50)
	}
	//fmt.Printf("generated password: %v", password)
}

func initUtf() ([]rune, []rune, []rune) {
	utfUppercaseBoundary := [2]rune{'A', 'Z'}
	var utfUppercase []rune
	for r := utfUppercaseBoundary[0]; r <= utfUppercaseBoundary[1]; r++ {
		utfUppercase = append(utfUppercase, r)
	}

	utfLowercaseBoundary := [2]rune{'a', 'z'}
	var utfLowercase []rune
	for r := utfLowercaseBoundary[0]; r <= utfLowercaseBoundary[1]; r++ {
		utfLowercase = append(utfLowercase, r)
	}

	utfNumbersBoundary := [2]rune{'0', '9'}
	var utfNumbers []rune
	for r := utfNumbersBoundary[0]; r <= utfNumbersBoundary[1]; r++ {
		utfNumbers = append(utfNumbers, r)
	}

	return utfUppercase, utfLowercase, utfNumbers
}
