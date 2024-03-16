package main

import (
	"testing"

	"github.com/fatih/color"

	"github.com/dimaportenko/go-test-pretty/style"
	"github.com/dimaportenko/go-test-pretty/test"
)

func TestOutputStylers(t *testing.T) {
	color.NoColor = false

	t.Run("Test test lower case", func(t *testing.T) {
		line := "--- FAIL: TestOutputStylers/Test_pretty_lines_failed"
		stylers := getTestOutputStylers()

		expected := "--- \x1b[31mFAIL\x1b[0m: TestOutputStylers/Test_pretty_lines_failed"
		result := style.StyleLine(line, stylers)

		test.AssertCorrectMessage(t, result, expected)
	})

	t.Run("Test surounded with underscores", func(t *testing.T) {
		line := "--- PASS: TestStyler/Red_FAIL_PASS_style_appliend (0.00s)"
		stylers := getTestOutputStylers()

		expected := "--- \x1b[32mPASS\x1b[0m: TestStyler/Red_FAIL_PASS_style_appliend (0.00s)"
		result := style.StyleLine(line, stylers)

		test.AssertCorrectMessage(t, result, expected)
	})

	t.Run("Test pretty lines failed", func(t *testing.T) {
		t.Errorf("This test failed")
	})
}
