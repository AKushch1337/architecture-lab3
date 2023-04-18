package painter

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRectOp struct {
	X1, Y1, X2, Y2 int
}

func (op *BgRectOp) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.X1, op.Y1, op.X2, op.Y2), color.Black, screen.Src)
	return false
}

type FigureOp struct {
	X, Y int
	C    color.RGBA
}

func (op *FigureOp) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.X-200, op.Y+100, op.X+200, op.Y-100), op.C, draw.Src)
	t.Fill(image.Rect(op.X-100, op.Y-200, op.X+100, op.Y+200), op.C, draw.Src)
	return false
}

type MoveOp struct {
	X, Y    int
	Figures []FigureOp
}

func (op *MoveOp) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].X = op.X
		op.Figures[i].Y = op.Y
		op.Figures[i].Do(t)
	}
	return false
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, draw.Src)
}
