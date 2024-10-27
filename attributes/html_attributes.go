package attributes

import (
	"github.com/julvo/htmlgo/attributes"
	"github.com/spf13/cast"
)

// Begin of generated *attributes.Attributes

func Accept_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Accept_(value)
	return &attr
}

func AcceptCharset_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.AcceptCharset_(value)
	return &attr
}

func AccessKey_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Accesskey_(value)
	return &attr
}

func Action_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Action_(value)
	return &attr
}

func Align_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Align_(value)
	return &attr
}

func Alt_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Alt_(value)
	return &attr
}

func AriaExpanded_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.AriaExpanded_(value)
	return &attr
}

func AriaHidden_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.AriaHidden_("aria-hidden")
	return &attr
}

func AriaLabel_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.AriaLabel_(value)
	return &attr
}

func Async_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Async_(value)
	return &attr
}

func Autocomplete_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Autocomplete_(value)
	return &attr
}

func Autofocus_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Autofocus_("autofocus")
	return &attr
}

func Autoplay_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Autoplay_("autoplay")
	return &attr
}

func Bgcolor_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Bgcolor_(value)
	return &attr
}

func Border_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Border_(value)
	return &attr
}

func Charset_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Charset_(value)
	return &attr
}

func Checked_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Checked_("checked")
	return &attr
}

func Cite_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Cite_(value)
	return &attr
}

type Class []string

func (c *Class) Add(classes ...string) {
	if *c == nil {
		*c = make([]string, 0)
	}
	for _, class := range classes {
		if class == "" {
			continue
		}
		exists := false
		for _, v := range *c {
			if v == class {
				exists = true
				break
			}
		}
		if !exists {
			*c = append(*c, class)
		}
	}
}

func (c *Class) Remove(classes ...string) {
	m := make(map[string]bool)
	for _, v := range classes {
		m[v] = true
	}

	tmpCls := make([]string, 0)
	for _, v := range *c {
		if !m[v] {
			tmpCls = append(tmpCls, v)
		}
	}
	*c = tmpCls
}

func (c *Class) Clearn() {
	*c = make(Class, 0)
}

func (c Class) Attr() *attributes.Attribute {
	if len(c) == 0 {
		return nil
	}
	attr := attributes.Class_(c...)
	return &attr
}

func Color_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Color_(value)
	return &attr
}

func Cols_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Cols_(cast.ToString(value))
	return &attr
}

func Colspan_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Colspan_(cast.ToString(value))
	return &attr
}

func Content_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Content_(value)
	return &attr
}

func Contenteditable_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Contenteditable_(value)
	return &attr
}

func Controls_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Controls_("controls")
	return &attr
}

func Coords_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Coords_(value)
	return &attr
}

func Data_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Data_(value)
	return &attr
}

func Datetime_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Datetime_(value)
	return &attr
}

func Default_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Default_(value)
	return &attr
}

func Defer_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Defer_(value)
	return &attr
}

func Dir_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Dir_(value)
	return &attr
}

func Dirname_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Dirname_(value)
	return &attr
}

func Disabled_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Disabled_("disabled")
	return &attr
}

func Download_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Download_(value)
	return &attr
}

func Draggable_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Draggable_(value)
	return &attr
}

func Dropzone_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Dropzone_(value)
	return &attr
}

func Enctype_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Enctype_(value)
	return &attr
}

func For_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.For_(value)
	return &attr
}

func Form_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Form_(value)
	return &attr
}

func FormAction_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Formaction_(value)
	return &attr
}

func FormEnctype_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("formenctype", value)
	return &attr
}
func FormMethod_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("formmethod", value)
	return &attr
}

func FormNoValidate_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := Attr_("formnovalidate", "formnovalidate")
	return &attr
}
func FormTarget_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("formtarget", value)
	return &attr
}

func Headers_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Headers_(value)
	return &attr
}

func Height_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Height_(cast.ToString(value))
	return &attr
}

func Hidden_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Hidden_(value)
	return &attr
}

func High_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.High_(value)
	return &attr
}

func Href_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Href_(value)
	return &attr
}

func Hreflang_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Hreflang_(value)
	return &attr
}

func HttpEquiv_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.HttpEquiv_(value)
	return &attr
}

func Id_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Id_(value)
	return &attr
}

func Background_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("background", value)
	return &attr
}

func InitialScale_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.InitialScale_(value)
	return &attr
}

func Ismap_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ismap_(value)
	return &attr
}

func Kind_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Kind_(value)
	return &attr
}

func Label_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Label_(value)
	return &attr
}

func Lang_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Lang_(value)
	return &attr
}

func List_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.List_(value)
	return &attr
}

func Loop_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Loop_("loop")
	return &attr
}

func Low_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Low_(value)
	return &attr
}

func Max_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Max_(value)
	return &attr
}
func MinLength_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := Attr_("minlength", cast.ToString(value))
	return &attr
}

func MaxLength_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Maxlength_(cast.ToString(value))
	return &attr
}

func Media_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Media_(value)
	return &attr
}

func HrefLang_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("hreflang", value)
	return &attr
}

func Method_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Method_(value)
	return &attr
}

func Min_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Min_(value)
	return &attr
}

func Multiple_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Multiple_("multiple")
	return &attr
}

func Muted_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Muted_("muted")
	return &attr
}

func Name_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Name_(value)
	return &attr
}

func Novalidate_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Novalidate_(value)
	return &attr
}

func Onabort_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onabort_(value)
	return &attr
}

func Onafterprint_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onafterprint_(value)
	return &attr
}

func Onbeforeprint_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onbeforeprint_(value)
	return &attr
}

func Onbeforeunload_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onbeforeunload_(value)
	return &attr
}

func Onblur_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onblur_(value)
	return &attr
}

func Oncanplay_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oncanplay_(value)
	return &attr
}

func Oncanplaythrough_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oncanplaythrough_(value)
	return &attr
}

func Onchange_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onchange_(value)
	return &attr
}

func OnClick_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onclick_(value)
	return &attr
}

func Oncontextmenu_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oncontextmenu_(value)
	return &attr
}

func Oncopy_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oncopy_(value)
	return &attr
}

func Oncuechange_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oncuechange_(value)
	return &attr
}

func Oncut_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oncut_(value)
	return &attr
}

func Ondblclick_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondblclick_(value)
	return &attr
}

func Ondrag_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondrag_(value)
	return &attr
}

func Ondragend_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondragend_(value)
	return &attr
}

func Ondragenter_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondragenter_(value)
	return &attr
}

func Ondragleave_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondragleave_(value)
	return &attr
}

func Ondragover_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondragover_(value)
	return &attr
}

func Ondragstart_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondragstart_(value)
	return &attr
}

func Ondrop_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondrop_(value)
	return &attr
}

func Ondurationchange_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ondurationchange_(value)
	return &attr
}

func Onemptied_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onemptied_(value)
	return &attr
}

func Onended_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onended_(value)
	return &attr
}

func Onerror_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onerror_(value)
	return &attr
}

func Onfocus_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onfocus_(value)
	return &attr
}

func Onhashchange_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onhashchange_(value)
	return &attr
}

func Oninput_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oninput_(value)
	return &attr
}

func Oninvalid_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Oninvalid_(value)
	return &attr
}

func Onkeydown_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onkeydown_(value)
	return &attr
}

func Onkeypress_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onkeypress_(value)
	return &attr
}

func Onkeyup_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onkeyup_(value)
	return &attr
}

func Onload_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onload_(value)
	return &attr
}

func Onloadeddata_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onloadeddata_(value)
	return &attr
}

func Onloadedmetadata_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onloadedmetadata_(value)
	return &attr
}

func Onloadstart_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onloadstart_(value)
	return &attr
}

func Onmousedown_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onmousedown_(value)
	return &attr
}

func Onmousemove_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onmousemove_(value)
	return &attr
}

func Onmouseout_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onmouseout_(value)
	return &attr
}

func Onmouseover_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onmouseover_(value)
	return &attr
}

func Onmouseup_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onmouseup_(value)
	return &attr
}

func Onmousewheel_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onmousewheel_(value)
	return &attr
}

func Onoffline_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onoffline_(value)
	return &attr
}

func Ononline_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ononline_(value)
	return &attr
}

func Onpagehide_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onpagehide_(value)
	return &attr
}

func Onpageshow_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onpageshow_(value)
	return &attr
}

func Onpaste_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onpaste_(value)
	return &attr
}

func Onpause_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onpause_(value)
	return &attr
}

func Onplay_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onplay_(value)
	return &attr
}

func Onplaying_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onplaying_(value)
	return &attr
}

func Onpopstate_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onpopstate_(value)
	return &attr
}

func Onprogress_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onprogress_(value)
	return &attr
}

func Onratechange_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onratechange_(value)
	return &attr
}

func Onreset_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onreset_(value)
	return &attr
}

func Onresize_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onresize_(value)
	return &attr
}

func Onscroll_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onscroll_(value)
	return &attr
}

func Onsearch_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onsearch_(value)
	return &attr
}

func Onseeked_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onseeked_(value)
	return &attr
}

func Onseeking_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onseeking_(value)
	return &attr
}

func Onselect_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onselect_(value)
	return &attr
}

func Onstalled_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onstalled_(value)
	return &attr
}

func Onstorage_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onstorage_(value)
	return &attr
}

func Onsubmit_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onsubmit_(value)
	return &attr
}

func Onsuspend_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onsuspend_(value)
	return &attr
}

func Ontimeupdate_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ontimeupdate_(value)
	return &attr
}

func Ontoggle_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Ontoggle_(value)
	return &attr
}

func Onunload_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onunload_(value)
	return &attr
}

func Onvolumechange_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onvolumechange_(value)
	return &attr
}

func Onwaiting_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onwaiting_(value)
	return &attr
}

func Onwheel_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Onwheel_(value)
	return &attr
}

func Open_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Open_("open")
	return &attr
}

func Optimum_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Optimum_(value)
	return &attr
}

func Pattern_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Pattern_(value)
	return &attr
}

func Placeholder_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Placeholder_(value)
	return &attr
}

func Poster_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Poster_(value)
	return &attr
}

func Preload_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Preload_(value)
	return &attr
}

func ReadOnly_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Readonly_("true")
	return &attr
}

func Rel_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Rel_(value)
	return &attr
}

func Required_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Required_("required")
	return &attr
}

func Reversed_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Reversed_("reversed")
	return &attr
}

func Role_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Role_(value)
	return &attr
}

func Rows_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Rows_(cast.ToString(value))
	return &attr
}

func Rowspan_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Rowspan_(cast.ToString(value))
	return &attr
}

func Sandbox_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Sandbox_(value)
	return &attr
}

func Scope_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Scope_(value)
	return &attr
}

func Selected_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Selected_("selected")
	return &attr
}

func Shape_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Shape_(value)
	return &attr
}

func Size_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Size_(cast.ToString(value))
	return &attr
}

func Sizes_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Sizes_(value)
	return &attr
}

func Span_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Span_(cast.ToString(value))
	return &attr
}

func Spellcheck_(value bool) *attributes.Attribute {
	if !value {
		return nil
	}
	attr := attributes.Spellcheck_("true")
	return &attr
}

func Src_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Src_(value)
	return &attr
}

func Srcdoc_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Srcdoc_(value)
	return &attr
}

func Srclang_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Srclang_(value)
	return &attr
}

func Srcset_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Srcset_(value)
	return &attr
}

func Start_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Start_(cast.ToString(value))
	return &attr
}

func Step_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Step_(value)
	return &attr
}

func Style_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Style_(value)
	return &attr
}

func TabIndex_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Tabindex_(cast.ToString(value))
	return &attr
}

func AriaDescribedby_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("aria-describedby", value)
	return &attr
}

func Target_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Target_(value)
	return &attr
}

func Title_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Title_(value)
	return &attr
}

func Translate_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Translate_(value)
	return &attr
}

const (
	Input_Type_Tel           = "tel"
	Input_Type_Text          = "text"
	Input_Type_Checkbox      = "checkbox"
	Input_Type_Radio         = "radio"
	Input_Type_Email         = "email"
	Input_Type_Password      = "password"
	Input_Type_Number        = "number"
	Input_Type_Date          = "date"
	Input_Type_DatetimeLocal = "datetime-local"
	Input_Type_Month         = "month"
	Input_Type_Week          = "week"
	Input_Type_Time          = "time"
	Input_Type_Color         = "color"
	Input_Type_Search        = "search"
	Input_Type_Url           = "url"
	Input_Type_Range         = "range"
	Input_Type_File          = "file"
	Input_Type_Hidden        = "hidden"
	Input_Type_Image         = "image"
	Input_Type_Submit        = "submit"
	Input_Type_Reset         = "reset"
	Input_Type_Button        = "button"
)

func Type_(typ string) *attributes.Attribute {
	if typ == "" {
		typ = Input_Type_Text
	}
	attr := attributes.Type_(typ)
	return &attr
}

func Usemap_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Usemap_(value)
	return &attr
}

func Value_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Value_(value)
	return &attr
}

func Width_(value int) *attributes.Attribute {
	if value == 0 {
		return nil
	}
	attr := attributes.Width_(cast.ToString(value))
	return &attr
}
func InputMode_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := Attr_("inputmode", value)
	return &attr
}

func Wrap_(value string) *attributes.Attribute {
	if value == "" {
		return nil
	}
	attr := attributes.Wrap_(value)
	return &attr
}

func Enterkeyhint_(enterkeyhint string) *attributes.Attribute {
	if enterkeyhint == "" {
		return nil
	}
	attr := Attr_("enterkeyhint", enterkeyhint)
	return &attr
}
func Inert_(inert string) *attributes.Attribute {
	if inert == "" {
		return nil
	}
	attr := Attr_("inert", inert)
	return &attr
}

func Popover_(popover string) *attributes.Attribute {
	if popover == "" {
		return nil
	}
	attr := Attr_("popover", popover)
	return &attr
}
