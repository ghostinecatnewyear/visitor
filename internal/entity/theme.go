package entity

type Color string

const (
	ColorRed    Color = "red"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorYellow Color = "yellow"
)

func (c Color) IsValid() bool {
	return c == ColorRed || c == ColorGreen || c == ColorBlue || c == ColorYellow
}

type Theme struct {
	Color Color `json:"color"`
}

func (t *Theme) IsValid() bool {
	return t.Color.IsValid()
}

type UpdateTheme struct {
	Color *Color `json:"color"`
}

func (ut *UpdateTheme) Apply(t Theme) *Theme {
	if ut.Color != nil {
		t.Color = *ut.Color
	}

	return &t
}
