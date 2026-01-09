package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Lerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

func DrawBitmap(
	screen *ebiten.Image,
	m *[]byte,
	x, y, s float32,
	l int,
	c color.Color,
) {
	i := 0

	p := &vector.Path{}

	for j := range l {
		for k := range l {
			pX := x + float32(k)*s
			pY := y + float32(j)*s

			if (*m)[i] == 1 {
				p.MoveTo(pX, pY)
				p.LineTo(pX+s, pY)
				p.LineTo(pX+s, pY+s)
				p.LineTo(pX, pY+s)
				p.Close()
			}

			i++
		}
	}

	col := ebiten.ColorScale{}
	col.ScaleWithColor(c)

	fillOpts := &vector.FillOptions{}
	pathOpts := &vector.DrawPathOptions{ColorScale: col}

	vector.FillPath(screen, p, fillOpts, pathOpts)
}
