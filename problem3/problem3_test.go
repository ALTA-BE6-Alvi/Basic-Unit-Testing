package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPrintNama(t *testing.T) {
	// your code here
	t.Run("Test alvin", func (t *testing.T) {
		result := PrintNama("alvin")
		assert.Equal(t, "Hello alvin! Saya Golang bahasa yang sangat menyenangkan", result)
	})
	t.Run("Test kobar", func (t *testing.T) {
		result := PrintNama("kobar")
		assert.Equal(t, "Hello kobar! Saya Golang bahasa yang sangat menyenangkan", result)
	})
	t.Run("Test sandi", func (t *testing.T) {
		result := PrintNama("sandi")
		assert.Equal(t, "Hello sandi! Saya Golang bahasa yang sangat menyenangkan", result)
	})
	t.Run("Test lalang", func (t *testing.T) {
		result := PrintNama("lalang")
		assert.Equal(t, "Hello lalang! Saya Golang bahasa yang sangat menyenangkan", result)
	})
	t.Run("Test paijo", func (t *testing.T) {
		result := PrintNama("paijo")
		assert.Equal(t, "Hello paijo! Saya Golang bahasa yang sangat menyenangkan", result)
	})
}
