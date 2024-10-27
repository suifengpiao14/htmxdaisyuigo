package elements

import (
	"sort"

	"github.com/go-playground/validator/v10"
	"github.com/julvo/htmlgo"
	htmlattr "github.com/julvo/htmlgo/attributes"
	"github.com/suifengpiao14/htmxdaisyuigo/attributes"
)

func validate(i any) (err error) {
	validate := validator.New()
	return validate.Struct(i)
}

type ComponentI interface {
	Tag() string
	SerialNumber() SerialNumber
	SetSerialNumber(serialNumber SerialNumber)
	Component() ComponentI
	Class() attributes.Class
	Attrs() attributes.Attrs
	Children() Components
	AddClass(classnames ...string)
	AddAttrs(attrs ...*htmlattr.Attribute)
	AddChildren(children ...ComponentI)
	Html() htmlgo.HTML
	IsNil() bool
}

type Components []ComponentI

func (a Components) Html() []htmlgo.HTML {
	sort.Sort(a)
	htmls := make([]htmlgo.HTML, 0)
	for _, c := range a {
		htmls = append(htmls, c.Html())
	}
	return htmls
}

func (a *Components) Add(cs ...ComponentI) {
	for _, c := range cs {
		if c.IsNil() {
			continue
		}
		*a = append(*a, c)
	}

}

func (a Components) Len() int           { return len(a) }
func (a Components) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Components) Less(i, j int) bool { return a[i].SerialNumber() < a[j].SerialNumber() }

// BasicComponent  组件基础,主要用于收集text,class 和children 方便延迟修改
type BasicComponent struct {
	tag          string
	class        attributes.Class
	text         string
	attrs        attributes.Attrs
	children     Components
	serialNumber SerialNumber
}

func (i *BasicComponent) Tag() string {
	if i.tag == "" {
		panic("Tag required  for component ")
	}
	return i.tag

}

func (i *BasicComponent) SetTag(tag string) {
	i.tag = tag

}

func (i *BasicComponent) SetSerialNumber(serialNumber SerialNumber) {
	i.serialNumber = serialNumber
}

func (i *BasicComponent) SerialNumber() (serialNumber SerialNumber) {
	return i.serialNumber
}

func (i *BasicComponent) Component() ComponentI {
	panic("implement Component method")
}

func (i BasicComponent) Class() attributes.Class {
	return i.class
}
func (i *BasicComponent) AddClass(classnames ...string) {
	i.class.Add(classnames...)
}
func (i BasicComponent) Text() string {
	return i.text
}
func (i *BasicComponent) SetText(text string) {
	i.text = text
}
func (i BasicComponent) Attrs() attributes.Attrs {
	i.attrs.AddRef(i.Class().Attr())
	return i.attrs
}
func (i *BasicComponent) AddAttrs(attrs ...*htmlattr.Attribute) {
	i.attrs = attrs
}
func (i BasicComponent) Children() Components {
	return i.children
}
func (i *BasicComponent) AddChildren(children ...ComponentI) {
	i.children.Add(children...)
}
func (i *BasicComponent) IsNil() bool {
	panic("implement IsNil method")
}

func (i *BasicComponent) Html() htmlgo.HTML {
	subComponents := Components{}
	subComponents.Add(&Text{Text: i.Text()})
	subComponents.Add(i.Children()...)
	children := make([]htmlgo.HTML, 0)
	sort.Sort(subComponents)
	for _, subComponent := range subComponents {
		children = append(children, subComponent.Html())
	}
	element := htmlgo.Element(i.Tag(), i.Attrs().Attrs(), children...)
	return element
}

type SerialNumber int

type Text struct {
	serialNumber SerialNumber
	Text         string
}

func (i *Text) Component() ComponentI {
	return i
}
func (i *Text) Tag() string {
	return ""
}
func (i *Text) IsNil() bool {
	return i.Text == ""
}
func (i *Text) SerialNumber() SerialNumber {
	return i.serialNumber
}

func (i *Text) SetSerialNumber(serialNumber SerialNumber) {
	i.serialNumber = serialNumber
}
func (i *Text) Class() attributes.Class {
	return attributes.Class{}
}
func (i *Text) AddClass(classnames ...string) {
}

func (i *Text) Attrs() attributes.Attrs {
	return attributes.Attrs{}
}
func (i *Text) Children() Components {
	return Components{}
}

func (i *Text) AddAttrs(attrs ...*htmlattr.Attribute) {
}
func (i *Text) AddChildren(children ...ComponentI) {
}
func (i *Text) Html() htmlgo.HTML {
	return htmlgo.Text(i.Text)
}

type Checkbox struct {
	Input
}

func (i *Checkbox) Component() ComponentI {
	return i
}

func (i *Checkbox) Attrs() attributes.Attrs {
	typ := htmlattr.Type_(attributes.Input_Type_Checkbox)
	i.Input.AddAttrs(&typ)
	return i.Input.Attrs()

}

type Radio struct {
	Input
}

func (i *Radio) Component() ComponentI {
	return i
}

func (i *Radio) Attrs() attributes.Attrs {
	typ := htmlattr.Type_(attributes.Input_Type_Radio)
	i.Input.AddAttrs(&typ)
	return i.Input.Attrs()

}

type Input struct {
	attributes.AttrHtmlGlobal
	Type           string `json:"type,omitempty"`           // 输入类型，如 text, password, checkbox 等
	Name           string `json:"name" validate:"required"` // 表单数据的字段名称
	Value          string `json:"value,omitempty"`          // 输入字段的初始值
	Placeholder    string `json:"placeholder,omitempty"`    // 输入字段的提示文字
	Required       bool   `json:"required,omitempty"`       // 是否为必填字段
	ReadOnly       bool   `json:"readonly,omitempty"`       // 是否为只读字段
	Disabled       bool   `json:"disabled,omitempty"`       // 是否禁用输入
	Autocomplete   string `json:"autocomplete,omitempty"`   // 浏览器自动填充选项
	Autofocus      bool   `json:"autofocus,omitempty"`      // 页面加载时自动聚焦
	Min            string `json:"min,omitempty"`            // 最小值，适用于 number, date 等类型
	Max            string `json:"max,omitempty"`            // 最大值，适用于 number, date 等类型
	MinLength      int    `json:"minlength,omitempty"`      // 文本输入的最小字符数
	MaxLength      int    `json:"maxlength,omitempty"`      // 文本输入的最大字符数
	Pattern        string `json:"pattern,omitempty"`        // 正则表达式模式匹配
	Step           string `json:"step,omitempty"`           // 步长，适用于 number, range 等类型
	Form           string `json:"form,omitempty"`           // 指定字段所属表单的 ID
	FormAction     string `json:"formaction,omitempty"`     // 表单提交时的 URL，适用于 submit 按钮
	FormEnctype    string `json:"formenctype,omitempty"`    // 表单数据的编码类型
	FormMethod     string `json:"formmethod,omitempty"`     // 表单提交时的 HTTP 方法
	FormNoValidate bool   `json:"formnovalidate,omitempty"` // 是否跳过表单验证，适用于 submit 按钮
	FormTarget     string `json:"formtarget,omitempty"`     // 表单提交的目标浏览上下文
	Accept         string `json:"accept,omitempty"`         // 上传文件的 MIME 类型
	Size           int    `json:"size,omitempty"`           // 输入字段的可见宽度
	Multiple       bool   `json:"multiple,omitempty"`       // 是否允许选择多个文件/选项
	List           string `json:"list,omitempty"`           // 绑定 <datalist> 的 ID，用于自动完成选项
	Height         int    `json:"height,omitempty"`         // 图片按钮的高度
	Width          int    `json:"width,omitempty"`          // 图片按钮的宽度
	InputMode      string `json:"inputmode,omitempty"`      // 虚拟键盘的输入模式
	Spellcheck     bool   `json:"spellcheck,omitempty"`     // 是否检查拼写和语法
	Title          string `json:"title,omitempty"`          // 鼠标悬停时显示的文字
	BasicComponent
}

func (i *Input) Component() ComponentI {
	return i
}
func (i *Input) Tag() string {
	return "input"
}
func (i *Input) IsNil() bool {
	return false
}

func (i Input) Clone() Input {
	return i
}

func (i Input) Attrs() attributes.Attrs {
	attrs := i.AttrHtmlGlobal.Attrs()
	if i.Type == "" {
		i.Type = attributes.Input_Type_Text
	}
	attrs.AddRef(
		attributes.Required_(i.Required),
		attributes.Type_(i.Type),
		attributes.Class.Attr((i.AttrHtmlGlobal.ClassName)),
		attributes.Placeholder_(i.Placeholder),
		attributes.Value_(i.Value),
		attributes.Name_(i.Name),
		attributes.ReadOnly_(i.ReadOnly),
		attributes.Disabled_(i.Disabled),
		attributes.Autocomplete_(i.Autocomplete),
		attributes.Autofocus_(i.Autofocus),
		attributes.Min_(i.Min),
		attributes.Max_(i.Max),
		attributes.MinLength_(i.MinLength),
		attributes.MaxLength_(i.MaxLength),
		attributes.Pattern_(i.Pattern),
		attributes.Step_(i.Step),
		attributes.Form_(i.Form),
		attributes.FormAction_(i.FormAction),
		attributes.FormEnctype_(i.FormEnctype),
		attributes.FormMethod_(i.FormMethod),
		attributes.FormNoValidate_(i.FormNoValidate),
		attributes.FormTarget_(i.FormTarget),
		attributes.Accept_(i.Accept),
		attributes.Size_(i.Size),
		attributes.Multiple_(i.Multiple),
		attributes.List_(i.List),
		attributes.Height_(i.Height),
		attributes.Width_(i.Width),
		attributes.InputMode_(i.InputMode),
		attributes.Spellcheck_(i.Spellcheck),
		attributes.Title_(i.Title),
		attributes.Id_(i.Id),
	)
	attrs.AddRef(i.BasicComponent.Attrs()...)
	return attrs
}

func (i Input) Html() (html htmlgo.HTML) {
	attrs := i.Attrs()
	return htmlgo.Input(attrs.Attrs())
}

type A struct {
	Href            string           `json:"href,omitempty"`             // 超链接目标 URL
	Target          string           `json:"target,omitempty"`           // 打开链接的方式 (_self, _blank, _parent, _top)
	Rel             string           `json:"rel,omitempty"`              // 与链接目标的关系
	Download        string           `json:"download,omitempty"`         // 下载链接内容的建议文件名
	Media           string           `json:"media,omitempty"`            // 适用的媒体类型
	HrefLang        string           `json:"hreflang,omitempty"`         // 目标 URL 的语言
	Title           string           `json:"title,omitempty"`            // 链接的提示文本
	AccessKey       string           `json:"accesskey,omitempty"`        // 激活链接的快捷键
	TabIndex        int              `json:"tabindex,omitempty"`         // 链接的 Tab 键顺序
	Id              string           `json:"id,omitempty"`               // 元素的唯一标识符
	Class           attributes.Class `json:"class,omitempty"`            // CSS 类名
	Style           string           `json:"style,omitempty"`            // 行内 CSS 样式
	AriaLabel       string           `json:"aria-label,omitempty"`       // 可访问性标签
	AriaDescribedby string           `json:"aria-describedby,omitempty"` // 描述元素的 ID
	AriaHidden      bool             `json:"aria-hidden,omitempty"`      // 指示元素是否可见
	OnClick         string           `json:"onclick,omitempty"`          // 点击事件的 JavaScript 代码
}

func (a A) Validate() (err error) {
	return validate(a)
}
func (a A) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Href_(a.Href),
		attributes.Target_(a.Target),
		attributes.Rel_(a.Rel),
		attributes.Download_(a.Download),
		attributes.Media_(a.Media),
		attributes.HrefLang_(a.HrefLang),
		attributes.Title_(a.Title),
		attributes.AccessKey_(a.AccessKey),
		attributes.TabIndex_(a.TabIndex),
		attributes.Id_(a.Id),
		attributes.Class.Attr((a.Class)),
		attributes.Style_(a.Style),
		attributes.AriaLabel_(a.AriaLabel),
		attributes.AriaDescribedby_(a.AriaDescribedby),
		attributes.AriaHidden_(a.AriaHidden),
		attributes.OnClick_(a.OnClick),
	)
	tag := htmlgo.A(attrs.Attrs())
	return string(tag)
}

type Abbr struct {
	Title string           `json:"title,omitempty"` // 解释缩写的标题
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Acronym struct {
	Title string           `json:"title,omitempty"` // 解释缩写的标题
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Address struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Applet struct {
	Code   string           `json:"code,omitempty"`   // Applet 的代码
	Width  int              `json:"width,omitempty"`  // Applet 的宽度
	Height int              `json:"height,omitempty"` // Applet 的高度
	Alt    string           `json:"alt,omitempty"`    // 替代文本
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

type Area struct {
	Href   string           `json:"href,omitempty"`   // 链接的 URL
	Alt    string           `json:"alt,omitempty"`    // 替代文本
	Shape  string           `json:"shape,omitempty"`  // 形状
	Coords string           `json:"coords,omitempty"` // 坐标
	Target string           `json:"target,omitempty"` // 链接目标
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

func (a Area) Validate() (err error) {
	return validate(a)
}

func (a Area) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Href_(a.Href),
		attributes.Alt_(a.Alt),
		attributes.Shape_(a.Shape),
		attributes.Coords_(a.Coords),
		attributes.Target_(a.Target),
		attributes.Id_(a.Id),
		attributes.Class.Attr((a.Class)),
	)
	tag := htmlgo.Area(attrs.Attrs())
	return string(tag)
}

type Article struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (a Article) Validate() (err error) {
	return validate(a)
}
func (a Article) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(a.Id),
		attributes.Class.Attr((a.Class)),
	)
	tag := htmlgo.Aside(attrs.Attrs())
	return string(tag)
}

type Aside struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (a Aside) Validate() (err error) {
	return validate(a)
}
func (a Aside) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(a.Id),
		attributes.Class.Attr((a.Class)),
	)
	tag := htmlgo.Aside(attrs.Attrs())
	return string(tag)
}

type Audio struct {
	Src      string `json:"src,omitempty"`      // 音频文件 URL
	Controls bool   `json:"controls,omitempty"` // 是否显示控件
	Autoplay bool   `json:"autoplay,omitempty"` // 是否自动播放
	Loop     bool   `json:"loop,omitempty"`     // 是否循环播放
	Muted    bool   `json:"muted,omitempty"`    // 是否静音
	Class    string `json:"class,omitempty"`    // CSS 类名
}

func (a Audio) Validate() (err error) {
	return validate(a)
}
func (a Audio) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Src_(a.Src),
		attributes.Controls_(a.Controls),
		attributes.Autoplay_(a.Autoplay),
		attributes.Loop_(a.Loop),
		attributes.Muted_(a.Muted),
	)
	tag := htmlgo.Audio(attrs.Attrs())
	return string(tag)
}

type B struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Base struct {
	Href  string           `json:"href,omitempty"`  // 基础 URL
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Basefont struct {
	Size  string           `json:"size,omitempty"`  // 字体大小
	Color string           `json:"color,omitempty"` // 字体颜色
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Bdi struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Bdo struct {
	Dir   string           `json:"dir,omitempty"`   // 文本方向
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Bgsound struct {
	Src   string           `json:"src,omitempty"`   // 音频文件 URL
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Big struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Blink struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Blockquote struct {
	Cite  string           `json:"cite,omitempty"`  // 引用来源
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (b Blockquote) Validate() (err error) {
	return validate(b)
}
func (b Blockquote) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Cite_(b.Cite),
		attributes.Id_(b.Id),
		attributes.Class.Attr((b.Class)),
	)
	tag := htmlgo.Blockquote(attrs.Attrs())
	return string(tag)
}

type Body struct {
	Background string           `json:"background,omitempty"` // 背景图像 URL
	Class      attributes.Class `json:"class,omitempty"`      // CSS 类名
}

func (b Body) Validate() (err error) {
	return validate(b)
}
func (b Body) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Background_(b.Background),
		attributes.Class.Attr((b.Class)),
	)
	tag := htmlgo.Body(attrs.Attrs())
	return string(tag)
}

type Br struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Button struct {
	Type     string           `json:"type,omitempty"`     // 按钮类型
	Name     string           `json:"name,omitempty"`     // 按钮名称
	Value    string           `json:"value,omitempty"`    // 按钮值
	Disabled bool             `json:"disabled,omitempty"` // 是否禁用按钮
	Class    attributes.Class `json:"class,omitempty"`    // CSS 类名
}

func (b Button) Validate() (err error) {
	return validate(b)
}
func (b Button) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Type_(b.Type),
		attributes.Name_(b.Name),
		attributes.Value_(b.Value),
		attributes.Disabled_(b.Disabled),
		attributes.Class.Attr((b.Class)),
	)
	tag := htmlgo.Body(attrs.Attrs())
	return string(tag)
}

type Canvas struct {
	Width  int              `json:"width,omitempty"`  // 画布宽度
	Height int              `json:"height,omitempty"` // 画布高度
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

type Caption struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Center struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Cite struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Code struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (c Code) Validate() (err error) {
	return validate(c)
}
func (c Code) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(c.Id),
		attributes.Class.Attr((c.Class)),
	)
	tag := htmlgo.Code(attrs.Attrs())
	return string(tag)
}

type Col struct {
	Span  int              `json:"span,omitempty"`  // 列的跨度
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (c Col) Validate() (err error) {
	return validate(c)
}
func (c Col) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Span_(c.Span),
		attributes.Id_(c.Id),
		attributes.Class.Attr((c.Class)),
	)
	tag := htmlgo.Col(attrs.Attrs())
	return string(tag)
}

type Colgroup struct {
	Span  int              `json:"span,omitempty"`  // 列组的跨度
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Datalist struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Dd struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Del struct {
	Cite  string           `json:"cite,omitempty"`  // 被删除内容的来源
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Details struct {
	Open  bool             `json:"open,omitempty"`  // 是否展开
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (d Details) Validate() (err error) {
	return validate(d)
}
func (d Details) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Open_(d.Open),
		attributes.Id_(d.Id),
		attributes.Class.Attr((d.Class)),
	)
	tag := htmlgo.Details(attrs.Attrs())
	return string(tag)
}

type Dfn struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Dir struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Div struct {
	attributes.AttrHtmlGlobal
	Title           string `json:"title,omitempty"`            // 提供额外信息的提示文本
	AriaLabel       string `json:"aria-label,omitempty"`       // 可访问性标签
	AriaDescribedby string `json:"aria-describedby,omitempty"` // 描述元素的 ID
	AriaHidden      bool   `json:"aria-hidden,omitempty"`      // 指示元素是否可见
	OnClick         string `json:"onclick,omitempty"`          // 点击事件的 JavaScript 代码
	Role            string `json:"role,omitempty"`             // 角色属性，定义元素的语义
	BasicComponent
}

func (d Div) IsNil() bool {
	return false
}

func (d *Div) Component() (c ComponentI) {
	return d
}
func (d Div) Tag() string {
	return "div"
}

func (d Div) Attrs() (attrs attributes.Attrs) {
	attrs = d.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Title_(d.Title),
		attributes.AriaLabel_(d.AriaLabel),
		attributes.AriaDescribedby_(d.AriaDescribedby),
		attributes.AriaHidden_(d.AriaHidden),
		attributes.OnClick_(d.OnClick),
		attributes.Role_(d.Role),
	)
	attrs.AddRef(d.BasicComponent.Attrs()...)
	return attrs
}
func (d Div) Html() (html htmlgo.HTML) {
	d.BasicComponent.SetTag(d.Tag())
	d.BasicComponent.AddAttrs(d.Attrs()...)
	return d.BasicComponent.Html()
}

type Dl struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"` // 子元素
}

func (d Dl) Attrs() (attrs attributes.Attrs) {
	return d.AttrHtmlGlobal.Attrs()
}

func (d Dl) Html() (html string) {

	tag := htmlgo.Dl(d.Attrs().Attrs(), d.Children...)
	return string(tag)
}

type Dt struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"` // 子元素
}

func (d Dt) Attrs() (attrs attributes.Attrs) {
	return d.AttrHtmlGlobal.Attrs()
}

func (d Dt) Html() (html string) {
	tag := htmlgo.Dt(d.Attrs().Attrs(), d.Children...)
	return string(tag)
}

type DD struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"` // 子元素
}

func (d DD) Attrs() (attrs attributes.Attrs) {
	return d.AttrHtmlGlobal.Attrs()
}

func (d DD) Html() (html string) {
	tag := htmlgo.Dt(d.Attrs().Attrs(), d.Children...)
	return string(tag)
}

type Em struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Embed struct {
	Src    string           `json:"src,omitempty"`    // 嵌入资源的 URL
	Width  int              `json:"width,omitempty"`  // 嵌入资源的宽度
	Height int              `json:"height,omitempty"` // 嵌入资源的高度
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

type Fieldset struct {
	Disabled bool   `json:"disabled,omitempty"` // 是否禁用字段集
	Legend   string `json:"legend,omitempty"`   // 字段集的标题
	Class    string `json:"class,omitempty"`    // CSS 类名
}

type Figcaption struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Figure struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (f Figure) Validate() (err error) {
	return validate(f)
}
func (f Figure) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(f.Id),
		attributes.Class.Attr((f.Class)),
	)
	tag := htmlgo.Figure(attrs.Attrs())
	return string(tag)

}

type Font struct {
	Size  string           `json:"size,omitempty"`  // 字体大小
	Color string           `json:"color,omitempty"` // 字体颜色
	Face  string           `json:"face,omitempty"`  // 字体名称
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (f Font) Validate() (err error) {
	return validate(f)
}
func (f Font) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(f.Id),
		attributes.Class.Attr((f.Class)),
	)
	tag := htmlgo.Font(attrs.Attrs())
	return string(tag)
}

type Footer struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (f Footer) Validate() (err error) {
	return validate(f)
}
func (f Footer) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(f.Id),
		attributes.Class.Attr((f.Class)),
	)
	tag := htmlgo.Footer(attrs.Attrs())
	return string(tag)
}

type Form struct {
	Action string           `json:"action,omitempty"` // 表单提交 URL
	Method string           `json:"method,omitempty"` // 提交方法
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

func (f Form) Validate() (err error) {
	return validate(f)
}
func (f Form) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(f.Id),
		attributes.Class.Attr((f.Class)),
		attributes.Action_(f.Action),
		attributes.Method_(f.Method),
	)
	tag := htmlgo.Form(attrs.Attrs())
	return string(tag)
}

type Frame struct {
	Name  string           `json:"name,omitempty"`  // 框架名称
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (f Frame) Validate() (err error) {
	return validate(f)
}
func (f Frame) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(f.Id),
		attributes.Class.Attr((f.Class)),
		attributes.Name_(f.Name),
	)
	tag := htmlgo.Frame(attrs.Attrs())
	return string(tag)
}

type Frameset struct {
	attributes.AttrHtmlGlobal
	Cols int `json:"cols,omitempty"` // 列定义
	Rows int `json:"rows,omitempty"` // 行定义
}

func (f Frameset) Attrs() attributes.Attrs {
	attrs := f.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Cols_(f.Cols),
		attributes.Rows_(f.Rows),
	)
	return attrs
}
func (f Frameset) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(f.Id),
		attributes.Class.Attr((f.ClassName)),
		attributes.Cols_(f.Cols),
		attributes.Rows_(f.Rows),
	)
	tag := htmlgo.Frameset(attrs.Attrs())
	return string(tag)

}

type H1 struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h H1) Validate() (err error) {
	return validate(h)
}
func (h H1) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.H1(attrs.Attrs())
	return string(tag)
}

type H2 struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h H2) Validate() (err error) {
	return validate(h)
}
func (h H2) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.H2(attrs.Attrs())
	return string(tag)

}

type H3 struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h H3) Validate() (err error) {
	return validate(h)
}
func (h H3) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.H3(attrs.Attrs())
	return string(tag)
}

type H4 struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h H4) Validate() (err error) {
	return validate(h)
}
func (h H4) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.H4(attrs.Attrs())
	return string(tag)
}

type H5 struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h H5) Validate() (err error) {
	return validate(h)
}
func (h H5) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.H5(attrs.Attrs())
	return string(tag)
}

type H6 struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h H6) Validate() (err error) {
	return validate(h)
}
func (h H6) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.H6(attrs.Attrs())
	return string(tag)
}

type Head struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Head) Validate() (err error) {
	return validate(h)
}
func (h Head) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Head(attrs.Attrs())
	return string(tag)
}

type Header struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Header) Validate() (err error) {
	return validate(h)
}
func (h Header) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Header(attrs.Attrs())
	return string(tag)
}

type Hgroup struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Hgroup) Validate() (err error) {
	return validate(h)
}
func (h Hgroup) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Hgroup(attrs.Attrs())
	return string(tag)
}

type Hr struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Hr) Validate() (err error) {
	return validate(h)
}
func (h Hr) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Hr(attrs.Attrs())
	return string(tag)
}

type Html struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Html) Validate() (err error) {
	return validate(h)
}
func (h Html) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Html(attrs.Attrs(), h.Children...)
	return string(tag)
}

type I struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Iframe struct {
	Src    string           `json:"src,omitempty"`    // 内嵌页面的 URL
	Width  int              `json:"width,omitempty"`  // 宽度
	Height int              `json:"height,omitempty"` // 高度
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

type Img struct {
	Src    string           `json:"src,omitempty"`    // 图片 URL
	Alt    string           `json:"alt,omitempty"`    // 替代文本
	Width  int              `json:"width,omitempty"`  // 宽度
	Height int              `json:"height,omitempty"` // 高度
	Id     string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class  attributes.Class `json:"class,omitempty"`  // CSS 类名
}

func (h Img) Validate() (err error) {
	return validate(h)
}
func (h Img) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Img(attrs.Attrs())
	return string(tag)
}

type Ins struct {
	Cite  string           `json:"cite,omitempty"`  // 插入的来源
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Isindex struct {
	// 指示用户输入的字段，已经被弃用
}

type Kbd struct {
	// 键盘文本，没有额外属性
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Keygen struct {
	Name  string           `json:"name,omitempty"`  // 字段名称
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
	// 其他属性可以根据需要添加
}

type Label struct {
	attributes.AttrHtmlGlobal
	For string `json:"for,omitempty"` // 关联输入字段的 ID
	BasicComponent
}

func (h *Label) Component() ComponentI {
	return h
}
func (h *Label) Tag() (tag string) {
	return "label"
}
func (h *Label) IsNil() bool {
	return h.Text() == ""
}
func (h Label) Clone() Label {
	return h
}

func (h Label) Attrs() (attrs attributes.Attrs) {
	attrs = h.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.For_(h.For),
	)
	attrs.AddRef(h.BasicComponent.attrs...)
	return attrs
}
func (h Label) Html() (html htmlgo.HTML) {
	h.BasicComponent.attrs.AddRef(h.Attrs()...)
	h.BasicComponent.SetTag(h.Tag())
	return h.BasicComponent.Html()
}

type Legend struct {
	// 字段集的标题，没有额外属性
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Li struct {
	Value    int              `json:"value,omitempty"` // 列表项的值
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Li) Validate() (err error) {
	return validate(h)
}
func (h Li) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Li(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Link struct {
	Href  string           `json:"href,omitempty"`  // 关联的样式表 URL
	Rel   string           `json:"rel,omitempty"`   // 关系属性
	Media string           `json:"media,omitempty"` // 媒体类型
	Type  string           `json:"type,omitempty"`  // MIME 类型
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Link) Validate() (err error) {
	return validate(h)
}
func (h Link) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
		attributes.Href_(h.Href),
		attributes.Rel_(h.Rel),
		attributes.Media_(h.Media),
		attributes.Type_(h.Type),
	)
	tag := htmlgo.Link(attrs.Attrs())
	return string(tag)
}

type Listing struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Listing) Validate() (err error) {
	return validate(h)
}
func (h Listing) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Listing(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Main struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Map struct {
	Name  string           `json:"name,omitempty"`  // 地图名称
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Mark struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Marquee struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Menu struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Menu) Validate() (err error) {
	return validate(h)
}
func (h Menu) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Menu(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Meta struct {
	Name    string `json:"name,omitempty"`    // 元数据的名称
	Content string `json:"content,omitempty"` // 元数据的内容
	Charset string `json:"charset,omitempty"` // 字符集
	Class   string `json:"class,omitempty"`   // CSS 类名
}

func (h Meta) Validate() (err error) {
	return validate(h)
}
func (h Meta) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Name_(h.Name),
		attributes.Content_(h.Content),
		attributes.Charset_(h.Charset),
	)
	tag := htmlgo.Meta(attrs.Attrs())
	return string(tag)
}

type Meter struct {
	Value   string           `json:"value,omitempty"`   // 当前值
	Min     string           `json:"min,omitempty"`     // 最小值
	Max     string           `json:"max,omitempty"`     // 最大值
	Low     string           `json:"low,omitempty"`     // 低值
	High    string           `json:"high,omitempty"`    // 高值
	Optimum string           `json:"optimum,omitempty"` // 最佳值
	Class   attributes.Class `json:"class,omitempty"`   // CSS 类名
}

func (h Meter) Validate() (err error) {
	return validate(h)
}
func (h Meter) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Value_(h.Value),
		attributes.Min_(h.Min),
		attributes.Max_(h.Max),
		attributes.Low_(h.Low),
		attributes.High_(h.High),
		attributes.Optimum_(h.Optimum),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Meter(attrs.Attrs())
	return string(tag)
}

type Nav struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Nav) Validate() (err error) {
	return validate(h)
}
func (h Nav) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Nav(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Nobr struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Noframes struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Noscript struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Object struct {
	Data     string           `json:"data,omitempty"`   // 嵌入对象的 URL
	Type     string           `json:"type,omitempty"`   // 对象的 MIME 类型
	Width    int              `json:"width,omitempty"`  // 对象宽度
	Height   int              `json:"height,omitempty"` // 对象高度
	Id       string           `json:"id,omitempty"`     // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"`  // CSS 类名
	Children []htmlgo.HTML    `json:"object,omitempty"` // 嵌入对象
}

func (h Object) Validate() (err error) {
	return validate(h)
}
func (h Object) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Data_(h.Data),
		attributes.Type_(h.Type),
		attributes.Width_(h.Width),
		attributes.Height_(h.Height),
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Object(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Ol struct {
	Reversed bool          `json:"reversed,omitempty"` // 是否倒序
	Start    int           `json:"start,omitempty"`    // 列表起始值
	Class    string        `json:"class,omitempty"`    // CSS 类名
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (h Ol) Validate() (err error) {
	return validate(h)
}
func (h Ol) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Reversed_(h.Reversed),
		attributes.Start_(h.Start),
	)
	tag := htmlgo.Ol(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Optgroup struct {
	Label string           `json:"label,omitempty"` // 组标签
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Option struct {
	attributes.AttrHtmlGlobal
	Value    string `json:"value,omitempty"`    // 选项的值
	Label    string `json:"label,omitempty"`    // 选项的标签
	Selected bool   `json:"selected,omitempty"` // 是否选中
	Disabled bool   `json:"disabled,omitempty"` // 是否禁用
	BasicComponent
}

func (i Option) Attrs() attributes.Attrs {
	attrs := i.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Value_(i.Value),
		attributes.Selected_(i.Selected),
		attributes.Disabled_(i.Disabled),
	)
	attrs.AddRef(i.BasicComponent.Attrs()...)
	return attrs
}
func (i *Option) Children() Components {
	i.BasicComponent.AddChildren(&Text{Text: i.Label})
	return i.BasicComponent.Children()
}

func (i *Option) Component() ComponentI {
	return i
}
func (i *Option) Tag() string {
	return "option"
}
func (i *Option) IsNil() bool {
	return false
}

func (i Option) Clone() Option {
	return i
}

func (i Option) Html() (html htmlgo.HTML) {
	attrs := i.Attrs()
	children := i.Children().Html()
	return htmlgo.Option(attrs.Attrs(), children...)
}

type Output struct {
	For   string           `json:"for,omitempty"`   // 关联的输入字段
	Name  string           `json:"name,omitempty"`  // 输出字段名称
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Output) Validate() (err error) {
	return validate(h)
}
func (h Output) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.For_(h.For),
		attributes.Name_(h.Name),
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Output(attrs.Attrs())
	return string(tag)
}

type P struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h P) Validate() (err error) {
	return validate(h)
}
func (h P) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.P(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Param struct {
	Name  string           `json:"name,omitempty"`  // 参数名称
	Value string           `json:"value,omitempty"` // 参数值
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Param) Validate() (err error) {
	return validate(h)
}
func (h Param) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Name_(h.Name),
		attributes.Value_(h.Value),
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Param(attrs.Attrs())
	return string(tag)
}

type Plaintext struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Plaintext) Validate() (err error) {
	return validate(h)
}
func (h Plaintext) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Plaintext(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Pre struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Pre) Validate() (err error) {
	return validate(h)
}
func (h Pre) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Pre(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Progress struct {
	Value string           `json:"value,omitempty"` // 当前进度
	Max   string           `json:"max,omitempty"`   // 最大进度
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Progress) Validate() (err error) {
	return validate(h)
}
func (h Progress) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Value_(h.Value),
		attributes.Max_(h.Max),
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Progress(attrs.Attrs())
	return string(tag)

}

type Q struct {
	Cite     string           `json:"cite,omitempty"`  // 引用的来源
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Q) Validate() (err error) {
	return validate(h)
}
func (h Q) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Cite_(h.Cite),
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Q(attrs.Attrs())
	return string(tag)
}

type Rp struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Rt struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Ruby struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type S struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Samp struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Section struct {
	Id       string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class    attributes.Class `json:"class,omitempty"` // CSS 类名
	Children []htmlgo.HTML    `json:"children,omitempty"`
}

func (h Section) Validate() (err error) {
	return validate(h)
}
func (h Section) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Section(attrs.Attrs(), h.Children...)
	return string(tag)
}

type Select struct {
	attributes.AttrHtmlGlobal
	Name     string `json:"name,omitempty"`     // 关联的输入字段名称
	Multiple bool   `json:"multiple,omitempty"` // 是否多选
	Size     int    `json:"size,omitempty"`     // 显示的选项数量
	BasicComponent
	Options []Option
}

func (i *Select) Component() ComponentI {
	return i
}
func (i *Select) Tag() string {
	return "select"
}
func (i *Select) IsNil() bool {
	return false
}

func (i Select) Clone() Select {
	return i
}

func (i Select) Attrs() attributes.Attrs {
	attrs := i.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Name_(i.Name),
		attributes.Size_(i.Size),
		attributes.Multiple_(i.Multiple),
	)
	attrs.AddRef(i.BasicComponent.Attrs()...)
	return attrs
}

func (i Select) Html() (html htmlgo.HTML) {
	attrs := i.Attrs()
	children := i.Children().Html()
	return htmlgo.Section(attrs.Attrs(), children...)
}

type Small struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Source struct {
	Src   string           `json:"src,omitempty"`   // 资源 URL
	Type  string           `json:"type,omitempty"`  // MIME 类型
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Source) Validate() (err error) {
	return validate(h)
}
func (h Source) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Src_(h.Src),
		attributes.Type_(h.Type),
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Source(attrs.Attrs())
	return string(tag)
}

type Spacer struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Spacer) Validate() (err error) {
	return validate(h)
}
func (h Spacer) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Spacer(attrs.Attrs())
	return string(tag)
}

type Span struct {
	BasicComponent
	attributes.AttrHtmlGlobal
}

func (i *Span) Component() ComponentI {
	return i
}
func (i *Span) Tag() string {
	return ""
}
func (i *Span) IsNil() bool {
	return false
}

func (i Span) Attrs() attributes.Attrs {
	attrs := i.AttrHtmlGlobal.Attrs()
	attrs.AddRef(i.BasicComponent.Attrs()...)
	return attrs
}
func (i Span) Html() (html htmlgo.HTML) {
	attrs := i.Attrs()
	i.BasicComponent.AddAttrs(attrs...)
	return i.BasicComponent.Html()
}

type Strike struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Strong struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Style struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

func (h Style) Validate() (err error) {
	return validate(h)
}
func (h Style) Html() (html string) {
	attrs := attributes.NewAttrs(
		attributes.Id_(h.Id),
		attributes.Class.Attr((h.Class)),
	)
	tag := htmlgo.Style(attrs.Attrs())
	return string(tag)
}

type Sub struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Summary struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Sup struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Table struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (h Table) Validate() (err error) {
	return validate(h)
}

func (h Table) Attrs() attributes.Attrs {
	return h.AttrHtmlGlobal.Attrs()
}

func (h Table) Html() (html string) {
	tag := htmlgo.Table(h.Attrs().Attrs(), h.Children...)
	return string(tag)
}

type Tbody struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

type Td struct {
	attributes.AttrHtmlGlobal
	Colspan  int           `json:"colspan,omitempty"` // 跨越的列数
	Rowspan  int           `json:"rowspan,omitempty"` // 跨越的行数
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (t Td) Attrs() attributes.Attrs {
	attrs := t.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Colspan_(t.Colspan),
		attributes.Rowspan_(t.Rowspan),
	)
	return attrs
}

func (t Td) Html() (html string) {
	tag := htmlgo.Td(t.Attrs().Attrs(), t.Children...)
	return string(tag)
}

type Textarea struct {
	attributes.AttrHtmlGlobal
	Name        string `json:"name,omitempty"`        // 输入字段名称
	Rows        int    `json:"rows,omitempty"`        // 行数
	Cols        int    `json:"cols,omitempty"`        // 列数
	Placeholder string `json:"placeholder,omitempty"` // 提示文字
	Required    bool   `json:"required,omitempty"`    // 是否为必填字段
	BasicComponent
}

func (i *Textarea) Component() ComponentI {
	return i
}
func (i *Textarea) Tag() string {
	return "textarea"
}
func (i *Textarea) IsNil() bool {
	return false
}

func (t Textarea) Attrs() attributes.Attrs {
	attrs := t.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Name_(t.Name),
		attributes.Rows_(t.Rows),
		attributes.Cols_(t.Cols),
		attributes.Placeholder_(t.Placeholder),
		attributes.Required_(t.Required),
	)
	return attrs
}
func (t Textarea) Html() (html htmlgo.HTML) {
	tag := htmlgo.Textarea(t.Attrs().Attrs())
	return tag
}

type Tfoot struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (t Tfoot) Attrs() attributes.Attrs {
	return t.AttrHtmlGlobal.Attrs()
}
func (t Tfoot) Html() (html string) {
	tag := htmlgo.Tfoot(t.Attrs().Attrs(), t.Children...)
	return string(tag)
}

type Th struct {
	attributes.AttrHtmlGlobal
	Colspan  int           `json:"colspan,omitempty"` // 跨越的列数
	Rowspan  int           `json:"rowspan,omitempty"` // 跨越的行数
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (t Th) Attrs() attributes.Attrs {
	attrs := t.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Colspan_(t.Colspan),
		attributes.Rowspan_(t.Rowspan),
	)
	return attrs
}
func (t Th) Html() (html string) {
	tag := htmlgo.Th(t.Attrs().Attrs(), t.Children...)
	return string(tag)
}

type Thead struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (t Thead) Attrs() attributes.Attrs {
	attrs := t.AttrHtmlGlobal.Attrs()
	return attrs
}
func (t Thead) Html() (html string) {
	tag := htmlgo.Thead(t.Attrs().Attrs(), t.Children...)
	return string(tag)
}

type Time struct {
	Datetime string `json:"datetime,omitempty"` // 时间的 ISO 格式
	Class    string `json:"class,omitempty"`    // CSS 类名
}

type Title struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (t Title) Attrs() attributes.Attrs {
	return t.AttrHtmlGlobal.Attrs()
}
func (t Title) Html() (html string) {
	tag := htmlgo.Title(t.Attrs().Attrs(), t.Children...)
	return string(tag)
}

type Tr struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (t Tr) Attrs() attributes.Attrs {
	return t.AttrHtmlGlobal.Attrs()
}
func (t Tr) Html() (html string) {
	tag := htmlgo.Title(t.Attrs().Attrs(), t.Children...)
	return string(tag)
}

type Track struct {
	Src     string `json:"src,omitempty"`     // 字幕文件的 URL
	Kind    string `json:"kind,omitempty"`    // 字幕类型
	Srclang string `json:"srclang,omitempty"` // 字幕语言
	Label   string `json:"label,omitempty"`   // 字幕的标签
	Class   string `json:"class,omitempty"`   // CSS 类名
}

type Tt struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type U struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Ul struct {
	attributes.AttrHtmlGlobal
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (u Ul) Attrs() attributes.Attrs {
	return u.AttrHtmlGlobal.Attrs()
}
func (u Ul) Html() (html string) {
	tag := htmlgo.Ul(u.Attrs().Attrs(), u.Children...)
	return string(tag)
}

type Var struct {
	Id    string           `json:"id,omitempty"`    // 元素的唯一标识符
	Class attributes.Class `json:"class,omitempty"` // CSS 类名
}

type Video struct {
	attributes.AttrHtmlGlobal
	Src      string        `json:"src,omitempty"`      // 视频文件 URL
	Controls bool          `json:"controls,omitempty"` // 是否显示控件
	Autoplay bool          `json:"autoplay,omitempty"` // 是否自动播放
	Loop     bool          `json:"loop,omitempty"`     // 是否循环播放
	Muted    bool          `json:"muted,omitempty"`    // 是否静音
	Children []htmlgo.HTML `json:"children,omitempty"`
}

func (v Video) Attrs() attributes.Attrs {
	attrs := v.AttrHtmlGlobal.Attrs()
	attrs.AddRef(
		attributes.Src_(v.Src),
		attributes.Controls_(v.Controls),
		attributes.Autoplay_(v.Autoplay),
		attributes.Loop_(v.Loop),
		attributes.Muted_(v.Muted),
	)
	return attrs
}
func (v Video) Html() (html string) {
	tag := htmlgo.Video(v.Attrs().Attrs(), v.Children...)
	return string(tag)
}

type Wbr struct {
	attributes.AttrHtmlGlobal
}
