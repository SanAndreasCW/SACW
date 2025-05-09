package commons

import (
	"fmt"
	"strconv"
)

func IntToString[T int | int32 | int64 | int16 | int8 | uint | uint32 | uint64 | uint16 | uint8](i T) string {
	return fmt.Sprintf("%d", i)
}

func FloatToString[T float32 | float64](f T) string {
	return fmt.Sprintf("%f", f)
}

func StringToInt[T ~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8](s *string) (T, error) {
	parsed, err := strconv.Atoi(*s)
	if err != nil {
		return T(0), err
	}
	return T(parsed), nil
}

func StringToFloat[T ~float32 | ~float64](s *string) (T, error) {
	parsed, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return T(0), err
	}
	return T(parsed), nil
}
