package menu

import (
	"image/color"
	"strings"
	"unicode/utf8"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	la "github.com/laranatech/gorana/layout"
)

var bg = color.Black
var fg = color.White
var pale = color.RGBA{100, 100, 100, 255}
var palest = color.RGBA{50, 50, 50, 255}

const (
	HeaderH = 64
	BtnW    = 64
)

type Menu struct {
	W        float32
	H        float32
	Root     *la.OutputItem
	Hovered  *la.OutputItem
	Selected string
}

func New(w, h float32) *Menu {
	m := &Menu{}

	m.Layout(w, h)

	return m
}

func (m *Menu) IsLevel(nl string) bool {
	return nl == "menu"
}

func (m *Menu) Draw(screen *ebiten.Image) {
	m.DrawNode(screen, m.Root)
}

func (m *Menu) DrawNode(screen *ebiten.Image, node *la.OutputItem) {
	if node.Id == "header" {
		m.DrawHeader(screen, node)
	} else if strings.HasPrefix(node.Id, "btn_") {
		m.DrawBtn(screen, node)
	}

	for _, child := range node.Children {
		m.DrawNode(screen, child)
	}
}

func (m *Menu) DrawHeader(screen *ebiten.Image, node *la.OutputItem) {
	s := float32(3)
	txt := "genuary 2026"
	ln := float32(utf8.RuneCountInString(txt))

	x := node.X + (node.W / 2) - ((ln * text.LetterWidth * s) / 2)
	y := node.Y + (node.H / 2) - ((text.LetterWidth * 2) / 2)

	text.Write(screen, txt, x, y, s, color.White)
}

func (m *Menu) DrawBtn(screen *ebiten.Image, node *la.OutputItem) {
	var c color.Color = pale

	if strings.HasSuffix(node.Id, "_disabled") {
		c = palest
	} else if m.Hovered != nil && m.Hovered.Id == node.Id {
		c = fg
	}

	vector.StrokeRect(
		screen,
		node.X,
		node.Y,
		node.W,
		node.H,
		2,
		c,
		false,
	)

	t := strings.Split(node.Id, "_")[1]

	s := float32(3)

	x := node.X + (node.W / 2) - ((text.LetterWidth * s * 2) / 2)
	y := node.Y + (node.H / 2) - ((text.LetterWidth * s) / 2)

	text.Write(screen, t, x, y, s, c)
}

func (m *Menu) Update() error {
	m.Selected = ""

	x, y := input.CursorPosition()

	m.Hovered = input.FindHovered(m.Root, x, y)

	if !input.IsPressed() {
		return nil
	}

	if m.Hovered == nil {
		return nil
	}

	m.Selected = m.Hovered.Id

	return nil
}

func (m *Menu) NextLevel() string {
	if m.Selected != "" {
		return strings.Split(m.Selected, "_")[1]
	}
	return ""
}

func (m *Menu) Layout(w, h float32) {
	m.W = w
	m.H = h

	root := la.Node(
		la.Id("root"),
		la.Column(),
		la.Width(la.Fix(w)),
		la.Height(la.Fix(h)),
		la.Gap(8),
		la.Padding(8),
		la.Children(
			la.Node(
				la.Id("header"),
				la.Height(la.Fix(HeaderH)),
				la.Width(la.Grow(1)),
			),
			la.Node(
				la.Id("body"),
				la.Column(),
				la.Width(la.Grow(1)),
				la.Height(la.Grow(1)),
				la.Gap(8),
				la.Children(
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("01"),
							btn("02"),
							btn("03_disabled"),
							btn("04_disabled"),
							btn("05_disabled"),
							btn("06_disabled"),
							btn("07_disabled"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("08_disabled"),
							btn("09_disabled"),
							btn("11_disabled"),
							btn("12_disabled"),
							btn("13_disabled"),
							btn("14_disabled"),
							btn("15_disabled"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("16_disabled"),
							btn("17_disabled"),
							btn("18_disabled"),
							btn("19_disabled"),
							btn("20_disabled"),
							btn("21_disabled"),
							btn("22_disabled"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("23_disabled"),
							btn("24_disabled"),
							btn("25_disabled"),
							btn("26_disabled"),
							btn("27_disabled"),
							btn("28_disabled"),
							btn("29_disabled"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("30_disabled"),
							btn("31_disabled"),
							la.Node(la.Width(la.Grow(1))),
							la.Node(la.Width(la.Grow(1))),
							la.Node(la.Width(la.Grow(1))),
							la.Node(la.Width(la.Grow(1))),
							la.Node(la.Width(la.Grow(1))),
						),
					),
				),
			),
		),
	)

	la.Layout(root)

	m.Root = la.Export(root)
}

func btn(id string) *la.NodeItem {
	return la.Node(
		la.Id("btn_"+id),
		la.Width(la.Grow(1)),
		la.Height(la.Grow(1)),
	)
}
