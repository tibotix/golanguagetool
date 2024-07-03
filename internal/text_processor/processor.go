package text_processor

import (
	"bytes"
	"errors"
	"unicode/utf8"
)

type ITP interface {
	Run() error
}

// This is an LBA (Linear-Bounded-Automaton).
// It is a normed TextProcessor M(Q,Σ,Γ,Δ,s,F).
// This LBA does only allow Instructions, which move the head to the right.
// It stops when it reaches the end of the input.
// Additionally it has two head pointers for one band, namely readHead and writeHead.
// Σ = [255]
// Γ = [INT_MAX] ∪ {□}
// s = 0
// For simplicity and performance reasons, the Predicates however receive a full rune as input.
type TextProcessor[T ITP] struct {
	currentState int
	InputBand    Band[byte]
	transitions  [][]Transition[T]
}

func NewTextProcessor[T ITP](text []byte) *TextProcessor[T] {
	return &TextProcessor[T]{
		currentState: 0,
		InputBand:    NewBandWithData(text),
		transitions:  [][]Transition[T]{},
	}
}
func NewTextProcessorWithTransitions[T ITP](text []byte, transitions [][]Transition[T]) *TextProcessor[T] {
	return &TextProcessor[T]{
		currentState: 0,
		InputBand:    NewBandWithData(text),
		transitions:  transitions,
	}
}

func (p *TextProcessor[T]) run(t T) error {
	runeBytes := make([]byte, utf8.UTFMax)
	outputRunesBytes := make([]byte, 0, 20*utf8.UTFMax)
	contextRunes := make([]rune, 0, 20)
	contextRunesWidth := 0
ReadLoop:
	for p.InputBand.readHead < len(p.InputBand.data) {
		outputRunesBytes = outputRunesBytes[:0]
		contextRunes = contextRunes[:0]
		contextRunesWidth = 0
	TransitionLoop:
		for _, transition := range p.transitions[p.currentState] {
			for len(contextRunes) < transition.predicate.contextSize {
				r, width := utf8.DecodeRune(p.InputBand.PeekAheadUptoN(contextRunesWidth, utf8.UTFMax))
				if width == 0 {
					continue TransitionLoop
				}
				if r == utf8.RuneError {
					return errors.New("invalid UTF8 encoding in input file")
				}
				contextRunes = append(contextRunes, r)
				contextRunesWidth += width
			}

			if transition.predicate.test(contextRunes[:transition.predicate.contextSize]) {
				p.currentState = transition.targetState
				output := transition.output(contextRunes[:transition.predicate.contextSize])
				for _, r := range output {
					width := utf8.EncodeRune(runeBytes, r)
					outputRunesBytes = append(outputRunesBytes, runeBytes[:width]...)
				}
				if len(outputRunesBytes) > 0 && (!p.InputBand.areHeadsTogether() || !bytes.Equal(p.InputBand.PeekN(len(outputRunesBytes)), outputRunesBytes)) {
					p.InputBand.WriteN(outputRunesBytes)
				} else {
					p.InputBand.MoveWriteHead(len(outputRunesBytes))
				}
				if transition.predicate.consuming {
					n := 0
					for _, r := range contextRunes[:transition.predicate.contextSize] {
						n += utf8.RuneLen(r)
					}
					p.InputBand.MoveReadHead(n)

				}
				transition.bandActions(t)
				continue ReadLoop
			}
		}
		return errors.New("no matching Transition found")
	}
	return nil
}

// Some variants with multiple Bands
type ThreeTP[B2 any, B3 any] struct {
	TextProcessor[*ThreeTP[B2, B3]]

	Band2 Band[B2]
	Band3 Band[B3]
}

func (t *ThreeTP[B1, B2]) Run() error {
	return t.run(t)
}
