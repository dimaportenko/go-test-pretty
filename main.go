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

	// Regular expressions for matching PASS and FAIL
	passRegex := regexp.MustCompile(`(?i)PASS`)
	failRegex := regexp.MustCompile(`(?i)FAIL`)
	//
	// // Colors
	// green := color.New(color.FgGreen).SprintFunc()
	// red := color.New(color.FgRed).SprintFunc()

  stylers := []style.Styler{
    {Regexp: *passRegex, Color: *color.New(color.FgGreen)},
    {Regexp: *failRegex, Color: *color.New(color.FgRed)},
  }

	// Read from stdin and process each line
	for scanner.Scan() {
		line := scanner.Text()

    line = style.StyleLine(line, stylers)
    // line = fmt.Sprintf("%q", line)
		// Check if the line contains PASS or FAIL and color accordingly
		// if passRegex.MatchString(line) {
		// 	line = passRegex.ReplaceAllStringFunc(line, func(s string) string {
		// 		return green(s)
		// 	})
		// }
		//
		// if failRegex.MatchString(line) {
		// 	line = failRegex.ReplaceAllStringFunc(line, func(s string) string {
		// 		return red(s)
		// 	})
		// }

		// Print the processed line
		println(line)
	}

	if err := scanner.Err(); err != nil {
		color.New(color.FgRed).Printf("Error reading input: %s\n", err)
	}
}
