package main

import "math/rand/v2"

type password struct {
	pass             string
	length           int
	specialCharset   []rune
	numberCharset    []rune
	uppercaseCharset []rune
	lowercoseCharset []rune
}

func newPass(l int) *password {
	return &password{ // return address to new password? or return new password's address
		length: l,
	}
}

func (p *password) utfCharConfig() error {
	utfUppercaseBoundary := [2]rune{'A', 'Z'} // define rune boundary
	var utfUppercase []rune                   // create empty rune array
	for r := utfUppercaseBoundary[0]; r <= utfUppercaseBoundary[1]; r++ {
		utfUppercase = append(utfUppercase, r) // append every rune in the rune boundary to utfUppercase
	}
	p.uppercaseCharset = append(p.uppercaseCharset, utfUppercase...) // probably couldve shorten this with a pointer to p.utfUppercase to save memory

	utfLowercaseBoundary := [2]rune{'a', 'z'}
	var utfLowercase []rune
	for r := utfLowercaseBoundary[0]; r <= utfLowercaseBoundary[1]; r++ {
		utfLowercase = append(utfLowercase, r)
	}
	p.uppercaseCharset = append(p.uppercaseCharset, utfLowercase...)

	utfNumbersBoundary := [2]rune{'0', '9'}
	var utfNumbers []rune
	for r := utfNumbersBoundary[0]; r <= utfNumbersBoundary[1]; r++ {
		utfNumbers = append(utfNumbers, r)
	}
	p.uppercaseCharset = append(p.uppercaseCharset, utfNumbers...)

	utfSpecial := []rune{'!', '#', '+', '/', '&', '%', ';', '?', '-', '_'} // this could later allow to unselect special characters
	p.uppercaseCharset = append(p.uppercaseCharset, utfSpecial...)

	return nil
}

func (p *password) genPass() error {
	tmppass := "" // empty pass

	// this could be modified later, not sure how you can shorten this
	baseChars := p.uppercaseCharset
	baseChars = append(baseChars, p.lowercoseCharset...)
	baseChars = append(baseChars, p.numberCharset...)
	baseChars = append(baseChars, p.specialCharset...)

	// iterate p.length times thru baseChars
	for i := 0; i < p.length; i++ {
		randIndex := rand.IntN(len(baseChars))  // retrieve random int in range of how many items baseChar has
		tmppass += string(baseChars[randIndex]) // append typecasted baseChar at random index to tmppass string
	}
	p.pass = tmppass
	return nil
}

func (p *password) newPass() error {
	p.utfCharConfig()
	p.genPass()

	return nil
}
