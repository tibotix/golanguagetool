package text_processor

type PredicateFunc func(rs []rune) bool
type Predicate struct {
	test        PredicateFunc
	contextSize int
	consuming   bool
}

func ConsumingPredicate(predicate PredicateFunc, contextSize int) Predicate {
	return Predicate{
		test:        predicate,
		contextSize: contextSize,
		consuming:   true,
	}
}
func NonConsumingPredicate(predicate PredicateFunc, contextSize int) Predicate {
	return Predicate{
		test:        predicate,
		contextSize: contextSize,
		consuming:   false,
	}
}

func All(r rune) PredicateFunc {
	return func(rs []rune) bool {
		for _, v := range rs {
			if v != r {
				return false
			}
		}
		return true
	}
}

func IsRune(r rune) PredicateFunc {
	return func(rs []rune) bool {
		return rs[0] == r
	}
}
func IsRunes(r ...rune) PredicateFunc {
	return func(rs []rune) bool {
		if len(r) != len(rs) {
			return false
		}
		for i := 0; i < len(rs); i++ {
			if r[i] != rs[i] {
				return false
			}
		}
		return true
	}
}

func PredicateTrue(rs []rune) bool {
	return true
}

var ConsumingPredicateTrue = ConsumingPredicate(PredicateTrue, 1)
var NonConsumingPredicateTrue = NonConsumingPredicate(PredicateTrue, 1)
