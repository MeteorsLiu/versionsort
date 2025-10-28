package versionsort

import (
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"a1", "a2", -1},
		{"a2", "a1", 1},
		{"a1", "a1", 0},
		{"a1", "a120", -1},
		{"a1", "a13", -1},
		{"a2", "a13", -1},
		{"a13", "a120", -1},
		{"1.0", "1.0.5", -1},
		{"1.0.5", "1.0", 1},
		{"8.1", "8.01", 0},
		{"8.5", "8.49", -1},
		{"8.10", "8.5", 1},
		{"foo07.7z", "foo7a.7z", 1},
		{"b1", "b3", -1},
		{"b3", "b11", -1},
		{"b11", "b20", -1},
		{"1~", "1", -1},
		{"~", "1", -1},
	}

	for _, tt := range tests {
		result := Compare(tt.a, tt.b)
		var expected int
		if tt.expected < 0 {
			expected = -1
		} else if tt.expected > 0 {
			expected = 1
		} else {
			expected = 0
		}

		var got int
		if result < 0 {
			got = -1
		} else if result > 0 {
			got = 1
		} else {
			got = 0
		}

		if got != expected {
			t.Errorf("Compare(%q, %q) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a1", "a120", "a13", "a2"},
			expected: []string{"a1", "a2", "a13", "a120"},
		},
		{
			input:    []string{"b3", "b11", "b1", "b20"},
			expected: []string{"b1", "b3", "b11", "b20"},
		},
		{
			input:    []string{"8.10", "8.5", "8.1", "8.01", "8.010", "8.100", "8.49"},
			expected: []string{"8.01", "8.1", "8.5", "8.010", "8.10", "8.49", "8.100"},
		},
		{
			input:    []string{"1", "1%", "1.2", "1~", "~"},
			expected: []string{"~", "1~", "1", "1%", "1.2"},
		},
		{
			input:    []string{"foo07.7z", "foo7a.7z"},
			expected: []string{"foo7a.7z", "foo07.7z"},
		},
	}

	for _, tt := range tests {
		items := make([]string, len(tt.input))
		copy(items, tt.input)
		Sort(items)

		for i := range items {
			if items[i] != tt.expected[i] {
				t.Errorf("Sort(%v) = %v, want %v", tt.input, items, tt.expected)
				break
			}
		}
	}
}
