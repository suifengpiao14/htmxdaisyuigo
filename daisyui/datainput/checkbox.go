package datainput

import (
	"github.com/julvo/htmlgo"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/css"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

type Option struct {
	Label    string
	Value    string
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Checked  bool   `json:"checked"`
}
type Options []Option

type Checkbox struct {
	Title string
	elements.BasicComponent
	Checkbox       elements.Checkbox
	Label          elements.Label
	Formcontrol    elements.ComponentI
	ResponseStatus css.ResponseStatus
	Color          css.Color
	Size           css.Size
	Options        Options
}

func (i *Checkbox) Tag() string {
	if i.Formcontrol == nil {
		return ""
	}
	return i.Formcontrol.Tag()
}

func (i *Checkbox) GetFormcontrol() (formcontrol elements.ComponentI) {
	if i.Formcontrol == nil {
		i.Formcontrol = &elements.Div{}
		i.Formcontrol.AddClass(css.Formcontrol)
	}
	return i.Formcontrol
}

func (i *Checkbox) Component() elements.ComponentI {
	return i
}

func (i *Checkbox) IsNil() bool {
	return false
}

func (i Checkbox) BasicClassName() string {
	return "checkbox"
}

func (i Checkbox) WithResponseStatus(status css.ResponseStatus) Checkbox {
	i.ResponseStatus = status
	return i
}

func (i Checkbox) WithColor(color css.Color) Checkbox {
	i.Color = color
	return i
}

func (i Checkbox) WithSize(size css.Size) Checkbox {
	i.Size = size
	return i
}

func (i Checkbox) Class() attributes.Class {
	i.Checkbox.AddClass(i.BasicClassName())
	if i.Color != "" {
		i.Checkbox.AddClass(i.Color.Combine(i.BasicClassName()))
	}
	if i.Size != "" {
		i.Checkbox.AddClass(i.Size.Combine(i.BasicClassName()))
	}
	if i.ResponseStatus != "" {
		i.Checkbox.AddClass(i.ResponseStatus.Combine(i.BasicClassName()))
	}
	i.Label.AddClass("label", "cursor-pointer")
	return i.BasicComponent.Class()
}

func (i Checkbox) Html() (html htmlgo.HTML) {
	i.Label.SetText(i.Title)
	wrap := i.GetFormcontrol()
	i.Label.For = i.Checkbox.Name
	for _, kv := range i.Options {
		label := i.Label.Clone()
		checkbox := i.Checkbox.Clone()
		label.SetText(kv.Label)
		span := elements.Span{}
		span.AddClass("label-text")
		txt := elements.Text{Text: kv.Label}
		span.AddChildren(&txt)
		checkbox.Value = kv.Value
		label.AddChildren(&checkbox)
		wrap.AddChildren(&label)
	}
	wrap.AddClass(i.Class()...)

	return wrap.Html()

}
