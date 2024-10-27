package datainput_test

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/css"
	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/datainput"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

func TestInput(t *testing.T) {
	input := datainput.Input{
		Title: "昵称",
		Input: elements.Input{
			Name:        "nickname",
			Placeholder: "请输入昵称",
			Value:       "suifeng",
		},
		Border: css.Border_bordered,
		Color:  css.Color_primary,
	}
	html := input.Html()
	fmt.Println(html)
}

func TestTextarea(t *testing.T) {
	input := datainput.Textarea{
		Title:  "昵称",
		Border: css.Border_bordered,
		Color:  css.Color_primary,
		Textarea: elements.Textarea{
			Name: "nickname",
		},
	}
	html := input.Html()
	fmt.Println(html)
}
