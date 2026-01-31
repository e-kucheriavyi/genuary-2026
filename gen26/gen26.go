package gen26

import (
	"image/color"
	"math/rand"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-26
// Recursive grids

const (
	InitialW = 640
	InitialH = 480
	MaxD     = 8
	MaxT     = 32
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen26 struct {
	W   float32
	H   float32
	T   int
	Img *ebiten.Image
}

func New() *Gen26 {
	l := &Gen26{
		W:   InitialW,
		H:   InitialH,
		Img: ebiten.NewImage(InitialW, InitialH),
	}

	return l
}

func (l *Gen26) IsLevel(nl string) bool {
	return nl == "26"
}

func (l *Gen26) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen26) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	if l.Img == nil {
		return
	}

	DrawGrid(l.Img, 0, 0, l.W, l.H, 0)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(l.Img, op)
}

func DrawGrid(img *ebiten.Image, x, y, w, h float32, d int) {
	if d > MaxD {
		return
	}
	xc := x + w*0.5
	yc := y + h*0.5

	vector.StrokeLine(img, x, yc, x+w, yc, 2, fg, false)
	vector.StrokeLine(img, xc, y, xc, y+h, 2, fg, false)

	r := rand.Float32()

	wh := w * 0.5
	hh := h * 0.5

	if r > 0.75 {
		DrawGrid(img, x, y, wh, hh, d+1)
	} else if r > 0.5 {
		DrawGrid(img, xc, y, wh, hh, d+1)
	} else if r > 0.25 {
		DrawGrid(img, x, yc, wh, hh, d+1)
	} else {
		DrawGrid(img, xc, yc, wh, hh, d+1)
	}
}

func (l *Gen26) Update() error {
	l.T += 1
	if l.T >= MaxT {
		l.T = 0
		l.Img = nil
		l.Img = ebiten.NewImage(int(l.W), int(l.H))
	}
	return nil
}

func (l *Gen26) Layout(w, h float32) {
	l.W = w
	l.H = h
}
