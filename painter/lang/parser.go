package lang

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/AKushch1337/architecture-lab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	var res []painter.Operation
	for scanner.Scan() {
		commandLine := scanner.Text()
		op := parse(commandLine) // parse the line to get Operation
		if op == nil {
			return nil, fmt.Errorf("failed to parse command: %s", commandLine)
		}
		if bgRectOp, ok := op.(*painter.BgRectOp); ok {
			for i, oldOp := range res {
				if _, ok := oldOp.(*painter.BgRectOp); ok {
					res[i] = bgRectOp
					break
				}
			}
		}
		res = append(res, op)
	}
	return res, nil
}

func parse(commandLine string) painter.Operation {
	parts := strings.Split(commandLine, " ")
	instruction := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	var intArgs []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err == nil {
			intArgs = append(intArgs, i)
		}
	}

	var figureOps []painter.FigureOp

	switch instruction {
	case "white":
		return painter.OperationFunc(painter.WhiteFill)
	case "green":
		return painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		return &painter.BgRectOp{X1: intArgs[0], Y1: intArgs[1], X2: intArgs[2], Y2: intArgs[3]}
	case "figure":
		col := color.RGBA{R: 219, G: 208, B: 48, A: 1}
		figure := painter.FigureOp{X: intArgs[0], Y: intArgs[1], C: col}
		figureOps = append(figureOps, figure)
		return &figure
	case "move":
		return &painter.MoveOp{X: intArgs[0], Y: intArgs[1], Figures: figureOps}
	case "reset":
		figureOps = figureOps[0:0]
		return painter.OperationFunc(painter.Reset)
	case "update":
		return painter.UpdateOp
	}
	return nil
}
