package daisyui

type Background string

func (b Background) Combine(prefix string) (className string) {
	return CombineClassName(prefix, string(b))
}

const (
	Background_None Background = "ghost"
)

type Color string

func (c Color) Combine(prefix string) (className string) {
	return CombineClassName(prefix, string(c))
}

const (
	Color_primary   Color = "primary"
	Color_secondary Color = "secondary"
	Color_accent    Color = "accent"
)

type ResponseStatus string

func (r ResponseStatus) Combine(prefix string) (className string) {
	return CombineClassName(prefix, string(r))
}

const (
	Response_status_success ResponseStatus = "success"
	Response_status_error   ResponseStatus = "error"
	Response_status_info    ResponseStatus = "info"
	Response_status_warning ResponseStatus = "warning"
)

type Size string

func (s Size) Combine(prefix string) (className string) {
	return CombineClassName(prefix, string(s))

}

const (
	Size_lg Size = "lg"
	Size_md Size = "md"
	Size_sm Size = "sm"
	Size_xs Size = "xs"
)

type Border string

func (b Border) Combine(prefix string) (className string) {
	return CombineClassName(prefix, string(b))
}

const (
	Border_bordered Border = "bordered"
)
