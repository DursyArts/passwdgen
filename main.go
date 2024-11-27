package main

import (
	"fmt"
)

func main() {
	pw := newPass(24)
	pw.newPass()

	fmt.Printf("password: %v\n", pw.pass)
}
