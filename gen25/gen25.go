package gen25

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-25
// Organic geometry

const (
	InitialW = 640
	InitialH = 480
	D        = 128
)

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen25 struct {
	W float32
	H float32
	T int
	D int
}

func New() *Gen25 {
	l := &Gen25{
		W: InitialW,
		H: InitialH,
		D: 1,
	}

	return l
}

func (l *Gen25) IsLevel(nl string) bool {
	return nl == "25"
}

func (l *Gen25) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen25) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	xc := l.W * 0.5
	yc := l.H * 0.5

	ln := float32(l.T)

	p := &vector.Path{}

	p.MoveTo(xc, yc-ln*0.5)
	p.LineTo(xc+ln*0.5, yc+ln*0.5)
	p.LineTo(xc-ln*0.5, yc+ln*0.5)
	p.Close()

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)
	strokeOpts := &vector.StrokeOptions{Width: 2}
	drawOpts := &vector.DrawPathOptions{ColorScale: c}

	vector.StrokePath(screen, p, strokeOpts, drawOpts)
}

func (l *Gen25) Update() error {
	l.T += l.D

	if l.T > D || l.T < 1 {
		l.D *= -1
	}

	return nil
}

func (l *Gen25) Layout(w, h float32) {
	l.W = w
	l.H = h
}
