package gen02

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-02
// 12 principles of animation

const (
	InitialW = 640
	InitialH = 480
	TriW     = 64
	Gap      = 8
	Step     = 0.1
)

var bg = color.RGBA{0, 100, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type KeyFrame struct {
	X float32
	Y float32
	W float32
	H float32
	D int
}

type Gen02 struct {
	X         float32
	Y         float32
	W         float32
	H         float32
	KeyFrames []*KeyFrame
	I         int
	T         int
}

func lerp(a, b, t float32) float32 {
	return a + (b-a)*t
}

func New() *Gen02 {
	l := &Gen02{
		W: InitialW,
		H: InitialH,
		KeyFrames: []*KeyFrame{
			{TriW * 2, 0, TriW, TriW, 20},
			{TriW * 2, 0, TriW * 1.5, TriW * 0.5, 20},
			{TriW * 2, 0, TriW * 1.5, TriW * 0.5, 20},
			{TriW * 2, TriW * 3, TriW * 0.5, TriW * 1.5, 20},
			{TriW * 2, TriW * 5, TriW * 0.5, TriW * 1.5, 20},
			{TriW * 2, TriW * 6, TriW * 1, TriW * 1, 20},
			{TriW * 2, 0, TriW * 0.8, TriW * 1.2, 40},
			{TriW * 2, 0, TriW * 1.5, TriW * 0.5, 10},
			{TriW * 2, 0, TriW, TriW, 20},
		},
	}

	return l
}

func (l *Gen02) IsLevel(nl string) bool {
	return nl == "02"
}

func (l *Gen02) NextLevel() string {
	if input.IsPressed() || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	return ""
}

func (l *Gen02) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, l.W, l.H, bg, false)

	x, y, w, h := l.GetBox()

	vector.FillRect(screen, x, y, w, h, fg, false)
}

func (l *Gen02) GetBox() (x, y, w, h float32) {
	c := l.KeyFrames[l.I]

	var p *KeyFrame

	if l.I == 0 {
		p = l.KeyFrames[len(l.KeyFrames)-1]
	} else {
		p = l.KeyFrames[l.I-1]
	}

	y = utils.Lerp(p.Y, c.Y, float32(l.T)/float32(c.D))
	w = utils.Lerp(p.W, c.W, float32(l.T)/float32(c.D))
	h = utils.Lerp(p.H, c.H, float32(l.T)/float32(c.D))
	x = utils.Lerp(p.X, c.X, float32(l.T)/float32(c.D)) - (w / 2)

	return x, y, w, h
}

func (l *Gen02) Update() error {
	l.T += 1

	if l.T > l.KeyFrames[l.I].D {
		l.T = 0
		l.I++
		if l.I >= len(l.KeyFrames) {
			l.I = 0
		}
	}
	return nil
}

func (l *Gen02) Layout(w, h float32) {
	l.W = w
	l.H = h
}
