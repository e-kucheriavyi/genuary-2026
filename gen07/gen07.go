package gen07

import (
	"image/color"
	"math"

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
	amp      = 70.0
	lambda   = 200.0
)

var bg = color.RGBA{0, 0, 0, 255}
var red = color.RGBA{150, 0, 0, 255}
var green = color.RGBA{0, 150, 0, 255}
var gray = color.RGBA{150, 150, 150, 255}

type Gen07 struct {
	W float32
	H float32
	T float64
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

	yCenter := float64(l.H / 2)

	vector.StrokeLine(screen, 0, float32(yCenter), l.W, float32(yCenter), 1, gray, false)

	s := float32(3)

	x := float64(0)

	twoPi := 2 * math.Pi

	for {
		phase := (x / lambda) * twoPi
		y := float64(yCenter) + math.Sin(phase+l.T)*amp

		char := '0'
		c := red

		if y < yCenter {
			char = '1'
			c = green
		}

		text.DrawLetter(screen, char, float32(x), float32(y), s, c)

		x += float64(text.LetterWidth * s)

		if x > float64(l.W) {
			break
		}
	}

	text.DrawLetter(screen, 'x', 10, 10, 4, color.White)
}

func (l *Gen07) Update() error {
	l.T += 0.05
	return nil
}

func (l *Gen07) Layout(w, h float32) {
	l.W = w
	l.H = h
}
