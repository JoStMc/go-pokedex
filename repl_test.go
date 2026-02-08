package main

import "testing"

func TestCleanInput(t *testing.T) {
	test := map[string]struct {
	    input string
		expected []string
	}{
		"single word": {input: "hello", 
						expected: []string{"hello"}},
		"several words": { input: "hello, world", 
							expected: []string{"hello,", "world"}},
		"trailing space": { input: " hello world  ", 
							expected: []string{"hello", "world"}},
		"double space": { input: "hello  world",
							expected: []string{"hello", "world"}},
		"upper case": { input: "HELLO",
							expected: []string{"hello"}},

	} 

	for name, c := range test {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%s: expected: %v, got: %v", name, c.expected, actual)
		} 
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("%s: expected: %v, got: %v", name, c.expected, actual)
			} 
		} 
	} 
}
