package util

type refernce[T any] interface {
	*T
}

func MapRef[I any, PI refernce[I], O any, PO refernce[O]](inp []I, fn func(PI) PO) []O {
	res := make([]O, len(inp))
	for i, item := range inp {
		res[i] = *fn(&item)
	}
	return res
}

func Map[I any, O any](inp []I, fn func(I) O) []O {
	res := make([]O, len(inp))
	for i, item := range inp {
		res[i] = fn(item)
	}
	return res
}

func AsRef[T any](inp T) *T {
	return &inp
}

func DerefOrDefault[I any](inp *I) I {
	var res I
	if inp != nil {
		res = *inp
	}
	return res
}
