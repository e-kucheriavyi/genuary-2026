package gen27

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-27
// Lifeform

const (
	InitialW = 640
	InitialH = 480
	Ln       = 100
	W        = 16
	D        = 160
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen27 struct {
	W float32
	H float32
	X float32
	T float32
	D float32
	B float32
}

func New() *Gen27 {
	l := &Gen27{
		W: InitialW,
		H: InitialH,
		D: 1,
		X: Ln * 1.2,
	}

	return l
}

func (l *Gen27) IsLevel(nl string) bool {
	return nl == "27"
}

func (l *Gen27) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen27) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	p := &vector.Path{}

	yc := l.H / 2

	a := (1 - (l.B * 2)) / 3

	p.MoveTo(l.X, yc)
	p.LineTo(l.X-Ln*a, yc)

	p.LineTo(l.X-Ln*a, yc-Ln*l.B)
	p.LineTo(l.X-(Ln*a)*2, yc-Ln*l.B)
	p.LineTo(l.X-(Ln*a)*2, yc)
	p.LineTo(l.X-(Ln*a)*3, yc)

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)

	strokeOpts := &vector.StrokeOptions{Width: W, LineCap: 1, LineJoin: 2}
	drawOpts := &vector.DrawPathOptions{ColorScale: c}

	vector.StrokePath(screen, p, strokeOpts, drawOpts)

	vector.FillRect(screen, l.X+6, yc-3, 3, 4, bg, false)
	vector.FillRect(screen, l.X, yc-3, 4, 4, bg, false)
}

func (l *Gen27) Update() error {
	l.T += l.D

	if l.T >= D || l.T <= 0 {
		l.D *= -1
	}

	if l.D > 0 {
		l.B += 0.01
	} else {
		l.B -= 0.01
		if l.B > 0 {
			l.X += 0.02 * Ln
		}
	}

	if l.B >= 0.4 {
		l.B = 0.4
	}
	if l.B < 0 {
		l.B = 0
	}

	if l.X > l.W+Ln*1.1 {
		l.X = 0 - Ln*0.1
	}

	return nil
}

func (l *Gen27) Layout(w, h float32) {
	l.W = w
	l.H = h
}
