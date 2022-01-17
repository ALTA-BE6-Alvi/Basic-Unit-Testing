package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPrintLuasPermukaan(t *testing.T) {
	t.Run("Radius 4 Tinggi 20", func (t *testing.T) {
		result := PrintLuasPermukaan(4, 20)
		assert.Equal(t, float64(602.88), result)
	})
	t.Run("Radius 7 Tinggi 90", func (t *testing.T) {
		result := PrintLuasPermukaan(7, 90)
		assert.Equal(t, float64(4264.12), result)
	})
	t.Run("Radius 3 Tinggi 7", func (t *testing.T) {
		result := PrintLuasPermukaan(3, 7)
		assert.Equal(t, float64(188.4), result)
	})
	t.Run("Radius 10 Tinggi 9", func (t *testing.T) {
		result := PrintLuasPermukaan(10, 9)
		assert.Equal(t, float64(1193.2), result)
	})
	t.Run("Radius 4 Tinggi 8", func (t *testing.T) {
		result := PrintLuasPermukaan(4, 8)
		assert.Equal(t, float64(301.44), result)
	})
}
