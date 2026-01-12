package gen08

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-08
// City

const (
	InitialW = 640
	InitialH = 480
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen08 struct {
	W float32
	H float32
}

func New() *Gen08 {
	l := &Gen08{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen08) IsLevel(nl string) bool {
	return nl == "08"
}

func (l *Gen08) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen08) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	vector.FillRect(screen, 0, 0, 10, 10, fg, false)
}

func (l *Gen08) Update() error {
	return nil
}

func (l *Gen08) Layout(w, h float32) {
	l.W = w
	l.H = h
}
