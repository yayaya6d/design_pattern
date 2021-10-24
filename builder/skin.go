package builder

type Color string

const (
	Blaxk Color = "black"
	White Color = "white"
	Red   Color = "red"
)

type Skin struct {
	color    Color
	hardness int
}
