package text_processor

type Transition[T ITP] struct {
	predicate   Predicate
	targetState int
	output      Output
	bandActions BandActions[T]
}

func NewTransition[T ITP]() *Transition[T] {
	return &Transition[T]{
		predicate:   ConsumingPredicateTrue,
		targetState: 0,
		output:      OutputNothing,
		bandActions: NoBandAction[T],
	}
}
func (t *Transition[T]) WithPredicate(predicate Predicate) *Transition[T] {
	t.predicate = predicate
	return t
}
func (t *Transition[T]) WithTargetState(targetState int) *Transition[T] {
	t.targetState = targetState
	return t
}
func (t *Transition[T]) WithOutput(output Output) *Transition[T] {
	t.output = output
	return t
}
func (t *Transition[T]) WithBandActions(bandActions BandActions[T]) *Transition[T] {
	t.bandActions = bandActions
	return t
}
