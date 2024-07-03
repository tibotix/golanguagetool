package text_processor

type BandActions[T ITP] func(T)

func NoBandAction[T ITP](T) {}

func ChainActions[T ITP](actions ...BandActions[T]) BandActions[T] {
	return func(t T) {
		for _, a := range actions {
			a(t)
		}
	}
}
