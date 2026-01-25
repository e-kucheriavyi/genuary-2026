package gen08

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-08
// City

const (
	InitialW = 640
	InitialH = 480
	D        = 64
)

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen08 struct {
	W   float32
	H   float32
	T   int
	Img *ebiten.Image
}

func New() *Gen08 {
	l := &Gen08{
		W:   InitialW,
		H:   InitialH,
		Img: ebiten.NewImage(InitialW, InitialH),
	}

	return l
}

func (l *Gen08) IsLevel(nl string) bool {
	return nl == "08"
}

func (l *Gen08) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen08) Draw(screen *ebiten.Image) {
	if l.Img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	if l.T != 0 {
		screen.DrawImage(l.Img, op)
		return
	}

	gap := float32(16)

	y := gap

	step := l.W * 0.1

	for {
		x := gap
		h := step * rand.Float32()

		if y+h >= l.H-gap {
			h = l.H - y - gap
		}

		for {
			w := step * rand.Float32()

			if x+w >= l.W-gap {
				w = l.W - x - gap
			}

			vector.StrokeRect(l.Img, x, y, w, h, 2, fg, false)

			x += w + gap

			if x >= l.W {
				break
			}
		}

		y += h + gap

		if y >= l.H {
			break
		}
	}

	screen.DrawImage(l.Img, op)
}

func (l *Gen08) Update() error {
	l.T += 1

	if l.T >= D {
		l.T = 0
		l.Img = nil
		l.Img = ebiten.NewImage(int(l.W), int(l.H))
	}

	return nil
}

func (l *Gen08) Layout(w, h float32) {
	l.W = w
	l.H = h
}
