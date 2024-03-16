package style

import (
	"regexp"

	"github.com/fatih/color"
)

type Styler struct {
	Regexp regexp.Regexp
	Color  color.Color
}

func (styler *Styler) Line(line string) string {
	if styler.Regexp.MatchString(line) {
		line = styler.Regexp.ReplaceAllStringFunc(line, func(str string) string {
			result := styler.Color.Sprint(str)
			return result
		})
	}
  return line
}

func StyleLine(line string, stylers []Styler) string {
	for _, styler := range stylers {
		line = styler.Line(line)
	}
	return line
}
