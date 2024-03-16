package main

import "testing"

func TestPrettyLines(t *testing.T) {
  t.Run("Test pretty lines", func(t *testing.T) {})

  t.Run("Test pretty lines failed", func(t *testing.T) {
    t.Errorf("This test failed")
  })
}
