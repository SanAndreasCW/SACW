package commons

import (
	"fmt"
	"strconv"
)

func IntToString[T int | int32 | int64 | int16 | int8](i T) string {
	return fmt.Sprintf("%d", i)
}

func StringToInt[T ~int | ~int64 | ~int32 | ~int16 | ~int8](s *string) (T, error) {
	parsed, err := strconv.Atoi(*s)
	if err != nil {
		panic(err)
		return T(0), err
	}
	return T(parsed), nil
}
