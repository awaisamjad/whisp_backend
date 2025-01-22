package internal

import (
	"testing"
)

func TestIsPasswordValid(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"1", false},
		{"Password1!", true},
		{"password1!", false},
		{"PASSWORD1!", false},
		{"Password!", false},
		{"Password1", false},
		{"Pass1!", false},
		{"P@ssw0rd", true},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			result := IsPasswordValid(test.password)
			if result != test.expected {
				t.Errorf("IsPasswordValid(%q) = %v; want %v", test.password, result, test.expected)
			}
		})
	}
}
