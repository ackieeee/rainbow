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

		r := c[1:3]
		g := c[3:5]
		b := c[5:7]

		ruint, err := strconv.ParseUint(r, 16, 64)
		if err != nil {
			panic(err)
		}
		guint, err := strconv.ParseUint(g, 16, 64)
		if err != nil {
			panic(err)
		}
		buint, err := strconv.ParseUint(b, 16, 64)
		if err != nil {
			panic(err)
		}

		color.RGB(int(ruint), int(guint), int(buint)).Printf(p)
	}
}
