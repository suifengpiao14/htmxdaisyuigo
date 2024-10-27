package datainput

import (
	"github.com/julvo/htmlgo"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/daisyui/css"
	"github.com/suifengpiao14/htmxdaisyuigo/elements"
)

type Radio struct {
	Title string
	elements.BasicComponent
	Radio          elements.Radio
	Label          elements.Label
	Formcontrol    elements.ComponentI
	ResponseStatus css.ResponseStatus
	Color          css.Color
	Size           css.Size
	KVs            KVs
}

func (i *Radio) Tag() string {
	if i.Formcontrol == nil {
		return ""
	}
	return i.Formcontrol.Tag()
}

func (i *Radio) GetFormcontrol() (formcontrol elements.ComponentI) {
	if i.Formcontrol == nil {
		i.Formcontrol = &elements.Div{}
		i.Formcontrol.AddClass(css.Formcontrol)
	}
	return i.Formcontrol
}

func (i *Radio) Component() elements.ComponentI {
	return i
}

func (i *Radio) IsNil() bool {
	return false
}

func (i Radio) BasicClassName() string {
	return "radio"
}

func (i Radio) WithResponseStatus(status css.ResponseStatus) Radio {
	i.ResponseStatus = status
	return i
}

func (i Radio) WithColor(color css.Color) Radio {
	i.Color = color
	return i
}

func (i Radio) WithSize(size css.Size) Radio {
	i.Size = size
	return i
}

func (i Radio) Class() attributes.Class {
	i.Radio.AddClass(i.BasicClassName())
	if i.Color != "" {
		i.Radio.AddClass(i.Color.Combine(i.BasicClassName()))
	}
	if i.Size != "" {
		i.Radio.AddClass(i.Size.Combine(i.BasicClassName()))
	}
	if i.ResponseStatus != "" {
		i.Radio.AddClass(i.ResponseStatus.Combine(i.BasicClassName()))
	}
	i.Label.AddClass("label", "cursor-pointer")
	return i.BasicComponent.Class()
}

func (i Radio) Html() (html htmlgo.HTML) {
	i.Label.SetText(i.Title)
	wrap := i.GetFormcontrol()
	i.Label.For = i.Radio.Name
	for _, kv := range i.KVs {
		label := i.Label.Clone()
		checkbox := i.Radio.Clone()
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
