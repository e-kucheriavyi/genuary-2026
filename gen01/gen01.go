package gen01

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	InitialW = 640
	InitialH = 480
	TriW     = 32
	Gap      = 8
	Step     = 0.1
)

var bg = color.RGBA{0, 100, 0, 255}
var triC = color.RGBA{0, 150, 0, 255}

type Gen01 struct {
	W      float32
	H      float32
	Offset int
}

func New() *Gen01 {
	l := &Gen01{
		W:      InitialW,
		H:      InitialH,
		Offset: 0,
	}

	return l
}

func (l *Gen01) IsLevel(nl string) bool {
	return nl == "01"
}

func (l *Gen01) NextLevel() string {
	if input.IsPressed() {
		return "menu"
	}
	return ""
}

func (l *Gen01) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, l.W, l.H, bg, false)

	for i := range int(l.H/TriW) + 1 {
		x := (float32(l.Offset) * Step) * -1

		if i%2 == 0 {
			x = x*-1 - TriW*10
		}

		l.DrawTeeth(screen, x, float32(i*TriW))
	}
}

func (l *Gen01) DrawTeeth(screen *ebiten.Image, x, y float32) {
	p := &vector.Path{}

	for {
		p.MoveTo(x, y)
		p.LineTo(x+TriW/2, y+TriW)
		p.LineTo(x+TriW, y)
		p.Close()

		if x + TriW < l.W {
			x += TriW
		} else {
			break
		}
	}

	c := ebiten.ColorScale{}

	c.ScaleWithColor(triC)

	fillOpts := &vector.FillOptions{}
	pathOts := &vector.DrawPathOptions{
		AntiAlias:  true,
		ColorScale: c,
	}

	vector.FillPath(
		screen,
		p,
		fillOpts,
		pathOts,
	)
}

func (l *Gen01) Update() error {
	l.Offset += 1

	if float32(l.Offset)*Step >= TriW {
		l.Offset = 0
	}

	return nil
}

func (l *Gen01) Layout(w, h float32) {
	l.W = w
	l.H = h
}
