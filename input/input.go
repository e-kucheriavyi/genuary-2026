package input

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	la "github.com/laranatech/gorana/layout"
)

const (
	clickInputDebounce = 250
)

func FindHovered(node *la.OutputItem, x, y float32) *la.OutputItem {
	if strings.HasPrefix(node.Id, "btn_") && !strings.HasSuffix(node.Id, "_disabled") {
		if Collide(node, x, y) {
			return node
		}
	}

	for _, child := range node.Children {
		hovered := FindHovered(child, x, y)
		if hovered != nil {
			return hovered
		}
	}

	return nil
}

func Collide(node *la.OutputItem, x, y float32) bool {
	if x < node.X || x > node.X+node.W {
		return false
	}

	if y < node.Y || y > node.Y+node.H {
		return false
	}

	return true
}

func IsPressed() bool {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		return true
	}

	touches := inpututil.AppendJustReleasedTouchIDs(nil)

	return len(touches) > 0
}

func CursorPosition() (float32, float32) {
	x, y := ebiten.CursorPosition()
	if x != 0 && y != 0 {
		return float32(x), float32(y)
	}

	touches := inpututil.AppendJustReleasedTouchIDs(nil)

	if len(touches) == 0 {
		return 0, 0
	}

	x, y = ebiten.TouchPosition(touches[0])

	return float32(x), float32(y)
}
