package gen04

import (
	"image/color"
	"unicode/utf8"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-04
// Lowres

const (
	InitialW = 640
	InitialH = 480
	W        = 16
	hH       = 7
	bH       = 9
	D        = 40
)

var bg = color.RGBA{0, 100, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

var head = []byte{
	0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0,
	0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0,
	0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var body = []byte{
	0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
	0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0,
	0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0,
	0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0,
	0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0,
	0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

type Gen04 struct {
	W float32
	H float32
	B bool
	T int
}

func New() *Gen04 {
	l := &Gen04{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen04) IsLevel(nl string) bool {
	return nl == "04"
}

func (l *Gen04) NextLevel() string {
	if input.IsPressed() || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	return ""
}

func (l *Gen04) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, l.W, l.H, bg, false)

	wS := float32(l.W/W) * 0.75
	hS := float32(l.H/W) * 0.75

	s := wS

	if hS < wS {
		s = hS
	}

	x := (l.W / 2) - ((W * s) / 2)
	y := (l.H / 2) - ((bH * s) / 2) + (hH * s / 2)

	drawBitmap(screen, &body, x, y, s, W, bH)

	y -= hH * s * 0.8

	if l.B {
		y += s
	}

	drawBitmap(screen, &head, x, y, s, W, hH)

	if l.B {
		txt := "ква"
		s := float32(4)
		x := (l.W / 2) - (float32(utf8.RuneCountInString(txt)) * (text.LetterWidth * s) / 2)
		text.Write(screen, txt, x, y - text.LetterWidth * s * 2, s, fg)
	}
}

func drawBitmap(
	screen *ebiten.Image,
	m *[]byte,
	x, y, s float32,
	w, h int,
) {
	i := 0

	for j := range h {
		for k := range w {
			pX := x + float32(k)*s
			pY := y + float32(j)*s

			c := bg

			if (*m)[i] == 1 {
				c = fg
			}

			vector.FillRect(screen, pX, pY, s, s, c, false)

			i++
		}
	}
}

func (l *Gen04) Update() error {
	l.T += 1

	if l.T > D {
		l.B = !l.B
		l.T = 0
	}

	return nil
}

func (l *Gen04) Layout(w, h float32) {
	l.W = w
	l.H = h
}
