package gen22

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-22
// Pen plotter ready

const (
	InitialW = 640
	InitialH = 480
)

var (
	bg   = color.RGBA{255, 255, 255, 255}
	fg   = color.RGBA{0, 150, 0, 255}
	pale = color.RGBA{0, 0, 100, 100}
	red  = color.RGBA{150, 0, 0, 255}
)

type Gen22 struct {
	W float32
	H float32
}

func New() *Gen22 {
	l := &Gen22{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen22) IsLevel(nl string) bool {
	return nl == "22"
}

func (l *Gen22) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen22) Draw(screen *ebiten.Image) {
	screen.Fill(bg)
	p := &vector.Path{}

	gap := float32(32)

	lx := (float32(int(l.W/gap)) - 2) * gap

	for i := range int(l.W/gap) + 1 {
		x := float32(i) * gap
		vector.StrokeLine(screen, x, 0, x, l.H, 1, pale, false)
	}

	for i := range int(l.H/gap) + 1 {
		y := float32(i) * gap
		vector.StrokeLine(screen, 0, y, l.W, y, 1, pale, false)
	}

	x := lx
	y := float32(0)

	m := 0

	vector.StrokeLine(screen, x-gap*0.5, 0, x-gap*0.5, l.H, 1, red, false)

	for i := range int(l.H/gap) + 1 {
		if i != 0 {
			y += gap
		}

		switch m {
		case 0, 2:
			p.MoveTo(x, y)
			p.LineTo(x, y+gap)

			p.MoveTo(x+gap, y)
			p.LineTo(x+gap, y+gap)

			p.MoveTo(x+gap*2, y)
			p.LineTo(x+gap*2, y+gap)
			break
		case 1:
			p.MoveTo(x, y)
			p.LineTo(x+gap, y+gap)

			p.MoveTo(x+gap, y)
			p.LineTo(x+gap*2, y+gap)

			p.MoveTo(x+gap*2, y)
			p.LineTo(x+gap*1.5, y+gap*0.5)

			p.MoveTo(x+gap*0.5, y+gap*0.5)
			p.LineTo(x, y+gap)
			break
		case 3:
			p.MoveTo(x, y)
			p.LineTo(x+gap*0.5, y+gap*0.5)

			p.MoveTo(x+gap*1.5, y+gap*0.5)
			p.LineTo(x+gap*2, y+gap)

			p.MoveTo(x+gap, y)
			p.LineTo(x, y+gap)

			p.MoveTo(x+gap*2, y)
			p.LineTo(x+gap, y+gap)
			break
		}

		m += 1

		if m > 3 {
			m = 0
		}
	}

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)
	strokeOpts := &vector.StrokeOptions{Width: 2}
	drawOpts := &vector.DrawPathOptions{ColorScale: c}

	vector.StrokePath(screen, p, strokeOpts, drawOpts)
}

func (l *Gen22) Update() error {
	return nil
}

func (l *Gen22) Layout(w, h float32) {
	l.W = w
	l.H = h
}
