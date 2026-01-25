package gen24

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2024-01-24
// Perfectionist hell

const (
	InitialW = 640
	InitialH = 480
)

var bg = color.RGBA{0, 0, 0, 245}
var fg = color.RGBA{0, 150, 0, 245}

type Gen24 struct {
	W float32
	H float32
	X float32
	Y float32
}

func New() *Gen24 {
	l := &Gen24{
		W: InitialW,
		H: InitialH,
		X: 0,
		Y: 0,
	}

	return l
}

func (l *Gen24) IsLevel(nl string) bool {
	return nl == "24"
}

func (l *Gen24) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen24) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	vector.FillRect(screen, l.X, l.Y, 2, 2, fg, false)
}

func (l *Gen24) Update() error {
	return nil
}

func (l *Gen24) Layout(w, h float32) {
	if w != l.W || h != l.H {
		l.X = rand.Float32() * w
		l.Y = rand.Float32() * h
	}

	l.W = w
	l.H = h
}
