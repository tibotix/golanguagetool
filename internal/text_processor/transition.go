package text_processor

// For every step, the TP does the following for every transition
// in the current state:
// 1. if the predicate returns true, this transition is taken and it resumes with step 2.
// 2. if the predicate was a ConsumingPredicate, the read head of the input band is advanced the contextSize of the predicate
// 3. if the predicate was a NonConsumingPredicate, the read head of the input band stays in place.
// 4. the result of output, which receives the same arguments as predicate, is written to the input band.
// 5. the band actions are executed, enabling the possibility for customized band manipulations.
type Transition[T ITP] struct {
	predicate   Predicate
	targetState int
	output      Output
	bandActions BandActions[T]
}

// Creates a new Transition with default values
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
