package gen12

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-12
// Boxes only

const (
	InitialW = 640
	InitialH = 480
	S        = 0.05
	BW       = 2
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen12 struct {
	W   float32
	H   float32
	Img *ebiten.Image
}

func New() *Gen12 {
	l := &Gen12{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen12) IsLevel(nl string) bool {
	return nl == "12"
}

func (l *Gen12) NextLevel() string {
	if input.IsPressed() || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	return ""
}

func (l *Gen12) Draw(screen *ebiten.Image) {
	if l.Img == nil {
		l.Img = ebiten.NewImage(int(l.W), int(l.H))
	}

	g := ebiten.GeoM{}

	g.Scale(0.99, 0.99)

	opts := &ebiten.DrawImageOptions{GeoM: g}

	screen.DrawImage(l.Img, opts)

	l.Img = nil
	l.Img = ebiten.NewImageFromImage(screen)

	s := l.W * S
	sh := l.H * S

	if sh < s {
		s = sh
	}

	// s *= rand.Float32()

	x := rand.Float32() * l.W
	y := rand.Float32() * l.H

	c := bg

	if rand.Float32() > 0.5 {
		c = fg
	}

	vector.FillRect(l.Img, x, y, s, s, c, false)
	vector.StrokeRect(l.Img, x, y, s, s, BW, bg, false)

	if 1 != 2 {
		return
	}

	for i := range int(l.W/s) + 1 {
		x := float32(i) * s
		for j := range int(l.H/s) + 1 {
			y := float32(j) * s

			c := bg

			r := rand.Float32()

			if r > 0.5 {
				continue
			}
			if r > 0.25 {
				c = fg
			}

			vector.FillRect(l.Img, x, y, s, s, c, false)
		}
	}
}

func (l *Gen12) Update() error {
	return nil
}

func (l *Gen12) Layout(w, h float32) {
	l.W = w
	l.H = h
}
