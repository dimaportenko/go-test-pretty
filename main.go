package main

import (
	"bufio"
	"os"
	"regexp"

	"github.com/dimaportenko/go-test-pretty/style"
	"github.com/fatih/color"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

  stylers := getTestOutputStylers()

	for scanner.Scan() {
		line := scanner.Text()

    line = style.StyleLine(line, stylers)
		println(line)
	}

	if err := scanner.Err(); err != nil {
		color.New(color.FgRed).Printf("Error reading input: %s\n", err)
	}
}

func getTestOutputStylers() []style.Styler {
	passRegex := regexp.MustCompile(`\bPASS\b`)
	failRegex := regexp.MustCompile(`\bFAIL\b`)

  stylers := []style.Styler{
    {Regexp: *passRegex, Color: *color.New(color.FgGreen)},
    {Regexp: *failRegex, Color: *color.New(color.FgRed)},
  }

  return stylers
}
