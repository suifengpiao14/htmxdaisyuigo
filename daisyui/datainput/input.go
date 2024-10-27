package datainput

import (
	"github.com/julvo/htmlgo"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/css"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

// <input type="text" placeholder="Type here" class="input w-full max-w-xs" />
type Input struct {
	Title string
	elements.BasicComponent
	Input          elements.Input
	Label          elements.Label
	Formcontrol    elements.ComponentI
	Icon           elements.ComponentI
	Wrap           elements.ComponentI
	IconInside     bool
	LabelInside    bool
	Border         css.Border
	Background     css.Background
	ResponseStatus css.ResponseStatus
	Color          css.Color
	Size           css.Size
}

func (i *Input) Tag() string {
	return i.Wrap.Tag()
}

func (i *Input) Component() elements.ComponentI {
	return i
}

func (i *Input) IsNil() bool {
	return false
}

func (i Input) BasicClassName() string {
	return "input"
}

func (i Input) Disabled() Input {
	i.Input.Disabled = true
	return i
}

func (i Input) WithBorder() Input {
	i.Border = css.Border_bordered
	return i
}
func (i Input) WithBackground(background css.Background) Input {
	i.Background = background
	return i
}
func (i Input) WithIconInside() Input {
	i.IconInside = true
	return i
}
func (i Input) WithLabelInside() Input {
	i.LabelInside = true
	return i
}

func (i Input) WithResponseStatus(status css.ResponseStatus) Input {
	i.ResponseStatus = status
	return i
}

func (i Input) WithColor(color css.Color) Input {
	i.Color = color
	return i
}

func (i Input) WithSize(size css.Size) Input {
	i.Size = size
	return i
}

func (i Input) GetWrap() elements.ComponentI {
	if i.Wrap != nil {
		return i.Wrap
	}
	return &elements.Div{}
}

func (i *Input) Class() attributes.Class {
	i.AddClass(i.BasicClassName())
	if i.Color != "" {
		i.AddClass(i.Color.Combine(i.BasicClassName()))
	}
	if i.Size != "" {
		i.AddClass(i.Size.Combine(i.BasicClassName()))
	}
	if i.Background != "" {
		i.AddClass(i.Background.Combine(i.BasicClassName()))
	}
	if i.ResponseStatus != "" {
		i.AddClass(i.ResponseStatus.Combine(i.BasicClassName()))
	}
	if i.Border != "" {
		i.AddClass(i.Border.Combine(i.BasicClassName()))
	}
	i.AddClass("flex", "items-center", "gap-2")
	if !i.LabelInside {
		i.Input.AddClass("grow")
	}
	return i.BasicComponent.Class()
}

func (i Input) Html() (html htmlgo.HTML) {
	i.Label.SetText(i.Title)
	i.Label.For = i.Input.Name
	if i.LabelInside {
		i.Label.AddClass(i.Class()...)
		i.Label.AddChildren(&i.Input)
		return i.Label.Html()
	}
	wrap := i.GetWrap()
	wrap.AddClass(i.Class()...)
	wrap.AddChildren(&i.Label, &i.Input)

	return wrap.Html()
}
