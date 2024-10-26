package daisyui

import (
	"strings"

	"github.com/julvo/htmlgo"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

type Component interface {
	Html() string
}

// <input type="text" placeholder="Type here" class="input w-full max-w-xs" />
type Input struct {
	Title string
	Class attributes.Class
	elements.Input
	Label          elements.Label
	Formcontrol    Component
	Icon           Component
	withIconInside bool
	labelInside    bool
}

func (i Input) BasicClassName() string {
	return "input"
}

func (i Input) Disabled() Input {
	i.Input.Disabled = true
	return i
}

func (i Input) WithBorder() Input {
	i.Class.Add(Border_bordered.Combine(i.BasicClassName()))
	return i
}
func (i Input) WithBackground(background Background) Input {
	i.Class.Add(background.Combine(i.BasicClassName()))
	return i
}
func (i Input) WithIconInside() Input {
	i.withIconInside = true
	return i
}
func (i Input) WithLabelInside() Input {
	i.labelInside = true
	return i
}

func CombineClassName(segments ...string) (className string) {
	className = strings.Join(segments, "-")
	return className
}

func (i Input) WithResponseStatus(status ResponseStatus) Input {
	i.Class.Add(status.Combine(i.BasicClassName()))
	return i
}

func (i Input) WithColor(color Color) Input {
	i.Class.Add(color.Combine(i.BasicClassName()))
	return i
}

func (i Input) WithSize(size Size) Input {
	i.Class.Add(size.Combine(i.BasicClassName()))
	return i
}

func (i Input) Html() (html htmlgo.HTML) {
	i.Class.Add(i.BasicClassName())
	if i.labelInside {
		i.Label.Class = i.Class
		i.Label.Children = []htmlgo.HTML{
			htmlgo.Text_(i.Title),
			i.Input.Html(),
		}
		return i.Label.Html()
	}

	i.Input.Class = i.Class

	div := elements.Div{
		Children: []htmlgo.HTML{
			i.Label.Html(),
			i.Input.Html(),
		},
	}

	return div.Html()
}
