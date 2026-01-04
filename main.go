package main

import (
	"log"

	"github.com/e-kucheriavyi/genuary-2025/gen01"
	"github.com/e-kucheriavyi/genuary-2025/gen02"
	"github.com/e-kucheriavyi/genuary-2025/gen03"
	"github.com/e-kucheriavyi/genuary-2025/gen04"
	"github.com/e-kucheriavyi/genuary-2025/menu"
	"github.com/hajimehoshi/ebiten/v2"
)

type Level interface {
	Draw(screen *ebiten.Image)
	Update() error
	NextLevel() string
	IsLevel(nl string) bool
	Layout(w, h float32)
}

const (
	InitialW = 600
	InitialH = 600
)

type Game struct {
	W            float32
	H            float32
	Menu         *menu.Menu
	Levels       []Level
	CurrentLevel Level
}

func NewGame() *Game {
	m := menu.New(InitialW, InitialH)

	g01 := gen01.New()
	g02 := gen02.New()
	g03 := gen03.New()
	g04 := gen04.New()

	g := &Game{
		W: InitialW,
		H: InitialH,
		Levels: []Level{
			m,
			g01,
			g02,
			g03,
			g04,
		},
		CurrentLevel: m,
	}

	return g
}

func (g *Game) Next(nl string) {
	if nl == "" {
		return
	}

	for _, l := range g.Levels {
		if l.IsLevel(nl) {
			g.CurrentLevel = l
			return
		}
	}
}

func (g *Game) Update() error {
	err := g.CurrentLevel.Update()

	if err != nil {
		return err
	}

	g.Next(g.CurrentLevel.NextLevel())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentLevel.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	g.W = float32(w)
	g.H = float32(h)

	for _, l := range g.Levels {
		l.Layout(g.W, g.H)
	}

	return w, h
}

func main() {
	game := NewGame()

	ebiten.SetWindowTitle("Genuary 2025")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(InitialW, InitialH)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err.Error())
	}
}
