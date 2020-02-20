package strings

import "math/rand"

const lowerCaseAlphabet string = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers string = "1234567890"

type RandomString struct {
	characters string
	Value      string
}

func (rs *RandomString) IncludeUppercaseAlphabet() {
	rs.characters = rs.characters + upperCaseAlphabet
}

func (rs *RandomString) IncludeLowercaseAlphabet() {
	rs.characters = rs.characters + lowerCaseAlphabet
}

func (rs *RandomString) IncludeNumbers() {
	rs.characters = rs.characters + numbers
}

func (rs *RandomString) Generate(length int) {

	if rs.characters == "" {
		rs.characters = lowerCaseAlphabet + upperCaseAlphabet + numbers
	}

	letterRunes := []rune(rs.characters)
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	rs.Value = string(b)
}
