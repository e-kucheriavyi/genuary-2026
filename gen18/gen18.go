package gen18

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-18
// Unexpected path

const (
	InitialW = 640
	InitialH = 480
)

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen18 struct {
	W float32
	H float32
	X float32
	Y float32
}

func New() *Gen18 {
	l := &Gen18{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen18) IsLevel(nl string) bool {
	return nl == "18"
}

func (l *Gen18) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen18) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	p := &vector.Path{}

	l.X = 0
	l.Y = 0

	l.MoveTo(p, 0, 0)

	m := 0

	ln := float32(32)

	for range 64 {
		switch m {
		case 0:
			l.LineTo(p, 0, ln)
		case 1:
			l.LineTo(p, ln, 0)
		case 2:
			l.LineTo(p, -ln, -ln)
		}
		m += 1

		if m >= 3 {
			ln *= 1.25
			m = 0
		}
	}

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)
	drawOpts := &vector.DrawPathOptions{ColorScale: c}
	strokeOpts := &vector.StrokeOptions{Width: 2}

	vector.StrokePath(screen, p, strokeOpts, drawOpts)
}

func (l *Gen18) MoveTo(p *vector.Path, x, y float32) {
	l.X += x
	l.Y += y

	p.LineTo(l.X, l.Y)
}

func (l *Gen18) LineTo(p *vector.Path, x, y float32) {
	l.X += x
	l.Y += y

	p.LineTo(l.X, l.Y)
}

func (l *Gen18) Update() error {
	return nil
}

func (l *Gen18) Layout(w, h float32) {
	l.W = w
	l.H = h
}
