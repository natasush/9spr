package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест для функции generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int // ожидаемый размер возвращаемого слайса
	}{
		{
			name:     "Normal size",
			size:     10,
			expected: 10,
		},
		{
			name:     "Size equal 0",
			size:     0,
			expected: 0,
		},
		{
			name:     "Negative size",
			size:     -5,
			expected: 0,
		},
		{
			name:     "Big sise",
			size:     1000,
			expected: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			// Проверяем размер возвращаемого слайса
			assert.Len(t, result, tt.expected, "generateRandomElements(%d) вернул слайс размером %d, ожидалось %d", tt.size, len(result), tt.expected)
			// Проверка на непустой срез, только если size > 0
			if tt.size > 0 {
				assert.NotEmpty(t, result, "generateRandomElements(%d) вернул пустой слайс", tt.size)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slise", []int{}, 0},
		{"Single element", []int{3}, 3},
		{"Mixed numbers", []int{-6, -5, -1, 2, 3, 5}, 5},
		{"Equal numbers", []int{4, 4, 4, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result, "maximum(%v) = %d, expected %d", tt.input, result, tt.expected)
		})

	}
}
