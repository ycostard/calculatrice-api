package services

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
)

const maxHistory = 20

func EvaluateExpression(expression string) (float64, error) {
	tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, expression)
	if err != nil {
		return 0, fmt.Errorf("invalid expression: %v", err)
	}

	value, _ := constant.Float64Val(tv.Value)

	history, err := LoadHistory()
	if err != nil {
		return 0, err
	}

	history = append(history, HistoryEntry{Expression: expression, Result: value})

	if len(history) > maxHistory {
		history = history[len(history)-maxHistory:]
	}

	if err := SaveHistory(history); err != nil {
		return 0, err
	}

	return value, nil
}
