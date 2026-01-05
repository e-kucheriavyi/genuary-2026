package gen05

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-05
// Write "genuary" without font

const (
	InitialW = 640
	InitialH = 480
	T        = 255
)

var bg = color.RGBA{0, 100, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Point struct {
	X float32
	Y float32
	M bool
}

type Gen05 struct {
	W      float32
	H      float32
	X      float32
	Y      float32
	D      int
	I      int
	T      float32
	Points []*Point
}

func New() *Gen05 {
	l := &Gen05{
		W:      InitialW,
		H:      InitialH,
		Points: []*Point{},
		D:      1,
	}

	return l
}

func (l *Gen05) IsLevel(nl string) bool {
	return nl == "05"
}

func (l *Gen05) NextLevel() string {
	if input.IsPressed() || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	return ""
}

func (l *Gen05) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, l.W, l.H, bg, false)

	p := &vector.Path{}

	for i, pt := range l.Points {
		if pt.M {
			p.MoveTo(pt.X, pt.Y)
			continue
		}

		if i == l.I || true {
			t := l.T / T

			prev := l.Points[i-1]
			p.LineTo(
				utils.Lerp(prev.X, pt.X, t),
				utils.Lerp(prev.Y, pt.Y, t),
			)
		}
	}

	c := ebiten.ColorScale{}

	c.ScaleWithColor(fg)

	pathOpts := &vector.DrawPathOptions{
		AntiAlias:  true,
		ColorScale: c,
	}
	strokeOpts := &vector.StrokeOptions{
		Width:    6,
		LineJoin: 2,
		LineCap:  1,
	}

	vector.StrokePath(screen, p, strokeOpts, pathOpts)
}

func (l *Gen05) Update() error {
	l.T += float32(l.D)

	if l.T >= T || l.T <= 0 {
		l.D *= -1
	}

	return nil
}

func (l *Gen05) MoveTo(x, y float32) {
	l.X += x
	l.Y += y
	l.Points = append(l.Points, &Point{l.X, l.Y, true})
}

func (l *Gen05) LineTo(x, y float32) {
	l.X += x
	l.Y += y
	l.Points = append(l.Points, &Point{l.X, l.Y, false})
	l.Points = append(l.Points, &Point{l.X, l.Y, true})
}

func (l *Gen05) Layout(w, h float32) {
	l.W = w
	l.H = h

	u := w / 12

	l.X = u * 3
	l.Y = l.H / 2

	l.Points = []*Point{}

	// G
	l.MoveTo(0, 0)
	l.LineTo(-u, -u)
	l.LineTo(-u, +u)
	l.LineTo(+u, +u)
	l.LineTo(+u, -u)
	l.LineTo(-u, 0)

	// E
	l.MoveTo(+u*2.2, -u)
	l.LineTo(-u, 0)
	l.LineTo(0, +u*2)
	l.LineTo(u, 0)
	l.MoveTo(-u, -u)
	l.LineTo(u*0.8, 0)

	// N
	l.MoveTo(+u*0.4, +u)
	l.LineTo(0, -u*2)
	l.LineTo(u, +u*2)
	l.LineTo(0, -u*2)

	// U
	l.MoveTo(+u*0.2, 0)
	l.LineTo(0, +u*1.8)
	l.LineTo(+u*0.5, +u*0.2)
	l.LineTo(+u*0.5, -u*0.2)
	l.LineTo(0, -u*1.8)

	// A
	l.MoveTo(+u*0.2, +u*2)
	l.LineTo(u*0.5, -u*2)
	l.LineTo(+u*0.5, +u*2)
	l.MoveTo(0, -u)
	l.LineTo(-u, 0)

	// R
	l.MoveTo(+u*1.2, +u)
	l.LineTo(0, -u*2)
	l.LineTo(+u*0.8, 0)
	l.LineTo(+u*0.2, u*0.5)
	l.LineTo(-u*0.2, u*0.5)
	l.LineTo(-u*0.7, 0)
	l.LineTo(+u*0.7, u)

	// Y
	l.MoveTo(+u*0.2, -u*2)
	l.LineTo(+u*0.5, +u*0.6)
	l.MoveTo(+u*0.5, -u*0.6)
	l.LineTo(-u*0.5, +u*0.6)
	l.LineTo(0, +u*1.4)
}
