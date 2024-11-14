package rainbow

import (
	"github.com/fatih/color"
	"regexp"
	"strconv"
)

var (
	pattern = "^#([A-Fa-f0-9]{6})$"
	rege    *regexp.Regexp
)

type TextList struct {
	Texts []Text
}

type Text struct {
	Param string
	Color string
}

func init() {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	rege = r
}

func New(value ...Text) TextList {
	textList := TextList{}
	textList.Add(value...)
	return textList
}

func (textList *TextList) Add(value ...Text) {
	for _, text := range value {
		color := text.Color
		if !rege.MatchString(color) {
			panic("color does not match pattern")
		}
	}
	textList.Texts = append(textList.Texts, value...)
}

func (textList *TextList) Print() {
	for _, text := range textList.Texts {
		p := text.Param
		c := text.Color

		r, err := strconv.ParseUint(c[1:3], 16, 64)
		if err != nil {
			panic(err)
		}
		g, err := strconv.ParseUint(c[3:5], 16, 64)
		if err != nil {
			panic(err)
		}
		b, err := strconv.ParseUint(c[5:7], 16, 64)
		if err != nil {
			panic(err)
		}

		color.RGB(int(r), int(g), int(b)).Printf(p)
	}
}
