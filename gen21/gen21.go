package gen21

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-21
// Bauhaus

const (
	InitialW = 640
	InitialH = 480
)

var (
	bg     = color.RGBA{0xdd, 0xcb, 0xa6, 0xff}
	blue   = color.RGBA{0x00, 0x6f, 0xb6, 0xff}
	yellow = color.RGBA{0xf3, 0xc3, 0x01, 0xff}
	red    = color.RGBA{0xd8, 0x00, 0x0b, 0xff}
)

type Gen21 struct {
	W float32
	H float32
}

func New() *Gen21 {
	l := &Gen21{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen21) IsLevel(nl string) bool {
	return nl == "21"
}

func (l *Gen21) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen21) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	vector.FillRect(screen, 0.5*l.W, 0.6*l.H, l.W*0.4, l.H*0.2, yellow, false)

	vector.FillCircle(screen, l.W*0.5, l.H*0.4, l.W*0.3, blue, true)

	vector.FillRect(screen, 0, 0.1*l.H, 0.1*l.W, 0.2*l.H, red, false)

	p := &vector.Path{}

	p.MoveTo(0, 0.1*l.H)
	p.LineTo(l.W, 0.1*l.H)

	p.MoveTo(0, 0.3*l.H)
	p.LineTo(l.W, 0.3*l.H)

	p.MoveTo(0, 0.8*l.H)
	p.LineTo(l.W, 0.8*l.H)

	p.MoveTo(0.1*l.W, 0)
	p.LineTo(0.1*l.W, l.H)

	p.MoveTo(0.5*l.W, 0)
	p.LineTo(0.5*l.W, 0.8*l.H)

	p.MoveTo(0.9*l.W, 0.6*l.H)
	p.LineTo(0.9*l.W, l.H)

	p.MoveTo(0.5*l.W, 0.6*l.H)
	p.LineTo(l.W, 0.6*l.H)

	c := ebiten.ColorScale{}
	c.ScaleWithColor(color.Black)

	sOpts := &vector.StrokeOptions{Width: 16}
	pOpts := &vector.DrawPathOptions{ColorScale: c}

	vector.StrokePath(screen, p, sOpts, pOpts)

}

func (l *Gen21) Update() error {
	return nil
}

func (l *Gen21) Layout(w, h float32) {
	l.W = w
	l.H = h
}
