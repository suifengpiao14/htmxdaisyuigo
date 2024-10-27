package datainput

import (
	"github.com/julvo/htmlgo"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/css"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

// <input type="text" placeholder="Type here" class="input w-full max-w-xs" />
type Select struct {
	Title string
	elements.BasicComponent
	Select         elements.Select
	Label          elements.Label
	Formcontrol    elements.ComponentI
	Icon           elements.ComponentI
	Wrap           elements.ComponentI
	Border         css.Border
	Background     css.Background
	ResponseStatus css.ResponseStatus
	Color          css.Color
	Size           css.Size
}

func (i *Select) Tag() string {
	return i.Wrap.Tag()
}

func (i *Select) Component() elements.ComponentI {
	return i
}

func (i *Select) IsNil() bool {
	return false
}

func (i Select) BasicClassName() string {
	return "select"
}

func (i Select) WithBorder() Select {
	i.Border = css.Border_bordered
	return i
}
func (i Select) WithBackground(background css.Background) Select {
	i.Background = background
	return i
}

func (i Select) WithResponseStatus(status css.ResponseStatus) Select {
	i.ResponseStatus = status
	return i
}

func (i Select) WithColor(color css.Color) Select {
	i.Color = color
	return i
}

func (i Select) WithSize(size css.Size) Select {
	i.Size = size
	return i
}

func (i Select) GetWrap() elements.ComponentI {
	if i.Wrap != nil {
		return i.Wrap
	}
	return &elements.Div{}
}

func (i *Select) Class() attributes.Class {
	i.Select.AddClass(i.BasicClassName())
	if i.Color != "" {
		i.Select.AddClass(i.Color.Combine(i.BasicClassName()))
	}
	if i.Size != "" {
		i.Select.AddClass(i.Size.Combine(i.BasicClassName()))
	}
	if i.Background != "" {
		i.Select.AddClass(i.Background.Combine(i.BasicClassName()))
	}
	if i.ResponseStatus != "" {
		i.Select.AddClass(i.ResponseStatus.Combine(i.BasicClassName()))
	}
	if i.Border != "" {
		i.Select.AddClass(i.Border.Combine(i.BasicClassName()))
	}
	return i.BasicComponent.Class()
}

func (i Select) Html() (html htmlgo.HTML) {
	i.Label.SetText(i.Title)
	i.Label.For = i.Select.Name
	wrap := i.GetWrap()
	wrap.AddClass(i.Class()...)
	wrap.AddChildren(&i.Label, &i.Select)
	return wrap.Html()
}
