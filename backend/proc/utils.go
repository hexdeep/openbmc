package proc

func Idx[T any](slice []T, idx int) T {
	if len(slice) <= idx {
		return *new(T)
	} else {
		return slice[idx]
	}
}
