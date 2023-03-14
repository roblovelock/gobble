package utils

func NewLookupMap[T comparable](values []T) map[T]bool {
	lookup := make(map[T]bool, len(values))
	for _, v := range values {
		lookup[v] = true
	}
	return lookup
}

func NewByteLookupArray(bytes []byte) [256]bool {
	lookup := [256]bool{}
	for _, b := range bytes {
		lookup[b] = true
	}
	return lookup
}

func NewInverseByteLookupArray(bytes []byte) [256]bool {
	lookup := NewByteLookupArray(bytes)
	for i := range lookup {
		lookup[i] = !lookup[i]
	}
	return lookup
}
