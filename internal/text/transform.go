package text

import (
	"fmt"
	"unicode"

	tp "github.com/tibotix/golanguagetool/internal/text_processor"
)

func incrementRuneCountBy(n int) tp.BandActions[*tp.ThreeTP[int, int]] {
	return func(t *tp.ThreeTP[int, int]) {
		t.Band2.Overwrite(t.Band2.Peek() + n)
	}
}

var incrementRuneCount = incrementRuneCountBy(1)

func addLineBeginning(t *tp.ThreeTP[int, int]) {
	t.Band3.Write(t.Band2.Peek())
}

type TextTransformer func(text []byte) (*TransformResult, error)

type TransformResult struct {
	Data           []byte
	LineBeginnings LineBeginnings
}

func nlTransition(targetState int) tp.Transition[*tp.ThreeTP[int, int]] {
	return *tp.NewTransition[*tp.ThreeTP[int, int]]().
		WithPredicate(tp.ConsumingPredicate(tp.IsRune(nlRune), 1)).
		WithTargetState(targetState).
		WithOutput(tp.OutputEcho).
		WithBandActions(tp.ChainActions(incrementRuneCount, addLineBeginning))
}

func echoTransition(targetState int) tp.Transition[*tp.ThreeTP[int, int]] {
	return *tp.NewTransition[*tp.ThreeTP[int, int]]().
		WithPredicate(tp.ConsumingPredicate(tp.PredicateTrue, 1)).
		WithTargetState(targetState).
		WithOutput(tp.OutputEcho).
		WithBandActions(incrementRuneCount)
}
func tabReplaceTransition(targetState int) tp.Transition[*tp.ThreeTP[int, int]] {
	return *tp.NewTransition[*tp.ThreeTP[int, int]]().
		WithPredicate(tp.ConsumingPredicate(tp.IsRune(tabRune), 1)).
		WithTargetState(targetState).
		WithOutput(tp.OutputRunes(spaceRune)).
		WithBandActions(incrementRuneCount)
}

func allReplaceTransition(targetState int) tp.Transition[*tp.ThreeTP[int, int]] {
	return *tp.NewTransition[*tp.ThreeTP[int, int]]().
		WithPredicate(tp.ConsumingPredicate(tp.PredicateTrue, 1)).
		WithTargetState(targetState).
		WithOutput(tp.OutputNothing).
		WithBandActions(incrementRuneCount)
}

func filterTransition(targetState int) tp.Transition[*tp.ThreeTP[int, int]] {
	return *tp.NewTransition[*tp.ThreeTP[int, int]]().
		WithPredicate(tp.ConsumingPredicate(func(rs []rune) bool {
			r := rs[0]
			return r == crRune || (!unicode.IsGraphic(r) && r != nlRune)
		}, 1)).
		WithTargetState(targetState).
		WithOutput(tp.OutputNothing).
		WithBandActions(tp.NoBandAction[*tp.ThreeTP[int, int]])
}

func codeBlockTransition(targetState int) tp.Transition[*tp.ThreeTP[int, int]] {
	return *tp.NewTransition[*tp.ThreeTP[int, int]]().
		WithPredicate(tp.ConsumingPredicate(tp.IsRunes(nlRune, backtickRune, backtickRune, backtickRune), 4)).
		WithTargetState(targetState).
		WithOutput(tp.OutputRunes(nlRune)).
		WithBandActions(tp.ChainActions(incrementRuneCount, addLineBeginning))
}

var mdTransitions [][]tp.Transition[*tp.ThreeTP[int, int]] = [][]tp.Transition[*tp.ThreeTP[int, int]]{
	{ // State 0
		filterTransition(0),
		tabReplaceTransition(0),
		codeBlockTransition(1),
		nlTransition(0),
		echoTransition(0),
	},
	{ // State 1
		codeBlockTransition(0),
		nlTransition(1), // stay here on newline
		allReplaceTransition(1),
	},
}
var plainTransitions [][]tp.Transition[*tp.ThreeTP[int, int]] = [][]tp.Transition[*tp.ThreeTP[int, int]]{
	{ // State 0
		filterTransition(0),
		tabReplaceTransition(0),
		nlTransition(0),
		echoTransition(0),
	},
}

func TransformTextMarkdown(text []byte) (*TransformResult, error) {
	return transformText(text, mdTransitions)
}

func TransformTextPlain(text []byte) (*TransformResult, error) {
	return transformText(text, plainTransitions)
}

func transformText(text []byte, transitions [][]tp.Transition[*tp.ThreeTP[int, int]]) (*TransformResult, error) {
	tm := tp.ThreeTP[int, int]{
		TextProcessor: *tp.NewTextProcessorWithTransitions(text, transitions),
		Band2:         tp.NewBandWithCapacity[int](1),
		Band3:         tp.NewBandWithCapacity[int](50),
	}

	err := tm.Run()
	if err != nil {
		return nil, err
	}

	fmt.Printf("LineBreaks: %v\n", tm.Band3.AllWritten())
	return &TransformResult{
		Data:           tm.InputBand.AllWritten(),
		LineBeginnings: NewLineBeginningsFromArray(tm.Band3.AllWritten()),
	}, nil
}
