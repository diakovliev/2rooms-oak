package utils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/render"
)

var (
	ErrUnableToSplit = fmt.Errorf("unable to split")
)

// SplitToLines splits the given text into lines of maximum length.
//
// Parameters:
// - text: the text to be split into lines.
// - maxLen: the maximum length of each line.
//
// Returns:
// - []string: an array of strings representing the lines.
func SplitToLines(text string, maxLen int) ([]string, error) {
	var ret []string
	currentLexeme := ""
	currentLine := ""

	for _, r := range text {
		if unicode.IsSpace(r) && len(currentLine)+len(currentLexeme) <= maxLen {
			currentLine += currentLexeme + string(r)
			currentLexeme = ""
		} else if unicode.IsSpace(r) && len(currentLine)+len(currentLexeme) > maxLen {
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
			if len(currentLexeme) > maxLen {
				return nil, ErrUnableToSplit
			}
			currentLine = currentLexeme + string(r)
			currentLexeme = ""
		} else if unicode.IsSpace(r) {
			if len(currentLexeme) > maxLen {
				return nil, ErrUnableToSplit
			}
			currentLine += currentLexeme + string(r)
			currentLexeme = ""
		} else {
			currentLexeme += string(r)
		}
	}

	if len(currentLexeme) > 0 || len(currentLine) > 0 {
		if len(currentLine)+len(currentLexeme) <= maxLen {
			currentLine += currentLexeme
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
		} else {
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
			if len(currentLexeme) > maxLen {
				return nil, ErrUnableToSplit
			}
			ret = append(ret, strings.TrimRightFunc(currentLexeme, unicode.IsSpace))
		}
	}

	return ret, nil
}

// SplitToLinesWithFontMeasure splits a text into lines based on a given width and font measurement.
//
// Parameters:
// - text: the text to be split into lines.
// - width: the maximum width of each line.
// - font: the font used for measuring the width of the text.
//
// Returns:
// - []string: the lines resulting from the split.
// - error: an error that occurred during the split, if any.
func SplitToLinesWithFontMeasure(text string, width float64, font *render.Font) ([]string, error) {
	if font == nil {
		font = render.DefaultFont()
	}

	measureString := func(str string) float64 {
		return float64(font.MeasureString(str).Ceil())
	}

	var ret []string
	currentLexeme := ""
	currentLine := ""

	for _, r := range text {
		if unicode.IsSpace(r) && measureString(currentLine+currentLexeme) <= width {
			currentLine += currentLexeme + string(r)
			currentLexeme = ""
		} else if unicode.IsSpace(r) && measureString(currentLine+currentLexeme) > width {
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
			if measureString(currentLexeme) > width {
				return nil, ErrUnableToSplit
			}
			currentLine = currentLexeme + string(r)
			currentLexeme = ""
		} else if unicode.IsSpace(r) {
			if measureString(currentLexeme) > width {
				return nil, ErrUnableToSplit
			}
			currentLine += currentLexeme + string(r)
			currentLexeme = ""
		} else {
			currentLexeme += string(r)
		}
	}

	if len(currentLexeme) > 0 || len(currentLine) > 0 {
		if measureString(currentLine+currentLexeme) <= width {
			currentLine += currentLexeme
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
		} else {
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
			if measureString(currentLexeme) > width {
				return nil, ErrUnableToSplit
			}
			ret = append(ret, strings.TrimRightFunc(currentLexeme, unicode.IsSpace))
		}
	}

	return ret, nil

}

// TextFitWidth generates a list of lines that fit within the given width based on the provided text, font, and width.
//
// Parameters:
// - text: The text to fit within the width.
// - width: The width to fit the text into.
// - font: The font to use for measuring the text.
//
// Returns:
// - lines: A list of lines that fit within the given width.
// - rect: The bounding rectangle of the fitted text.
func TextFitWidth(text string, width float64, font *render.Font) (lines []string, rect floatgeom.Rect2, lh float64, err error) {
	if font == nil {
		font = render.DefaultFont()
	}

	h := font.Height()
	lh = float64(h)
	if len(text) == 0 {
		rect = floatgeom.Rect2{
			Min: floatgeom.Point2{0, 0},
			Max: floatgeom.Point2{width, h},
		}
		return
	}

	w := float64(font.MeasureString(text).Ceil())

	if w <= width {
		// whole string fits requested width
		lines = append(lines, text)
		rect = floatgeom.Rect2{
			Min: floatgeom.Point2{0, 0},
			Max: floatgeom.Point2{width, h},
		}
		return
	}

	lines, err = SplitToLinesWithFontMeasure(text, width, font)

	if err != nil {
		err = fmt.Errorf("%w: unable to fit text to width %v", err, width)
		return
	}

	rect = floatgeom.Rect2{
		Min: floatgeom.Point2{0, 0},
		Max: floatgeom.Point2{width, h * float64(len(lines))},
	}

	return
}
