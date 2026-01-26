package gen16

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-16
// Order and disorder

const (
	InitialW = 640
	InitialH = 480
	D        = 4
	S        = 8
)

var bg = color.RGBA{0, 0, 0, 255}

type Gen16 struct {
	W   float32
	H   float32
	T   int
	Img *ebiten.Image
}

func New() *Gen16 {
	l := &Gen16{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen16) IsLevel(nl string) bool {
	return nl == "16"
}

func (l *Gen16) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen16) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	if l.Img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	if l.T > 0 {
		screen.DrawImage(l.Img, op)
		return
	}

	y := float32(0)

	for {
		if y >= l.H {
			break
		}
		x := l.W / 2
		for {
			if x >= l.W {
				break
			}

			sh := uint8(rand.Float32() * 255)
			c := color.RGBA{sh, sh, sh, 255}
			vector.FillRect(l.Img, x, y, S, S, c, false)
			x += S
		}

		y += S
	}

	screen.DrawImage(l.Img, op)
}

func (l *Gen16) Update() error {
	l.T += 1

	if l.T > D {
		l.Img = nil
		l.Img = ebiten.NewImage(int(l.W), int(l.H))
		l.T = 0
	}
	return nil
}

func (l *Gen16) Layout(w, h float32) {
	l.W = w
	l.H = h
}
