package gen29

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-29
// Genetic evolution

const (
	InitialW = 640
	InitialH = 480
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen29 struct {
	W float32
	H float32
}

func New() *Gen29 {
	l := &Gen29{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen29) IsLevel(nl string) bool {
	return nl == "29"
}

func (l *Gen29) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen29) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	vector.FillRect(screen, 0, 0, 10, 10, fg, false)
}

func (l *Gen29) Update() error {
	return nil
}

func (l *Gen29) Layout(w, h float32) {
	l.W = w
	l.H = h
}
