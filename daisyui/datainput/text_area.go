package datainput

import (
	"github.com/julvo/htmlgo"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/css"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

type Textarea struct {
	Title string
	elements.BasicComponent
	Textarea       elements.Textarea
	Label          elements.Label
	Formcontrol    elements.ComponentI
	Wrap           elements.ComponentI
	Border         css.Border
	Background     css.Background
	ResponseStatus css.ResponseStatus
	Color          css.Color
	Size           css.Size
}

func (i *Textarea) Tag() string {
	return i.Wrap.Tag()
}

func (i *Textarea) Component() elements.ComponentI {
	return i
}

func (i *Textarea) IsNil() bool {
	return false
}

func (i Textarea) BasicClassName() string {
	return "textarea"
}
func (i Textarea) GetWrap() elements.ComponentI {
	if i.Wrap != nil {
		return i.Wrap
	}
	return &elements.Div{}
}

func (i Textarea) WithBorder() Textarea {
	i.Border = css.Border_bordered
	return i
}
func (i Textarea) WithBackground(background css.Background) Textarea {
	i.Background = background
	return i
}

func (i Textarea) WithResponseStatus(status css.ResponseStatus) Textarea {
	i.ResponseStatus = status
	return i
}

func (i Textarea) WithColor(color css.Color) Textarea {
	i.Color = color
	return i
}

func (i Textarea) WithSize(size css.Size) Textarea {
	i.Size = size
	return i
}

func (i Textarea) Class() attributes.Class {
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
	return i.BasicComponent.Class()
}

func (i Textarea) Html() (html htmlgo.HTML) {
	i.Label.SetText(i.Title)
	wrap := i.GetWrap()
	i.Label.For = i.Textarea.Name
	wrap.AddChildren(&i.Label, &i.Textarea)
	wrap.AddClass(i.Class()...)

	return wrap.Html()
}
