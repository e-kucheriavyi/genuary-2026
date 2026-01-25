package gen23

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-23
// Transparency

const (
	InitialW = 640
	InitialH = 480
)

var (
	bg = color.RGBA{255, 255, 255, 255}
	red = color.RGBA{255, 0, 0, 100}
	green = color.RGBA{0, 255, 0, 100}
	blue = color.RGBA{0, 255, 0, 100}
)

type Gen23 struct {
	W float32
	H float32
	R uint8
	G uint8
	B uint8
}

func New() *Gen23 {
	l := &Gen23{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen23) IsLevel(nl string) bool {
	return nl == "23"
}

func (l *Gen23) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen23) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	r := color.RGBA{255, 0, 0, l.R}
	g := color.RGBA{0, 255, 0, l.G}
	b := color.RGBA{0, 0, 255, l.B}

	vector.FillRect(screen, 0, 0, l.W * 0.40, l.H, r, false)
	vector.FillRect(screen, l.W * 0.33, 0, l.W * 0.40, l.H, g, false)
	vector.FillRect(screen, l.W * 0.66, 0, l.W * 0.40, l.H, b, false)
}

func (l *Gen23) Update() error {
	l.R += 1
	l.G += 1
	l.B += 1

	return nil
}

func (l *Gen23) Layout(w, h float32) {
	l.W = w
	l.H = h
}
