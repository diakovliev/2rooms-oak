package utils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/render"
)

var (
	// ErrUnableToSplit is returned when a string cannot be split
	ErrUnableToSplit = fmt.Errorf("unable to split")
)

// SplitMeasure is a measure used to split a string
type SplitMeasure interface {
	// LtOrEq checks if the measure of the given string is less than or equal to the maximum measure
	LtOrEq(str string) bool
	// Gt checks if the measure of the given string is greater than the maximum measure
	Gt(str string) bool
}

// StringLenMesure is a measure used to split a string based on its length
type StringLenMesure struct {
	MaxLen int
}

// LtOrEq checks if the length of the given string is less than or equal to the maximum length specified in the StringLenMeasure struct.
//
// Parameters:
// - str: the string to be checked.
//
// Returns:
// - bool: true if the length of the string is less than or equal to the maximum length, false otherwise.
func (s StringLenMesure) LtOrEq(str string) bool {
	return len(str) <= s.MaxLen
}

// Gt checks if the length of the given string is greater than the maximum length specified in StringLenMeasure.
//
// Parameter(s):
// - str: the string to be compared
//
// Return type:
// - bool: true if the length of the string is greater than the maximum length, false otherwise
func (s StringLenMesure) Gt(str string) bool {
	return len(str) > s.MaxLen
}

// FontWidthMesure is a measure used to split a string based on its width
type FontWidthMesure struct {
	MaxWidth float64
	Font     *render.Font
}

// mesure measures the width of a string in the given font.
//
// It takes a string parameter `str` representing the string to be measured.
// It returns a float64 representing the measured width of the string.
func (f FontWidthMesure) mesure(str string) float64 {
	return float64(f.Font.MeasureString(str).Ceil())
}

// LtOrEq determines if the font width measure of the given string is less than or equal to the maximum width.
//
// Parameters:
// - str: the string to measure the font width of.
//
// Returns:
// - bool: true if the font width measure is less than or equal to the maximum width, false otherwise.
func (f FontWidthMesure) LtOrEq(str string) bool {
	return f.mesure(str) <= f.MaxWidth
}

// Gt checks if the measured width of the given string is greater than the maximum width.
//
// str: the string to measure the width of.
// bool: true if the measured width is greater than the maximum width, false otherwise.
func (f FontWidthMesure) Gt(str string) bool {
	return f.mesure(str) > f.MaxWidth
}

// SplitToLines splits a text into lines based on a given split measure.
//
// Parameters:
// - text: the text to be split (string)
// - measure: the split measure (SplitMeasure)
//
// Returns:
// - []string: the lines of the split text
// - error: an error if unable to split the text
func SplitToLines(text string, measure SplitMeasure) ([]string, error) {
	var ret []string
	currentLexeme := ""
	currentLine := ""

	for _, r := range text {
		if unicode.IsSpace(r) && measure.LtOrEq(currentLine+currentLexeme) {
			currentLine += currentLexeme + string(r)
			currentLexeme = ""
		} else if unicode.IsSpace(r) && measure.Gt(currentLine+currentLexeme) {
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
			if measure.Gt(currentLexeme) {
				return nil, ErrUnableToSplit
			}
			currentLine = currentLexeme + string(r)
			currentLexeme = ""
		} else if unicode.IsSpace(r) {
			if measure.Gt(currentLexeme) {
				return nil, ErrUnableToSplit
			}
			currentLine += currentLexeme + string(r)
			currentLexeme = ""
		} else {
			currentLexeme += string(r)
		}
	}

	if len(currentLexeme) > 0 || len(currentLine) > 0 {
		if measure.LtOrEq(currentLine + currentLexeme) {
			currentLine += currentLexeme
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
		} else {
			ret = append(ret, strings.TrimRightFunc(currentLine, unicode.IsSpace))
			if measure.Gt(currentLexeme) {
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

	lines, err = SplitToLines(text, FontWidthMesure{MaxWidth: width, Font: font})

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
