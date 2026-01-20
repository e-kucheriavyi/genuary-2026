package gen15

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-15
// Invisible object with shadow

const (
	InitialW = 640
	InitialH = 480
	Size     = 0.2
)

var (
	bg   = color.RGBA{0, 0, 0, 255}
	pale = color.RGBA{0, 100, 0, 255}
	fg   = color.RGBA{0, 150, 0, 255}
)

type Point struct {
	X float32
	Y float32
}

type Gen15 struct {
	W      float32
	H      float32
	X      float32
	Y      float32
	Points []*Point
}

func New() *Gen15 {
	l := &Gen15{
		W: InitialW,
		H: InitialH,
		Points: []*Point{
			{X: 100, Y: 100},
			{X: 200, Y: 100},
			{X: 200, Y: 200},
			{X: 100, Y: 200},
			{X: 100, Y: 100},
		},
	}

	return l
}

func (l *Gen15) IsLevel(nl string) bool {
	return nl == "15"
}

func (l *Gen15) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	if input.IsPressed() {
		x, y := input.CursorPosition()

		p := float32(text.LetterWidth * 4 * 2)

		if x <= p && y <= p {
			return "menu"
		}
	}
	return ""
}

func (l *Gen15) Draw(screen *ebiten.Image) {
	p1 := &vector.Path{}

	p1.MoveTo(l.X, l.Y)

	m := float32(10)

	for i, pt := range l.Points {
		if i == 0 {
			p1.MoveTo(pt.X, pt.Y)
		}
		p1.LineTo(pt.X, pt.Y)
		x := (pt.X - l.X) * m
		y := (pt.Y - l.Y) * m
		p1.LineTo(x, y)
	}

	p1.Close()

	c := ebiten.ColorScale{}
	c.ScaleWithColor(pale)

	pathOpts := &vector.DrawPathOptions{ColorScale: c, AntiAlias: true}
	fillOpts := &vector.FillOptions{}

	vector.FillPath(screen, p1, fillOpts, pathOpts)

	text.DrawLetter(screen, 'x', 0, 0, float32(4), fg)
}

func (l *Gen15) Update() error {
	x, y := input.CursorPosition()

	l.X = x
	l.Y = y

	return nil
}

func (l *Gen15) Rebuild() {
	cx := l.W / 2
	cy := l.H / 2

	s := l.W * Size
	sh := s / 2

	l.Points[0].X = cx - sh
	l.Points[0].Y = cy - sh

	l.Points[1].X = cx + sh
	l.Points[1].Y = cy - sh

	l.Points[2].X = cx + sh
	l.Points[2].Y = cy + sh

	l.Points[3].X = cx - sh
	l.Points[3].Y = cy + sh

	l.Points[4].X = cx - sh
	l.Points[4].Y = cy - sh
}

func (l *Gen15) Layout(w, h float32) {
	if l.W != w || l.H != h {
		l.Rebuild()
	}
	l.W = w
	l.H = h
}
