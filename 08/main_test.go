package main

import (
	"testing"
)

func Test_isVisible(t *testing.T) {
	tests := map[string]struct {
		g        grid
		i        int
		j        int
		expected bool
	}{
		"trees have equal height": {
			g: [][]height{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			i:        1,
			j:        1,
			expected: false,
		},
		"trees on the top are smaller": {
			g: [][]height{
				{1, 0, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			i:        1,
			j:        1,
			expected: true,
		},

		"trees on the right are smaller": {
			g: [][]height{
				{1, 1, 1},
				{1, 1, 0},
				{1, 1, 1},
			},
			i:        1,
			j:        1,
			expected: true,
		},

		"trees on the bottom are smaller": {
			g: [][]height{
				{1, 1, 1},
				{1, 1, 1},
				{1, 0, 1},
			},
			i:        1,
			j:        1,
			expected: true,
		},
		"trees on the left are smaller": {
			g: [][]height{
				{1, 1, 1},
				{0, 1, 1},
				{1, 1, 1},
			},
			i:        1,
			j:        1,
			expected: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := isVisible(tt.g, tt.i, tt.j)
			if actual != tt.expected {
				t.Errorf("expected expected: %t acual: %t", tt.expected, actual)
			}
		})
	}
}
