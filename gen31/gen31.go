package gen31

import (
	_ "embed"
	"image/color"
	"log"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/hajimehoshi/ebiten/v2"
)

// 2025-01-31
// GLSL day

const (
	InitialW = 640
	InitialH = 480
)

//go:embed shader.kage
var shaderSrc []byte

var shader *ebiten.Shader

var (
	bg = color.RGBA{0, 0, 0, 255}
	fg = color.RGBA{0, 150, 0, 255}
)

type Gen31 struct {
	W float32
	H float32
}

func New() *Gen31 {
	s, err := ebiten.NewShader(shaderSrc)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	shader = s

	l := &Gen31{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen31) IsLevel(nl string) bool {
	return nl == "31"
}

func (l *Gen31) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || input.IsPressed() {
		return "menu"
	}

	return ""
}

func (l *Gen31) Draw(screen *ebiten.Image) {
	screen.Fill(bg)

	op := &ebiten.DrawRectShaderOptions{}
	screen.DrawRectShader(int(l.W), int(l.H), shader, op)
}

func (l *Gen31) Update() error {
	return nil
}

func (l *Gen31) Layout(w, h float32) {
	l.W = w
	l.H = h
}
