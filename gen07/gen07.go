package gen07

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-07
// Boolean algebra

const (
	InitialW = 640
	InitialH = 480
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{150, 0, 0, 255}

type Gen07 struct {
	W float32
	H float32
}

func New() *Gen07 {
	l := &Gen07{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen07) IsLevel(nl string) bool {
	return nl == "07"
}

func (l *Gen07) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	if input.IsPressed() {
		x, y := input.CursorPosition()

		p := float32(text.LetterWidth * 4 * 2)

		if x <= p && y <= p {
			return "menu"
		}
	}
	return ""
}

func (l *Gen07) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, l.W, l.H, bg, false)

	tl := float32(5)
	s := l.GetTextS()
	x := l.W / 2 - ((tl * text.LetterWidth * s) / 2)
	y := l.H / 2 - (text.LetterWidth * s / 2)

	text.Write(screen, "false", x, y, s, fg)
	text.DrawLetter(screen, 'x', 10, 10, 4, color.White)
}

func (l *Gen07) GetTextS() float32 {
	tl := float32(5)
	sW := l.W / text.LetterWidth / tl

	sH := l.H / text.LetterWidth

	if sH < sW {
		return sH
	}

	return sW
}

func (l *Gen07) Update() error {
	return nil
}

func (l *Gen07) Layout(w, h float32) {
	l.W = w
	l.H = h
}
