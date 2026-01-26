package gen17

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-17
// Wallpaper group

const (
	InitialW = 640
	InitialH = 480
)

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen17 struct {
	W float32
	H float32
}

func New() *Gen17 {
	l := &Gen17{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen17) IsLevel(nl string) bool {
	return nl == "17"
}

func (l *Gen17) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen17) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	p := &vector.Path{}

	gap := float32(32)

	sg := gap / 3

	for i := range int(l.H/gap) + 1 {
		y := float32(i) * gap
		for j := range int(l.W/gap) + 1 {
			x := float32(j) * gap

			if (i+j)%2 == 0 {
				p.MoveTo(x, y)
				p.LineTo(x, y+gap)
				p.MoveTo(x+sg, y)
				p.LineTo(x+sg, y+gap)
				p.MoveTo(x+sg*2, y)
				p.LineTo(x+sg*2, y+gap)
				continue
			}

			p.MoveTo(x, y)
			p.LineTo(x+gap, y)
			p.MoveTo(x, y+sg)
			p.LineTo(x+gap, y+sg)
			p.MoveTo(x, y+sg*2)
			p.LineTo(x+gap, y+sg*2)

		}
	}

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)
	strokeOpts := &vector.StrokeOptions{Width: 2}
	drawOpts := &vector.DrawPathOptions{ColorScale: c}
	vector.StrokePath(screen, p, strokeOpts, drawOpts)
}

func (l *Gen17) Update() error {
	return nil
}

func (l *Gen17) Layout(w, h float32) {
	l.W = w
	l.H = h
}
