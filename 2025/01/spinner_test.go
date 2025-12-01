package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSpinner_Spin(t *testing.T) {
	t.Run("increment overflow wraps around", func(t *testing.T) {
		spinner := NewSpinner(50, 100)
		spinner.Spin(100)
		assert.Equal(t, 50, spinner.Position())

		spinner.Spin(-100)
		assert.Equal(t, 50, spinner.Position())
	})
}
