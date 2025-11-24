package utils

import "strconv"

func IntToString[T int32 | int64 | int](n T) string {
	return strconv.FormatInt(int64(n), 10)
}
