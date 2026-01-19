package gen19

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-19
// City

const (
	InitialW = 640
	InitialH = 480
	Size     = 16
	D        = Size * Size
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen19 struct {
	W   float32
	H   float32
	T   int
	Img *ebiten.Image
}

func New() *Gen19 {
	l := &Gen19{
		W:   InitialW,
		H:   InitialH,
		Img: ebiten.NewImage(Size, Size),
	}

	return l
}

func (l *Gen19) IsLevel(nl string) bool {
	return nl == "19"
}

func (l *Gen19) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen19) Draw(screen *ebiten.Image) {
	if l.Img == nil {
		return
	}

	t := 0

	for y := range Size {
		for x := range Size {
			t += 1

			if t < l.T {
				continue
			}

			c := color.RGBA{
				uint8(rand.Float32() * 255),
				uint8(rand.Float32() * 255),
				uint8(rand.Float32() * 255),
				uint8(rand.Float32() * 255),
			}

			vector.FillRect(l.Img, float32(x), float32(y), 1, 1, c, false)
			break
		}

		if t >= l.T {
			break
		}
	}

	op := &ebiten.DrawImageOptions{}

	s := l.W

	if l.W > l.H {
		s = l.H
	}

	op.GeoM.Scale(float64(s/Size), float64(s/Size))

	screen.DrawImage(l.Img, op)
}

func (l *Gen19) Update() error {
	l.T += 1

	if l.T >= D {
		l.Img = nil
		l.Img = ebiten.NewImage(Size, Size)
		l.T = 0
	}

	return nil
}

func (l *Gen19) Layout(w, h float32) {
	l.W = w
	l.H = h
}
