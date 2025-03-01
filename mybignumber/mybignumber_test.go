package mybignumber

import (
	"testing"
)

func TestMyBigNumber_Sum(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"1234", "897", "2131"},
		{"1229999", "1", "1230000"},
		{"0", "0", "0"},
		{"0", "1", "1"},
	}

	m := &MyBigNumber{}
	for _, tt := range tests {
		t.Run(tt.str1+"+"+tt.str2, func(t *testing.T) {
			result := m.Sum(tt.str1, tt.str2)
			if result != tt.expected {
				t.Errorf("Sum(%q, %q) = %q, want %q", tt.str1, tt.str2, result, tt.expected)
			}
		})
	}
}
