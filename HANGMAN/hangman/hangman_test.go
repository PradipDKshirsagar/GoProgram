package hangman

import "testing"

func TestTablecheckInputLetter(t *testing.T) {
	var tests = []struct {
		intputLetter []string
		letter string
		expected bool
	}{
		{[]string{"A", "P", "S", "T",}, "T", true},
		{[]string{"Q", "E", "X", "P",},"P" , true},
		{[]string{"L", "P", "W", "Q", "M",}, "N", false},
		{[]string{}, "P", false},
	}
	for _, test :=range tests {
		if output := checkInputLetter(test.intputLetter, test.letter); output != test.expected {
			t.Error("Test Failed : {} InputLetter,{} Letter,{} expected, received: {}", test.intputLetter, test.letter, test.expected, output)
		} 
	}
}
