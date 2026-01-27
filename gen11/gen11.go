package gen11

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
)

// 2025-01-11
// Quine

const (
	InitialW = 640
	InitialH = 480
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen11 struct {
	W float32
	H float32
}

func New() *Gen11 {
	l := &Gen11{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen11) IsLevel(nl string) bool {
	return nl == "11"
}

func (l *Gen11) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen11) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	gap := float32(64)
	s := float32(6)

	for i := range int(l.H/gap)+1 {
		y := float32(i) * gap + gap*0.1
		for j := range int(l.W/gap)+1 {
			x := float32(j) * gap + gap*0.1
			l := '0'
			if rand.Float32() > 0.5 {
				l = '1'
			}
			text.DrawLetter(screen, l, x, y, s, fg)
		}
	}
}

func (l *Gen11) Update() error {
	return nil
}

func (l *Gen11) Layout(w, h float32) {
	l.W = w
	l.H = h
}
