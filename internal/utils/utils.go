package utils

func IsSet(data, mask uint16) bool {
	return (data & mask) != 0
}
