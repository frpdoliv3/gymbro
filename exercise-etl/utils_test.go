package main

import (
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			fn:       func(x int) int { return x * 2 },
			expected: []int{},
		},
		{
			name:     "multiply by 2",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(x int) int { return x * 2 },
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "identity",
			input:    []int{10, 20, 30},
			fn:       func(x int) int { return x },
			expected: []int{10, 20, 30},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3},
			fn:       func(x int) int { return x * -1 },
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.input, tt.fn)
			if len(got) != len(tt.expected) {
				t.Errorf("Map() length = %d, expected %d", len(got), len(tt.expected))
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Map()[%d] = %d, expected %d", i, got[i], tt.expected[i])
				}
			}
		})
	}

	// Test with string transformation
	t.Run("string transformation", func(t *testing.T) {
		input := []string{"a", "bb", "ccc"}
		expected := []int{1, 2, 3}
		got := Map(input, func(s string) int { return len(s) })
		if len(got) != len(expected) {
			t.Errorf("Map() length = %d, expected %d", len(got), len(expected))
			return
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("Map()[%d] = %d, expected %d", i, got[i], expected[i])
			}
		}
	})
}

func TestNormalizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "whitespace only",
			input:    "   \t\n  ",
			expected: "",
		},
		{
			name:     "lowercase",
			input:    "hello world",
			expected: "Hello World",
		},
		{
			name:     "uppercase",
			input:    "HELLO WORLD",
			expected: "Hello World",
		},
		{
			name:     "mixed case",
			input:    "hElLo WoRlD",
			expected: "Hello World",
		},
		{
			name:     "with leading/trailing spaces",
			input:    "  hello world  ",
			expected: "Hello World",
		},
		{
			name:     "single word",
			input:    "benchpress",
			expected: "Benchpress",
		},
		{
			name:     "multiple spaces between words",
			input:    "hello   world",
			expected: "Hello   World",
		},
		{
			name:     "with tabs and newlines",
			input:    "\thello\nworld\t",
			expected: "Hello\nWorld",
		},
		{
			name:     "unicode characters",
			input:    "café au lait",
			expected: "Café Au Lait",
		},
		{
			name:     "already title case",
			input:    "Hello World",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeString(tt.input)
			if got != tt.expected {
				t.Errorf("NormalizeString(%q) = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNormalizeStringPtr(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		got := NormalizeStringPtr(nil)
		if got != nil {
			t.Errorf("NormalizeStringPtr(nil) = %v, expected nil", got)
		}
	})

	tests := []struct {
		name     string
		input    *string
		expected *string
	}{
		{
			name:     "empty string",
			input:    stringPtr(""),
			expected: stringPtr(""),
		},
		{
			name:     "whitespace only",
			input:    stringPtr("   \t\n  "),
			expected: stringPtr(""),
		},
		{
			name:     "lowercase",
			input:    stringPtr("hello world"),
			expected: stringPtr("Hello World"),
		},
		{
			name:     "uppercase",
			input:    stringPtr("HELLO WORLD"),
			expected: stringPtr("Hello World"),
		},
		{
			name:     "with spaces",
			input:    stringPtr("  bench press  "),
			expected: stringPtr("Bench Press"),
		},
		{
			name:     "mixed case with spaces",
			input:    stringPtr(" bEnCh PrEsS "),
			expected: stringPtr("Bench Press"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeStringPtr(tt.input)
			if tt.expected == nil {
				if got != nil {
					t.Errorf("NormalizeStringPtr(%v) = %v, expected nil", tt.input, got)
				}
				return
			}
			if got == nil {
				t.Errorf("NormalizeStringPtr(%v) = nil, expected %v", tt.input, tt.expected)
				return
			}
			if *got != *tt.expected {
				t.Errorf("NormalizeStringPtr(%q) = %q, expected %q", *tt.input, *got, *tt.expected)
			}
		})
	}
}
