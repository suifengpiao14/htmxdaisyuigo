package htmxdaisyuigo

import (
	"github.com/julvo/htmlgo"
	htmlattr "github.com/suifengpiao14/htmxdaisyuigo/attributes"
)

type Input struct {
	Required    bool
	Type        string
	Class       htmlattr.Class
	Placeholder string
}

func (i Input) Html() (html string) {
	//<input required type="tel" class="grow" placeholder="请输入手机号"  />
	attrs := htmlattr.NewAttrs(
		htmlattr.Required_(i.Required),
		htmlattr.Type_(i.Type),
		htmlattr.Class.Attr(i.Class),
		htmlattr.Placeholder_(i.Placeholder),
	)
	tag := htmlgo.Input(attrs.Attrs())
	return string(tag)
}

type Label struct {
	For      string
	Class    htmlattr.Class
	children []htmlgo.HTML
}

func (l Label) Html() (html string) {
	//<label class="input px-0 input-bordered flex items-center gap-2" for="phone">手机号</label>
	attrs := htmlattr.NewAttrs(
		htmlattr.For_(l.For),
		htmlattr.Class.Attr(l.Class),
	)
	tag := htmlgo.Label(attrs.Attrs(), l.children...)
	return string(tag)
}
