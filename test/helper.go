package test

import "testing"

func AssertCorrectMessage(t testing.TB, result, expected string) {
  t.Helper()
  if result != expected {
    t.Errorf("\nResult Q: %q\nExpected Q: %q\nResult S: %s\nExpected S: %s", result, expected, result, expected)
  }
}


