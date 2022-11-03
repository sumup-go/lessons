package main

import "testing"

func TestDouble(t *testing.T) {
	expect := 4
	if res := double(2); res != expect {
		t.Fatalf("expected %v, got %v", 4, res)
	}
}

func TestDoubleTable(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "2",
			input:    2,
			expected: 4,
		},
		{
			name:     "3",
			input:    3,
			expected: 6,
		},
		// {
		// name:     "fails",
		// input:    3,
		// expected: 7,
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Helper(t)
			if res := double(tc.input); res != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, res)
			}
		})
	}
}

func Helper(t *testing.T) {
	t.Cleanup(func() {})
}
