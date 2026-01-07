package gen06

import (
	"image"
	"image/color"
	"math"

	"github.com/e-kucheriavyi/genuary-2025/input"
	"github.com/e-kucheriavyi/genuary-2025/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// 2025-01-06
// Lights on/off

const (
	InitialW = 640
	InitialH = 480
	LD       = 0.1
	R        = 0.2
)

var (
	bgImage       *ebiten.Image
	fgImage       *ebiten.Image
	maskedFgImage *ebiten.Image
	spotImage     *ebiten.Image
)

var bg = color.RGBA{0, 0, 0, 255}
var pale = color.RGBA{10, 10, 10, 255}

var fg = color.RGBA{0, 150, 0, 255}

type Gen06 struct {
	W float32
	H float32
	X float32
	Y float32
	R float32
}

func New() *Gen06 {
	l := &Gen06{
		W: InitialW,
		H: InitialH,
	}

	return l
}

func (l *Gen06) IsLevel(nl string) bool {
	return nl == "06"
}

func (l *Gen06) NextLevel() string {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return "menu"
	}
	if input.IsPressed() {
		x, y := input.CursorPosition()

		p := float32(text.LetterWidth * 4 * 2)

		if x <= p && y <= p {
			return "menu"
		}
	}
	return ""
}

func (l *Gen06) Draw(screen *ebiten.Image) {
	maskedFgImage.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.Blend = ebiten.BlendCopy
	op.GeoM.Translate(float64(l.X-l.R), float64(l.Y-l.R))
	if spotImage != nil {
		maskedFgImage.DrawImage(spotImage, op)
	}

	op = &ebiten.DrawImageOptions{}
	op.Blend = ebiten.BlendSourceIn
	maskedFgImage.DrawImage(fgImage, op)

	screen.Fill(color.RGBA{0x00, 0x00, 0x80, 0xff})
	screen.DrawImage(bgImage, &ebiten.DrawImageOptions{})
	screen.DrawImage(maskedFgImage, &ebiten.DrawImageOptions{})

	text.DrawLetter(screen, 'x', 10, 10, 4, color.White)
}

func (l *Gen06) Update() error {
	x, y := input.CursorPosition()

	l.X = x
	l.Y = y

	return nil
}

func DrawImage(screen *ebiten.Image, w, h float32, bg, fg color.Color) {
	vector.FillRect(screen, 0, 0, w, h, bg, false)

	lW := LD * w
	lH := LD * h

	p := &vector.Path{}

	for i := range int(h/(h*LD)) + 1 {
		y := lH * float32(i)
		for j := range int(w/(w*LD)) + 1 {
			x := lW * float32(j)

			if j == 0 {
				p.MoveTo(x, y+lH*0.5)
				continue
			}

			if j%2 == 0 {
				p.LineTo(x, y+lH*0.5)
				continue
			}

			p.LineTo(x, y)
		}
	}

	c := ebiten.ColorScale{}

	c.ScaleWithColor(fg)

	strokeOpts := &vector.StrokeOptions{Width: 5, LineCap: 2, LineJoin: 2}
	pathOpts := &vector.DrawPathOptions{
		ColorScale: c,
		AntiAlias:  true,
	}

	vector.StrokePath(screen, p, strokeOpts, pathOpts)
}

func (l *Gen06) Layout(w, h float32) {
	l.W = w
	l.H = h

	bgImage = ebiten.NewImage(int(w), int(h))
	fgImage = ebiten.NewImage(int(w), int(h))
	maskedFgImage = ebiten.NewImage(int(w), int(h))

	DrawImage(bgImage, w, h, bg, fg)
	DrawImage(fgImage, w, h, bg, pale)

	r := int(l.W * R)
	if l.H < l.W {
		r = int(l.H * R)
	}
	l.R = float32(r)

	alphas := image.Point{r * 2, r * 2}
	a := image.NewAlpha(image.Rectangle{image.Point{}, alphas})
	for j := range alphas.Y {
		for i := range alphas.X {
			d := math.Sqrt(float64((i-r)*(i-r) + (j-r)*(j-r)))
			b := uint8(max(0, min(0xff, int(3*d*0xff/float64(r))-2*0xff)))
			a.SetAlpha(i, j, color.Alpha{b})
		}
	}
	spotImage = ebiten.NewImageFromImage(a)
}
