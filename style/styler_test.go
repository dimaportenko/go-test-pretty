package style

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/fatih/color"
)

func TestStyler(t *testing.T) {
	t.Run("Red FAIL style appliend", func(t *testing.T) {
		line := "--- FAIL: TestPrettyLines (0.00s)"
		failRegex := regexp.MustCompile(`(?i)FAIL`)
		stylerColor := color.New(color.FgRed)
		styler := Styler{Regexp: *failRegex, Color: *stylerColor}

		expected := "--- \x1b[0m\x1b[31mFAIL\x1b[0m\x1b[31m: TestPrettyLines (0.00s)\x1b[0m"
		// result := fmt.Sprintf("%q", styler.Line(line))
		result := fmt.Sprintf("%q", styler.Line(line))

		if result != expected {
			t.Errorf("\nResult: %q\nExpected: %q\nline: %s", result, expected, line)
		}
	})

	t.Run("Green PASS style appliend", func(t *testing.T) {
		line := "--- PASS: TestPrettyLines (0.00s)"
		passRegex := regexp.MustCompile(`(?i)PASS`)
		var result string
		if passRegex.MatchString(line) {
			result = passRegex.ReplaceAllStringFunc(line, func(s string) string {
        green := color.New(color.FgGreen)
				t.Errorf("\nResult: %q\nExpected: %q", s, green.Sprintf(s))
				return green.Sprintf(s)
			})
		}

		expected := "--- PASS: TestPrettyLines (0.00s) 1"
		if result != expected {
			t.Errorf("\nResult: %q\nExpected: %q", line, expected)
		}
	})

	t.Run("Test pretty lines", func(t *testing.T) {
		t.Errorf("Test is not implemented")
	})
}
