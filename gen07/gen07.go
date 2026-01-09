package gen07

import (
	"image/color"
	"math"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/e-kucheriavyi/genuary-2025/utils"
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

var (
	bg    = color.RGBA{0, 0, 0, 255}
	red   = color.RGBA{150, 0, 0, 255}
	green = color.RGBA{0, 150, 0, 255}
	gray  = color.RGBA{150, 150, 150, 255}
)

var twoPi = math.Pi * 2

var (
	l1 = text.GetLetterMap('1')
	l0 = text.GetLetterMap('0')
)

type Gen07 struct {
	W      float32
	H      float32
	YC     float32
	S      float32
	T      float64
	Offset float64
}

func New() *Gen07 {
	l := &Gen07{
		W: InitialW,
		H: InitialH,
		S: 3,
	}

	return l
}

func (l *Gen07) IsLevel(nl string) bool {
	return nl == "07"
}

func (l *Gen07) NextLevel() string {
	if input.IsPressed() || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	return ""
}

func (l *Gen07) Draw(screen *ebiten.Image) {
	vector.StrokeLine(screen, 0, float32(l.YC), l.W, float32(l.YC), 1, gray, false)

	x := float64(0)

	for {
		phase := (x / lambda) * twoPi
		y := float64(l.YC) + math.Sin(phase+l.T)*amp

		char := l0
		c := red

		if float32(y) < l.YC {
			char = l1
			c = green
		}

		utils.DrawBitmap(
			screen,
			char,
			float32(x),
			float32(y),
			l.S,
			text.LetterWidth,
			c,
		)

		x += l.Offset

		if x > float64(l.W) {
			break
		}
	}
}

func (l *Gen07) Update() error {
	l.T += 0.05
	return nil
}

func (l *Gen07) Layout(w, h float32) {
	l.W = w
	l.H = h
	l.YC = h / 2
	l.Offset = float64(text.LetterWidth * l.S)
}
