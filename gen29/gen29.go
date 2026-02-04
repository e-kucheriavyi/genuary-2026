package gen29

import (
	"image/color"
	"math"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-29
// Genetic evolution

const (
	InitialW = 640
	InitialH = 480
	MaxT     = 8
	MinN     = 3
	MaxN     = 32
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen29 struct {
	W  float32
	H  float32
	T  int
	N  int
	DN int
	R  float32
}

func New() *Gen29 {
	l := &Gen29{
		W:  InitialW,
		H:  InitialH,
		N:  3,
		DN: 1,
		R:  0.5,
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
	r := float32(math.Min(float64(l.H), float64(l.W))) * 0.45
	x := l.W * 0.5
	y := l.H * 0.5
	lw := float32(4)

	DrawNgon(screen, x, y, r, l.R, lw, l.N, fg)
}

func DrawNgon(
	screen *ebiten.Image,
	x, y, radius, rot, lw float32,
	n int,
	stroke color.Color,
) {
	if n < MinN {
		return
	}

	rotation := float32(-math.Pi * rot)
	step := float32(math.Pi*2.0) / float32(n)

	p := &vector.Path{}

	for i := range n {
		angle := rotation + step*float32(i)
		px := x + float32(math.Cos(float64(angle)))*radius
		py := y + float32(math.Sin(float64(angle)))*radius

		if i == 0 {
			p.MoveTo(px, py)
			continue
		}

		p.LineTo(px, py)
	}

	p.Close()

	c := ebiten.ColorScale{}
	c.ScaleWithColor(stroke)
	strokeOpts := &vector.StrokeOptions{Width: lw}
	pathOpts := &vector.DrawPathOptions{ColorScale: c}
	vector.StrokePath(screen, p, strokeOpts, pathOpts)
}

func (l *Gen29) Update() error {
	l.T += 1

	if l.T >= MaxT {
		l.T = 0
		l.N += l.DN

		if l.N >= MaxN || l.N <= MinN {
			l.DN *= -1
		}
	}

	l.R += 0.01

	return nil
}

func (l *Gen29) Layout(w, h float32) {
	l.W = w
	l.H = h
}
