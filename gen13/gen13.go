package gen13

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-13
// Self portrait

const (
	InitialW = 640
	InitialH = 480
)

var bg = color.RGBA{0, 150, 0, 255}
var fg = color.RGBA{0, 100, 0, 255}

type Gen13 struct {
	W float32
	H float32
}

func New() *Gen13 {
	l := &Gen13{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen13) IsLevel(nl string) bool {
	return nl == "13"
}

func (l *Gen13) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen13) Draw(screen *ebiten.Image) {
	s := float32(0)

	if l.W > l.H {
		s = l.H
	} else {
		s = l.W
	}

	p := &vector.Path{}

	p.MoveTo(0.400*s, 0.040*s)
	p.LineTo(0.350*s, 0.050*s)
	p.LineTo(0.250*s, 0.150*s)
	p.LineTo(0.201*s, 0.265*s)
	p.LineTo(0.182*s, 0.379*s)
	p.LineTo(0.221*s, 0.426*s)
	p.LineTo(0.178*s, 0.543*s)
	p.LineTo(0.221*s, 0.562*s)
	p.LineTo(0.214*s, 0.595*s)
	p.LineTo(0.229*s, 0.615*s)
	p.LineTo(0.221*s, 0.628*s)
	p.LineTo(0.237*s, 0.654*s)
	p.LineTo(0.230*s, 0.737*s)
	p.LineTo(0.495*s, 0.654*s)
	p.LineTo(0.542*s, 0.414*s)
	p.LineTo(0.740*s, 0.314*s)
	p.LineTo(0.764*s, 0.175*s)
	p.LineTo(0.700*s, 0.075*s)
	p.LineTo(0.548*s, 0.039*s)
	p.LineTo(0.548*s, 0.040*s)
	p.LineTo(0.535*s, 0.040*s)
	p.Close()

	c := ebiten.ColorScale{}
	c.ScaleWithColor(fg)

	pathOpts := &vector.DrawPathOptions{ColorScale: c, AntiAlias: false}
	strokeOpts := &vector.StrokeOptions{Width: 16}

	vector.StrokePath(screen, p, strokeOpts, pathOpts)
}

func (l *Gen13) Update() error {
	return nil
}

func (l *Gen13) Layout(w, h float32) {
	l.W = w
	l.H = h
}
