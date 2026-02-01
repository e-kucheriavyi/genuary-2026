package menu

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	X        float32
	Y        float32
	Root     *la.OutputItem
	Focused  string
	Hovered  string
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
	} else if m.Focused == node.Id || m.Hovered == node.Id {
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

	m.UpdateKeyboard()

	if x == m.X && y == m.Y {
		if input.IsPressed() && m.Focused != "" {
			m.Selected = m.Focused
		}
		return nil
	}

	m.X = x
	m.Y = y

	hovered := input.FindHovered(m.Root, x, y)

	if hovered != nil {
		m.Focused = hovered.Id
	}

	if !input.IsPressed() {
		return nil
	}

	if m.Focused == "" {
		return nil
	}

	m.Selected = m.Focused

	return nil
}

func (m *Menu) UpdateKeyboard() error {
	keys := inpututil.AppendJustReleasedKeys(nil)

	for _, key := range keys {
		switch key {
		case ebiten.KeyArrowLeft, ebiten.KeyH:
			return m.Focus(-1)
		case ebiten.KeyArrowRight, ebiten.KeyL:
			return m.Focus(+1)
		case ebiten.KeyArrowUp, ebiten.KeyK:
			return m.Focus(-7)
		case ebiten.KeyArrowDown, ebiten.KeyJ:
			return m.Focus(+7)
		case ebiten.KeyEnter:
			m.Selected = m.Focused
			return nil
		}
	}
	return nil
}

func (m *Menu) Focus(d int) error {
	if m.Focused == "" {
		m.Focused = "btn_01"
		return nil
	}

	id := strings.Split(m.Focused, "_")[1]
	curr, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	curr += d

	if curr <= 0 {
		curr = 1
	}
	if curr > 31 {
		curr = 31
	}

	currS := fmt.Sprintf("btn_%02d", curr)

	if d > 0 {
		d = 1
	} else {
		d = -1
	}

	m.Focused = currS

	if curr >= 31 {
		curr = 31
		return nil
	}

	if !IsDisabled(m.Root, currS) {
		return nil
	}

	return m.Focus(d)
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
							btn("03"),
							btn("04"),
							btn("05"),
							btn("06"),
							btn("07"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("08"),
							btn("09_disabled"),
							btn("10"),
							btn("11"),
							btn("12"),
							btn("13"),
							btn("14_disabled"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("15"),
							btn("16"),
							btn("17"),
							btn("18"),
							btn("19"),
							btn("20"),
							btn("21"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("22"),
							btn("23"),
							btn("24"),
							btn("25"),
							btn("26"),
							btn("27"),
							btn("28"),
						),
					),
					la.Node(
						la.Gap(8),
						la.Row(),
						la.Width(la.Grow(1)),
						la.Height(la.Grow(1)),
						la.Children(
							btn("29_disabled"),
							btn("30"),
							btn("31"),
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

func IsDisabled(node *la.OutputItem, id string) bool {
	if node.Id == id {
		return false
	}

	for _, child := range node.Children {
		d := IsDisabled(child, id)
		if d == false {
			return d
		}
	}

	return true
}

func btn(id string) *la.NodeItem {
	return la.Node(
		la.Id("btn_"+id),
		la.Width(la.Grow(1)),
		la.Height(la.Grow(1)),
	)
}
