package gen10

import (
	"image/color"
	"math"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-10
// Polar Coordinates

const (
	InitialW = 640
	InitialH = 480
	R        = 0.4
	Step     = 0.05
	LW       = 16
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}
var red = color.RGBA{150, 0, 0, 255}

type Gen10 struct {
	W float32
	H float32
	P float32
}

func New() *Gen10 {
	l := &Gen10{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen10) IsLevel(nl string) bool {
	return nl == "10"
}

func (l *Gen10) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen10) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	cX := l.W / 2
	cY := l.H / 2

	r := l.W * R

	if l.W > l.H {
		r = l.H * R
	}

	p := &vector.Path{}
	p1 := &vector.Path{}

	p.Arc(cX, cY, r, l.P, l.P-Step, 1)
	p1.Arc(cX, cY, r, l.P+math.Pi, l.P+math.Pi-Step, 1)

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)
	c1 := ebiten.ColorScale{}
	c1.ScaleWithColor(red)

	pathOpts := &vector.DrawPathOptions{ColorScale: c}
	pathOpts1 := &vector.DrawPathOptions{ColorScale: c1}
	strokeOpts := &vector.StrokeOptions{Width: LW}

	vector.StrokePath(screen, p, strokeOpts, pathOpts)
	vector.StrokePath(screen, p1, strokeOpts, pathOpts1)
}

func (l *Gen10) Update() error {
	l.P += Step

	return nil
}

func (l *Gen10) Layout(w, h float32) {
	l.W = w
	l.H = h
}
