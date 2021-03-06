package convert

import (
	"fmt"
	"strconv"
)

func ToInt(i interface{}) (int, error) {
	switch v := i.(type) {

	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case int:
		return int(v), nil
	case float64:
		return int(v), nil
	case float32:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("I don't know about type %T!", v)
	}
}

func ToString(i interface{}) (string, error) {
	switch v := i.(type) {

	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil

	default:
		return "", fmt.Errorf("I don't know about type %T!", v)
	}
}

func ToFloat(i interface{}) (float64, error) {
	switch v := i.(type) {

	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("I don't know about type %T!", v)
	}
}
