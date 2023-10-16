package walk

type Walker[T any] struct {
	values  []T
	pointer int
}

func NewWalker[T any](values []T) *Walker[T] {
	return &Walker[T]{
		values: values,
	}
}

func (w *Walker[T]) Next() {
	w.pointer++
}

func (w *Walker[T]) value(index int) (t T) {
	if w.pointer < len(w.values) {
		t = w.values[index]
	}

	return t
}

func (w *Walker[T]) Current() (t T) {
	return w.value(w.pointer)
}

func (w *Walker[T]) Lookahead() (t T) {
	return w.value(w.pointer + 1)
}

func (w *Walker[T]) LookaheadN(n int) (t T) {
	return w.value(w.pointer + n)
}
