package gen28

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	la "github.com/laranatech/gorana/layout"
)

// 2025-01-28
// No libs, no canvas, just HTML
// Since there's no HTML here, using only gorana

const (
	InitialW = 640
	InitialH = 480
	D        = 64
)

var (
	bg     = color.RGBA{0, 0, 0, 255}
	fg     = color.RGBA{0, 150, 0, 255}
	blue   = color.RGBA{0x00, 0x6f, 0xb6, 0xff}
	yellow = color.RGBA{0xf3, 0xc3, 0x01, 0xff}
	red    = color.RGBA{0xd8, 0x00, 0x0b, 0xff}
)

type Gen28 struct {
	W    float32
	H    float32
	T    int
	D    int
	Root *la.OutputItem
}

func New() *Gen28 {
	l := &Gen28{
		W: InitialW,
		H: InitialH,
		D: 1,
	}

	return l
}

func (l *Gen28) IsLevel(nl string) bool {
	return nl == "28"
}

func (l *Gen28) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen28) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	if l.Root == nil {
		return
	}

	l.DrawNode(screen, l.Root)

	vector.FillRect(screen, 0, 0, 10, 10, fg, false)
}

func (l *Gen28) DrawNode(screen *ebiten.Image, node *la.OutputItem) {
	var c color.Color

	switch node.Id {
	case "row_1_1":
		c = blue
	case "row_1_2":
		c = yellow
	case "row_2_1":
		c = red
	case "row_2_2":
		c = fg
	}

	if c != nil {
		vector.FillRect(screen, node.X, node.Y, node.W, node.H, c, false)
	}

	for _, child := range node.Children {
		l.DrawNode(screen, child)
	}
}

func (l *Gen28) Update() error {
	l.T += l.D

	if l.T > D || l.T < 1 {
		l.D *= -1
		l.T += l.D
	}

	l.Lay()

	return nil
}

func (l *Gen28) Lay() {
	l.Root = nil

	root := la.Node(
		la.Id("root"),
		la.Width(la.Fix(l.W)),
		la.Height(la.Fix(l.H)),
		la.Column(),
		la.Gap(32),
		la.Children(
			la.Node(
				la.Id("row_1"),
				la.Gap(32),
				la.Height(la.Grow(1)),
				la.Width(la.Grow(1)),
				la.Children(
					la.Node(
						la.Id("row_1_1"),
						la.Width(la.Grow(float32(l.T))),
						la.Height(la.Grow(1)),
					),
					la.Node(
						la.Id("row_1_2"),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
					),
				),
			),
			la.Node(
				la.Id("row_2"),
				la.Gap(32),
				la.Height(la.Grow(1)),
				la.Width(la.Grow(1)),
				la.Children(
					la.Node(
						la.Id("row_2_1"),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
					),
					la.Node(
						la.Id("row_2_2"),
						la.Width(la.Grow(float32(l.T))),
						la.Height(la.Grow(1)),
					),
				),
			),
		),
	)

	la.Layout(root)

	l.Root = la.Export(root)
}

func (l *Gen28) Layout(w, h float32) {
	l.W = w
	l.H = h
}
