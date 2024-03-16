package style

import (
	"regexp"
	"testing"

	"github.com/fatih/color"
)

func TestStyler(t *testing.T) {
	color.NoColor = false

	t.Run("Red FAIL style appliend", func(t *testing.T) {
		line := "--- FAIL: TestPrettyLines (0.00s)"
		styler := Styler{
			Regexp: *regexp.MustCompile(`(?i)FAIL`),
			Color:  *color.New(color.FgRed),
		}

		expected := "--- \x1b[31mFAIL\x1b[0m: TestPrettyLines (0.00s)"
		result := styler.Line(line)

    assertCorrectMessage(t, expected, result)
	})

	t.Run("Green PASS style appliend", func(t *testing.T) {
		line := "--- PASS: TestPrettyLines (0.00s)"
		styler := Styler{
			Regexp: *regexp.MustCompile(`(?i)PASS`),
			Color:  *color.New(color.FgGreen),
		}

		expected := "--- \x1b[32mPASS\x1b[0m: TestPrettyLines (0.00s)"
		result := styler.Line(line)

    assertCorrectMessage(t, expected, result)
	})
}

func assertCorrectMessage(t testing.TB, result, expected string) {
  t.Helper()
  if result != expected {
    t.Errorf("\nResult Q: %q, \nResult S: %s\nExpected Q: %q\nExpected S: %s", result, result, expected, expected)
  }
}

