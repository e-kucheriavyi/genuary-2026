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

	for j := range l {
		for k := range l {
			pX := x + float32(k)*s
			pY := y + float32(j)*s

			if (*m)[i] == 1 {
				vector.FillRect(screen, pX, pY, s, s, c, false)
			}

			i++
		}
	}
}
