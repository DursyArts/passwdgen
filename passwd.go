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
	return &password{
		length: l,
	}
}

func (p *password) utfCharConfig() error {
	utfUppercaseBoundary := [2]rune{'A', 'Z'}
	var utfUppercase []rune
	for r := utfUppercaseBoundary[0]; r <= utfUppercaseBoundary[1]; r++ {
		utfUppercase = append(utfUppercase, r)
	}
	p.uppercaseCharset = append(p.uppercaseCharset, utfUppercase...)

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

	utfSpecial := []rune{'!', '#', '+', '/', '&', '%', ';', '?', '-', '_'}
	p.uppercaseCharset = append(p.uppercaseCharset, utfSpecial...)

	return nil
}

func (p *password) genPass() error {
	tmppass := ""
	baseChars := p.uppercaseCharset
	baseChars = append(baseChars, p.lowercoseCharset...)
	baseChars = append(baseChars, p.numberCharset...)
	baseChars = append(baseChars, p.specialCharset...)

	for i := 0; i < p.length; i++ {
		randIndex := rand.IntN(len(baseChars))
		tmppass += string(baseChars[randIndex])
	}
	p.pass = tmppass
	return nil
}

func (p *password) newPass() error {
	p.utfCharConfig()
	p.genPass()

	return nil
}
