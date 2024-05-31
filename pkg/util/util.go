package util

func PtrOfCopy[T any](v T) *T {
	return &v
}
