package conv

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func NumericalValue(value reflect.Value) (float64, bool) {
	switch value.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int()), true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(value.Uint()), true

	case reflect.Float32, reflect.Float64:
		return value.Float(), true

	default:
		return 0, false
	}
}

func NumericalToString(value any) (string, bool) {
	var val string

	switch value.(type) {
	default:
		return "0", false
	case int:
		intVal, _ := value.(int)
		val = strconv.FormatInt(int64(intVal), 10)
	case int8:
		intVal, _ := value.(int8)
		val = strconv.FormatInt(int64(intVal), 10)
	case int16:
		intVal, _ := value.(int16)
		val = strconv.FormatInt(int64(intVal), 10)
	case int32:
		intVal, _ := value.(int32)
		val = strconv.FormatInt(int64(intVal), 10)
	case int64:
		intVal, _ := value.(int64)
		val = strconv.FormatInt(intVal, 10)
	case uint:
		intVal, _ := value.(uint)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint8:
		intVal, _ := value.(uint8)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint16:
		intVal, _ := value.(uint16)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint32:
		intVal, _ := value.(uint32)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint64:
		intVal, _ := value.(uint64)
		val = strconv.FormatUint(intVal, 10)
	case float32:
		floatVal, _ := value.(float32)
		val = strconv.FormatFloat(float64(floatVal), 'f', -1, 32)
	case float64:
		floatVal, _ := value.(float64)
		val = strconv.FormatFloat(floatVal, 'f', -1, 64)
	case string:
		val = value.(string)
	}
	return val, true
}

func LessValue(a, b reflect.Value) bool {
	aValue, aNumerical := NumericalValue(a)
	bValue, bNumerical := NumericalValue(b)

	if aNumerical && bNumerical {
		return aValue < bValue
	}

	if !aNumerical && !bNumerical {
		// In theory this should mean they are both strings. In reality,
		// they could be any other type and the String() representation
		// will be something like "<bool>" if it is not a string. Since
		// distinct values of non-strings still return the same value
		// here that's what makes the ordering undefined.
		return strings.Compare(a.String(), b.String()) < 0
	}

	// Numerical values are always treated as less than other types
	// (including strings that might represent numbers themselves). The
	// inverse is also true.
	return aNumerical && !bNumerical
}

func JsonNumberToInt64(number json.Number) int64 {
	v, _ := number.Int64()
	return v
}

func JsonNumberToFloat64(number json.Number) float64 {
	v, _ := number.Float64()
	return v
}
