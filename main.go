package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/e-kucheriavyi/genuary-2025/gen01"
	"github.com/e-kucheriavyi/genuary-2025/gen02"
	"github.com/e-kucheriavyi/genuary-2025/gen03"
	"github.com/e-kucheriavyi/genuary-2025/gen04"
	"github.com/e-kucheriavyi/genuary-2025/gen05"
	"github.com/e-kucheriavyi/genuary-2025/gen06"
	"github.com/e-kucheriavyi/genuary-2025/gen07"
	"github.com/e-kucheriavyi/genuary-2025/gen08"
	"github.com/e-kucheriavyi/genuary-2025/gen09"
	"github.com/e-kucheriavyi/genuary-2025/gen10"
	"github.com/e-kucheriavyi/genuary-2025/gen11"
	"github.com/e-kucheriavyi/genuary-2025/gen12"
	"github.com/e-kucheriavyi/genuary-2025/gen13"
	"github.com/e-kucheriavyi/genuary-2025/gen15"
	"github.com/e-kucheriavyi/genuary-2025/gen16"
	"github.com/e-kucheriavyi/genuary-2025/gen17"
	"github.com/e-kucheriavyi/genuary-2025/gen18"
	"github.com/e-kucheriavyi/genuary-2025/gen19"
	"github.com/e-kucheriavyi/genuary-2025/gen20"
	"github.com/e-kucheriavyi/genuary-2025/gen21"
	"github.com/e-kucheriavyi/genuary-2025/gen22"
	"github.com/e-kucheriavyi/genuary-2025/gen23"
	"github.com/e-kucheriavyi/genuary-2025/gen24"
	"github.com/e-kucheriavyi/genuary-2025/gen26"
	"github.com/e-kucheriavyi/genuary-2025/gen27"
	"github.com/e-kucheriavyi/genuary-2025/gen28"
	"github.com/e-kucheriavyi/genuary-2025/gen30"
	"github.com/e-kucheriavyi/genuary-2025/gen31"
	"github.com/e-kucheriavyi/genuary-2025/menu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	I            int
	Recording    bool
	RecordAt     time.Time
}

func NewGame() *Game {
	m := menu.New(InitialW, InitialH)

	g01 := gen01.New()
	g02 := gen02.New()
	g03 := gen03.New()
	g04 := gen04.New()
	g05 := gen05.New()
	g06 := gen06.New()
	g07 := gen07.New()
	g08 := gen08.New()
	g09 := gen09.New()
	g10 := gen10.New()
	g11 := gen11.New()
	g12 := gen12.New()
	g13 := gen13.New()
	g15 := gen15.New()
	g16 := gen16.New()
	g17 := gen17.New()
	g18 := gen18.New()
	g19 := gen19.New()
	g20 := gen20.New()
	g21 := gen21.New()
	g22 := gen22.New()
	g23 := gen23.New()
	g24 := gen24.New()
	g26 := gen26.New()
	g27 := gen27.New()
	g28 := gen28.New()
	g30 := gen30.New()
	g31 := gen31.New()

	g := &Game{
		W: InitialW,
		H: InitialH,
		Levels: []Level{
			m,
			g01, g02, g03, g04, g05, g06, g07,
			g08, g09, g10, g11, g12, g13,
			g15, g16, g17, g18, g19, g20, g21,
			g22, g23, g24, g26, g27, g28,
			g30, g31,
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
			g.CurrentLevel.Layout(g.W, g.H)
			return
		}
	}
}

func (g *Game) Update() error {
	err := g.CurrentLevel.Update()

	if err != nil {
		return err
	}

	g.ToggleRecording()

	if g.Recording {
		g.I += 1
	}

	g.Next(g.CurrentLevel.NextLevel())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentLevel.Draw(screen)

	if g.Recording {
		g.Record(screen)
	}
}

func (g *Game) Record(screen *ebiten.Image) {
	dir := fmt.Sprintf("dist/%d", g.RecordAt.Unix())
	os.MkdirAll(dir, 0777)
	f, err := os.Create(fmt.Sprintf("%s/%06d.png", dir, g.I))

	if err != nil {
		log.Fatal(err)
	}

	if err = png.Encode(f, screen); err != nil {
		f.Close()
		log.Fatal(err)
	}
}

func (g *Game) ToggleRecording() {
	keys := inpututil.AppendJustReleasedKeys(nil)

	for _, key := range keys {
		if key != ebiten.KeyR {
			continue
		}

		g.Recording = !g.Recording
		g.I = 0

		if g.Recording {
			g.RecordAt = time.Now()
		}
		return
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	g.W = float32(w)
	g.H = float32(h)

	g.CurrentLevel.Layout(g.W, g.H)

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
