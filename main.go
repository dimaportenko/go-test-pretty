package main

import (
	"bufio"
	"os"
	"regexp"

	"github.com/fatih/color"

	"github.com/dimaportenko/go-test-pretty/style"
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

func getStylerRegexFromText(text string) regexp.Regexp {
	return *regexp.MustCompile(`\b` + text + `\b`)
}

func getTestOutputStylers() []style.Styler {
	passRegex := getStylerRegexFromText("PASS")
	failRegex := getStylerRegexFromText("FAIL")

	stylers := []style.Styler{
		{Regexp: passRegex, Color: *color.New(color.FgGreen)},
		{Regexp: failRegex, Color: *color.New(color.FgRed)},
	}

	return stylers
}
