package gen30

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-30
// Feature, not a bug

const (
	InitialW = 640
	InitialH = 480
	S        = 100
)

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen30 struct {
	W   float32
	H   float32
	X   float32
	Y   float32
	Img *ebiten.Image
}

func New() *Gen30 {
	l := &Gen30{
		W:   InitialW,
		H:   InitialH,
		X:   InitialW*0.5 - S*0.5,
		Y:   InitialH*0.5 - S*0.5,
		Img: ebiten.NewImage(InitialW, InitialH),
	}

	return l
}

func (l *Gen30) IsLevel(nl string) bool {
	return nl == "30"
}

func (l *Gen30) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen30) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	if l.Img == nil {
		return
	}

	vector.FillRect(l.Img, l.X, l.Y, S, S, bg, false)
	vector.StrokeRect(l.Img, l.X, l.Y, S, S, 2, fg, false)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(l.Img, op)
}

func (l *Gen30) Update() error {
	if rand.Float32() > 0.5 {
		l.X += 1
	} else {
		l.X -= 1
	}
	if rand.Float32() > 0.5 {
		l.Y += 1
	} else {
		l.Y -= 1
	}

	return nil
}

func (l *Gen30) Layout(w, h float32) {
	if l.W != w || l.H != h {
		l.Img = nil
		l.Img = ebiten.NewImage(int(w), int(h))
	}
	l.W = w
	l.H = h
}
