package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestSpinner_Spin(t *testing.T) {
	t.Run("increment overflow wraps around", func(t *testing.T) {
		spinner := NewSpinner(50, 100)
		spinner.Spin(100)
		assert.Equal(t, 50, spinner.Position())

		spinner.Spin(-100)
		assert.Equal(t, 50, spinner.Position())
	})

	t.Run("spinner computes overflow rotations", func(t *testing.T) {
		spinner := NewSpinner(50, 100)
		rotations := spinner.Spin(-150)
		require.Equal(t, 0, spinner.Position())
		assert.Equal(t, 2, rotations)

		rotations = spinner.Spin(-100)
		require.Equal(t, 0, spinner.Position())
		assert.Equal(t, 1, rotations)

		rotations = spinner.Spin(-1000)
		require.Equal(t, 0, spinner.Position())
		assert.Equal(t, 10, rotations)
	})

	t.Run("spinner spins backwards correctly", func(t *testing.T) {
		s := NewSpinner(0, 10)
		
		s.Spin(-1)
		require.Equal(t, 9, s.Position())

		r := s.Spin(-20)
		require.Equal(t, 9, s.Position())
		require.Equal(t, 2, r)

		r = s.Spin(-19)
		require.Equal(t, 0, s.Position())
		require.Equal(t, 2, r)
	})
}

func TestSpinner_SpinV2(t *testing.T) {
	t.Run("number of forward spins", func(t *testing.T) {
		s := NewSpinner(0, 10)

		clicks := s.SpinV2(10)
		require.Equal(t, 1, clicks)

		clicks = s.SpinV2(100)
		require.Equal(t, 10, clicks)

		clicks = s.SpinV2(17)
		require.Equal(t, 1, clicks)
	})
}
