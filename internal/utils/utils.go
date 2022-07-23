package utils

func IsSet[T uint8 | uint16](data, mask T) bool {
	return (data & mask) != 0
}

func Min[T int](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}
