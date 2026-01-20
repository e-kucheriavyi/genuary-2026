package gen20

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-20
// One line

const (
	InitialW = 640
	InitialH = 480
	Step     = 0.1
	D        = 32
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen20 struct {
	W float32
	H float32
	T int
	P *vector.Path
}

func New() *Gen20 {
	l := &Gen20{
		W: InitialW,
		H: InitialH,
		P: &vector.Path{},
	}

	return l
}

func (l *Gen20) IsLevel(nl string) bool {
	return nl == "20"
}

func (l *Gen20) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen20) Draw(screen *ebiten.Image) {
	if l.P == nil {
		return
	}

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)

	strokeOpts := &vector.StrokeOptions{Width: 2}
	pathOpts := &vector.DrawPathOptions{ColorScale: c}

	vector.StrokePath(screen, l.P, strokeOpts, pathOpts)
}

func (l *Gen20) Update() error {
	l.T += 1

	if l.T >= D {
		l.T = 0
		l.Rebuild()
	}

	return nil
}

func (l *Gen20) Rebuild() {
	l.P = nil
	l.P = &vector.Path{}

	x := float32(l.W * 0.5)
	y := float32(l.H * 0.5)

	l.P.MoveTo(x, y)

	xD := float32(1)
	yD := float32(-1)

	for range 40 {
		x += l.W * Step * xD
		y += l.H * Step * yD

		l.P.LineTo(x, y)

		if rand.Float32() > 0.5 {
			xD *= -1
		} else {
			yD *= -1
		}
	}
}

func (l *Gen20) Layout(w, h float32) {
	if l.W != w || l.H != h {
		l.T = 0
		l.Rebuild()
	}

	l.W = w
	l.H = h
}
