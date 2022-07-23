package utils

func IsSet[T uint8 | uint16](data, mask T) bool {
	return (data & mask) != 0
}
