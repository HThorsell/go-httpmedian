package httpmedian

import (
	"strconv"
	"testing"
)

func TestStringMiddle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str      string
		length   int
		expected string
	}{
		{"abcdef", 0, ""},
		{"abcdef", 1, "c"},
		{"abcdef", 2, "cd"},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.length), func(t *testing.T) {
			result := stringMiddle(test.str, test.length)
			if result != test.expected {
				t.Errorf("stringMiddle(%q, %d) = %q; want %q", test.str, test.length, result, test.expected)
			}
		})
	}
}

func TestMedianElement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		elements []any
		expected any
	}{
		{"empty nil slice", []any{}, nil},
		{"one integer", []any{1}, 1},
		{"two strings", []any{"a", "b"}, "b"},
		{"three floats", []any{1.1, 2.2, 3.3}, 2.2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := medianElement(test.elements)
			if result != test.expected {
				t.Errorf("medianElement(%v) = %v; want %v", test.elements, result, test.expected)
			}
		})
	}
}
