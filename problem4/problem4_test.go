package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPrintLuas(t *testing.T) {
	// your code here
	t.Run("Alas 5 Tinggi 10", func (t *testing.T) {
		result := PrintLuas(5, 10)
		assert.Equal(t, float64(25), result)
	})
	t.Run("Alas 7 Tinggi 90", func (t *testing.T) {
		result := PrintLuas(7, 90)
		assert.Equal(t, float64(315), result)
	})
	t.Run("Alas 2 Tinggi 8", func (t *testing.T) {
		result := PrintLuas(2, 8)
		assert.Equal(t, float64(8), result)
	})
	t.Run("Alas 10 Tinggi 9", func (t *testing.T) {
		result := PrintLuas(10, 9)
		assert.Equal(t, float64(45), result)
	})
	t.Run("Alas -5 Tinggi 10", func (t *testing.T) {
		result := PrintLuas(4, 8)
		assert.Equal(t, float64(16), result)
	})
}
