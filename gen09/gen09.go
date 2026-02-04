package gen09

import (
	"image/color"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-09
// Crazy automaton

const (
	InitialW = 640
	InitialH = 480
	MaxT     = 2
)

var bg = color.RGBA{0, 0, 0, 255}
var fg = color.RGBA{0, 150, 0, 255}

type Gen09 struct {
	W     float32
	H     float32
	T     int
	Board [][]bool
}

func NewBoard() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}
}

func New() *Gen09 {
	l := &Gen09{
		W:     InitialW,
		H:     InitialH,
		Board: NewBoard(),
	}

	return l
}

func (l *Gen09) IsLevel(nl string) bool {
	return nl == "09"
}

func (l *Gen09) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen09) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	w := l.W / 10
	h := l.H / 10

	for i, r := range l.Board {
		for j, c := range r {
			if !c {
				continue
			}
			vector.FillRect(screen, float32(i)*w, float32(j)*h, w, h, fg, false)
		}
	}
}

func (l *Gen09) Update() error {
	l.T += 1

	if l.T >= MaxT {
		l.T = 0
		l.Live()
	}

	return nil
}

func (l *Gen09) Live() {
	nb := l.Board

	for i, r := range l.Board {
		for j, c := range r {
			if j == 0 {
				continue
			}
			if c == l.Board[i][j-1] {
				nb[i][j] = !l.Board[i][j]
			} else {
				nb[i][j-1] = !l.Board[i][j-1]
			}
			if i == 0 {
				continue
			}
			if c == l.Board[i-1][j] {
				nb[i][j] = !l.Board[i][j]
			} else {
				nb[i-1][j] = !l.Board[i-1][j]
			}
		}
	}

	l.Board = nb
}

func (l *Gen09) Layout(w, h float32) {
	l.W = w
	l.H = h
}
