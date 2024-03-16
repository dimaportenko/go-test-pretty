package style

import (
	"regexp"
	"testing"

	"github.com/dimaportenko/go-test-pretty/test"
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

		test.AssertCorrectMessage(t, result, expected)
	})

	t.Run("Green PASS style appliend", func(t *testing.T) {
		line := "--- PASS: TestPrettyLines (0.00s)"
		styler := Styler{
			Regexp: *regexp.MustCompile(`(?i)PASS`),
			Color:  *color.New(color.FgGreen),
		}

		expected := "--- \x1b[32mPASS\x1b[0m: TestPrettyLines (0.00s)"
		result := styler.Line(line)

		test.AssertCorrectMessage(t, result, expected)
	})
}

