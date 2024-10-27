package attributes

import (
	"fmt"
	"strings"

	"github.com/julvo/htmlgo/attributes"
	"github.com/suifengpiao14/funcs"
)

func NewAttrs(attrs ...*attributes.Attribute) Attrs {
	atrs := make(Attrs, 0)
	atrs.AddRef(attrs...)
	return atrs
}

// Attrs 自定义属性列表 声明 为 属性引用,方便程序中增加nil
type Attrs []*attributes.Attribute

func (atrs *Attrs) AddRef(attrs ...*attributes.Attribute) *Attrs {
	//实现函数体,支持相同属性的更新
	for _, attr := range attrs {
		if attr == nil {
			continue
		}
		exists := false
		for i, at := range *atrs {
			if at.Name == attr.Name {
				(*atrs)[i] = attr
				exists = true
				break
			}
		}
		if !exists {
			*atrs = append(*atrs, attr)
		}
	}
	return atrs
}
func (atrs *Attrs) Add(attrs ...attributes.Attribute) *Attrs {
	refs := make(Attrs, 0)
	for _, attr := range attrs {
		refs = append(refs, &attr)
	}
	return atrs.AddRef(refs...)
}

func (atrs Attrs) Attrs() []attributes.Attribute {
	var attrs []attributes.Attribute
	for _, attr := range atrs {
		attrs = append(attrs, *attr)
	}
	return attrs
}

func Attr_(name string, templs ...string) attributes.Attribute {
	tplName := funcs.ToCamel(name)
	attr := attributes.Attribute{Data: nil, Name: tplName}
	value := "{{.}}"
	if len(templs) > 0 {
		value = strings.Join(templs, " ")
	}
	attr.Templ = fmt.Sprintf(`{{define "%s"}}%s="%s"{{end}}`, tplName, name, value)
	return attr
}

// AttrHtmlGlobal html 全局属性
type AttrHtmlGlobal struct {
	Accesskey       string `json:"accesskey,omitempty"`       // Specifies a shortcut key to activate/focus an element
	ClassName       Class  `json:"class,omitempty"`           // Specifies one or more classnames for an element (refers to a class in a style sheet)
	Contenteditable string `json:"contenteditable,omitempty"` // Specifies whether the content of an element is editable or not
	Data            string `json:"data,omitempty"`            // Used to store custom data private to the page or application
	Dir             string `json:"dir,omitempty"`             // Specifies the text direction for the content in an element
	Draggable       string `json:"draggable,omitempty"`       // Specifies whether an element is draggable or not
	Enterkeyhint    string `json:"enterkeyhint,omitempty"`    // Specifies the text of the enter-key on a virtual keybord
	Hidden          string `json:"hidden,omitempty"`          // Specifies that an element is not yet, or is no longer, relevant
	Id              string `json:"id,omitempty"`              // Specifies a unique id for an element
	Inert           string `json:"inert,omitempty"`           // Specifies that the browser should ignore this section
	Inputmode       string `json:"inputmode,omitempty"`       // Specifies the mode of a virtual keyboard
	Lang            string `json:"lang,omitempty"`            // Specifies the language of the element's content
	Popover         string `json:"popover,omitempty"`         // Specifies a popover element
	Spellcheck      bool   `json:"spellcheck,omitempty"`      // Specifies whether the element is to have its spelling and grammar checked or not
	Style           string `json:"style,omitempty"`           // Specifies an inline CSS style for an element
	Tabindex        int    `json:"tabindex,omitempty"`        // Specifies the tabbing order of an element
	Title           string `json:"title,omitempty"`           // Specifies extra information about an element
	Translate       string `json:"translate,omitempty"`       // Specifies whether the content of an element should be translated or not

}

func (a AttrHtmlGlobal) Attrs() Attrs {
	attrs := NewAttrs(
		AccessKey_(a.Accesskey),
		Class.Attr((a.ClassName)),
		Contenteditable_(a.Contenteditable),
		Data_(a.Data),
		Dir_(a.Dir),
		Draggable_(a.Draggable),
		Enterkeyhint_(a.Enterkeyhint),
		Hidden_(a.Hidden),

		Id_(a.Id),
		Inert_(a.Inert),
		InputMode_(a.Inputmode),
		Lang_(a.Lang),
		Popover_(a.Popover),
		Spellcheck_(a.Spellcheck),
		Style_(a.Style),
		TabIndex_(a.Tabindex),
		Title_(a.Title),
		Translate_(a.Translate),
	)
	return attrs

}
