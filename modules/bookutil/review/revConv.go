package review

import (
	"errors"
)

var err = errors.New("Invalid comment")

func Perfreview[Rev string | float64](i ...Rev) float64 {

}

func ReviewConv(str string) (float64, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Fair":
		return 4, nil
	case "Poor":
		return 3, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		panic(err)
	}
}
