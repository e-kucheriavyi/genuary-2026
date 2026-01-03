package gen03

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-03
// Fibonacci

const (
	InitialW = 640
	InitialH = 480
	Step     = 0.3
)

var bg = color.RGBA{0, 100, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen03 struct {
	W float32
	H float32
	S float32
	D float32
}

func lerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

func New() *Gen03 {
	l := &Gen03{
		W: InitialW,
		H: InitialH,
		S: 50,
		D: -1,
	}

	return l
}

func (l *Gen03) IsLevel(nl string) bool {
	return nl == "03"
}

func (l *Gen03) NextLevel() string {
	if input.IsPressed() || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	return ""
}

func fib(i int) int {
	if i <= 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func (l *Gen03) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, l.W, l.H, bg, false)

	for i := range 20 {
		w := float32(fib(i)) * l.S
		h := float32(fib(i)) * l.S

		x := l.W/2 - w/2
		y := l.H/2 - h/2

		vector.StrokeRect(screen, x, y, w, h, 2, fg, true)
	}

	vector.StrokeLine(
		screen,
		0,
		0,
		l.W,
		l.H,
		2,
		fg,
		true,
	)

	vector.StrokeLine(
		screen,
		l.W,
		0,
		0,
		l.H,
		2,
		fg,
		true,
	)
}

func (l *Gen03) Update() error {
	l.S += Step * l.D

	if l.S <= 0 || l.S >= 100 {
		l.D *= -1
	}

	return nil
}

func (l *Gen03) Layout(w, h float32) {
	l.W = w
	l.H = h
}
