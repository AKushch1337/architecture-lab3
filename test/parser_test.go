package test

import (
	"github.com/AKushch1337/architecture-lab3/painter"
	"github.com/AKushch1337/architecture-lab3/painter/lang"
	"image/color"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parse_func(t *testing.T) {
	tests := []struct {
		command string
		op      painter.Operation
	}{
		{
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			command: "reset",
			op:      painter.OperationFunc(painter.Reset),
		},
	}
	parser := &lang.Parser{}
	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			op, _ := parser.Parse(strings.NewReader(tt.command))
			assert.Equal(t, reflect.TypeOf(tt.op), reflect.TypeOf(op[0]))
		})
	}
}

func Test_parse_struct(t *testing.T) {
	tests := []struct {
		command string
		op      painter.Operation
	}{
		{
			command: "bgrect 0 0 100 100",
			op:      &painter.BgRectOp{X1: 0, Y1: 0, X2: 100, Y2: 100},
		},
		{
			command: "figure 200 200",
			op:      &painter.FigureOp{X: 200, Y: 200, C: color.RGBA{R: 219, G: 208, B: 48, A: 1}},
		},
		{
			command: "move 100 100",
			op:      &painter.MoveOp{X: 100, Y: 100},
		},
		{
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			command: "invalidcommand",
			op:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			parser := &lang.Parser{}
			op, err := parser.Parse(strings.NewReader(tt.command))
			if err != nil {
				assert.Nil(t, tt.op)
			} else {
				assert.Equal(t, reflect.TypeOf(tt.op), reflect.TypeOf(op[0]))
				if tt.op != nil {
					assert.Equal(t, tt.op, op[0])
				}
			}
		})
	}
}
