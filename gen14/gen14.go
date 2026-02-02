package gen14

import (
	"image/color"
	"math"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-14
// Everything fits perfectly

const (
	InitialW = 640
	InitialH = 480
)

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen14 struct {
	W float32
	H float32
}

func New() *Gen14 {
	l := &Gen14{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen14) IsLevel(nl string) bool {
	return nl == "14"
}

func (l *Gen14) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen14) Draw(screen *ebiten.Image) {
	screen.Fill(fg)

	vector.FillRect(screen, 100, 100, 100, 100, bg, false)
	vector.FillRect(screen, 200, 300, 50, 100, bg, false)
	vector.FillCircle(screen, 100, 350, 50, bg, true)

	// triangle
	p := &vector.Path{}
	p.MoveTo(350, 400)
	p.LineTo(550, 400)
	p.LineTo(450, 300)
	p.Close()
	c := ebiten.ColorScale{}
	c.ScaleWithColor(bg)
	fillOpts := &vector.FillOptions{}
	pathOpts := &vector.DrawPathOptions{ColorScale: c, AntiAlias: true}

	vector.FillPath(screen, p, fillOpts, pathOpts)

	// semicircle
	p1 := &vector.Path{}
	p1.Arc(400, 200, 100, 0, math.Pi, 1)
	p1.Close()

	vector.FillPath(screen, p1, fillOpts, pathOpts)

	// arc
	vector.FillRect(screen, 200, 450, 200, 100, bg, false)
	vector.FillCircle(screen, 300, 550, 50, fg, true)
}

func (l *Gen14) Update() error {
	return nil
}

func (l *Gen14) Layout(w, h float32) {
	l.W = w
	l.H = h
}
