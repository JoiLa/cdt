package cdt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// isPtr is_prt
func isPtr(refVal reflect.Value) bool {
	return refVal.Kind() == reflect.Ptr
}

// isInterface is_interface
func isInterface(refVal reflect.Value) bool {
	return refVal.Kind() == reflect.Interface
}

// isString is_string
func isString(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isString(reflect.ValueOf(refVal.Elem().Interface()))
	}
	return refVal.Kind() == reflect.String
}

// isArray is_array
func isArray(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isArray(reflect.ValueOf(refVal.Elem().Interface()))
	}
	return refVal.Kind() == reflect.Slice || refVal.Kind() == reflect.Array
}

// isBytes is_bytes
func isBytes(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isBytes(reflect.ValueOf(refVal.Elem().Interface()))
	}
	return refVal.Kind() == reflect.Slice && refVal.Type().Elem().Kind() == reflect.Uint8
}

// isMap is_map
func isMap(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isMap(reflect.ValueOf(refVal.Elem().Interface()))
	}
	return refVal.Kind() == reflect.Map
}

// isStruct is_struct
func isStruct(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isStruct(reflect.ValueOf(refVal.Elem().Interface()))
	}
	return refVal.Kind() == reflect.Struct
}

// isBoolean is_bool
func isBoolean(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isBoolean(reflect.ValueOf(refVal.Elem().Interface()))
	}
	return refVal.Kind() == reflect.Bool

}

// isInt is_int
func isInt(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isInt(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

// isInt8 is_int8
func isInt8(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isInt8(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Int8:
		return true
	}
	return false
}

// isInt16 is_int16
func isInt16(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isInt16(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Int16:
		return true
	}
	return false
}

// isInt32 is_int32
func isInt32(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isInt32(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Int32:
		return true
	}
	return false
}

// isInt64 is_int64
func isInt64(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isInt64(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Int64:
		return true
	}
	return false
}

// isUint is_uint
func isUint(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isUint(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}
func isUint8(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isUint8(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uint8:
		return true
	}
	return false
}

// isUint16 is_uint16
func isUint16(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isUint16(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uint16:
		return true
	}
	return false
}

// isUint32 is_uint32
func isUint32(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isUint32(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uint32:
		return true
	}
	return false
}

// isUint64 is_uint64
func isUint64(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isUint64(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uint64:
		return true
	}
	return false
}

// isUintptr is_uintptr
func isUintptr(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isUintptr(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Uintptr:
		return true
	}
	return false
}

// isFloat is_float
func isFloat(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isFloat(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// isFloat32 is_float32
func isFloat32(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isFloat(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Float32:
		return true
	}
	return false
}

// isFloat64 is_float64
func isFloat64(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isFloat(reflect.ValueOf(refVal.Elem().Interface()))
	}
	switch refVal.Kind() {
	case reflect.Float64:
		return true
	}
	return false
}

// isNil is_nil
func isNil(refVal reflect.Value) bool {
	return refVal.Kind() == reflect.Invalid
}

// isTime is_time
func isTime(refVal reflect.Value) bool {
	if isPtr(refVal) {
		return isTime(reflect.ValueOf(refVal.Elem().Interface()))
	}
	if refVal.Kind() == reflect.Struct {
		timeType := reflect.TypeOf(time.Time{})
		return refVal.Type().AssignableTo(timeType)
	}
	return false
}

// IsNumeric is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
// In PHP hexadecimal (e.g. 0xf4c3b00c) is not supported, but IsNumeric is supported.
func isNumeric(refVal reflect.Value) bool {
	switch refVal.Kind() {
	case reflect.Ptr:
		return isNumeric(reflect.ValueOf(refVal.Elem().Interface()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		return true
	case reflect.Slice, reflect.String:
		var str string
		switch refVal.Kind() {
		case reflect.Slice:
			if isBytes(refVal) {
				str = string(refVal.Bytes())
			}
			break
		case reflect.String:
			str = refVal.String()
			break
		}
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9, Point, Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}
	return false
}

// toFloat64 function to convert other type data to float64
func toFloat64(refVal reflect.Value) (float64, error) {
	switch refVal.Kind() {
	case reflect.Invalid:
		return 0, nil
	case reflect.Ptr:
		return toFloat64(reflect.ValueOf(refVal.Elem().Interface()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(refVal.Uint()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(refVal.Int()), nil
	case reflect.Float32, reflect.Float64:
		return refVal.Float(), nil
	case reflect.String:
		if isNumeric(refVal) {
			return json.Number(strings.TrimSpace(refVal.String())).Float64()
		} else {
			return 0, newError(refVal)
		}
	case reflect.Slice:
		if isBytes(refVal) && isNumeric(refVal) {
			return json.Number(refVal.Bytes()).Float64()
		} else {
			return 0, newError(refVal)
		}
	case reflect.Bool:
		if refVal.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, newError(refVal)
	}
}

// toBoolean function to convert other type data to boolean
func toBoolean(refVal reflect.Value) (bool, error) {
	var val any
	switch refVal.Kind() {
	case reflect.Invalid:
		return false, nil
	case reflect.Ptr:
		return toBoolean(reflect.ValueOf(refVal.Elem().Interface()))
	case reflect.String:
		// val to lower
		val = strings.ToLower(refVal.String())
		break
	case reflect.Slice:
		if isBytes(refVal) {
			val = strings.ToLower(string(refVal.Bytes()))
			break
		}
		return false, newError(refVal)
	default:
		val = refVal.Interface()
		if isInt(refVal) {
			valInt, _ := toFloat64(refVal)
			val = int(valInt)
		}
	}
	switch val {
	case nil:
		return false, nil
	case "1", "t", "true", 1, true:
		return true, nil
	case "0", "f", "false", 0, false:
		return false, nil
	default:
		return false, newError(refVal)
	}
}

// toString function to convert other type data to string
func toString(refVal reflect.Value) (string, error) {
	switch refVal.Kind() {
	case reflect.Invalid:
		return "", nil
	case reflect.Ptr:
		return toString(reflect.ValueOf(refVal.Elem().Interface()))
	case reflect.String:
		return refVal.String(), nil
	case reflect.Slice:
		if isBytes(refVal) {
			return string(refVal.Bytes()), nil
		}
		b, err := json.Marshal(refVal.Interface())
		if err != nil {
			return "", err
		}
		return string(b), nil
	case reflect.Bool:
		return fmt.Sprintf("%v", refVal.Interface()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", refVal.Interface()), nil
	case reflect.Map, reflect.Struct, reflect.Interface, reflect.Array:
		b, err := json.Marshal(refVal.Interface())
		if err != nil {
			return "", err
		}
		return string(b), nil
	default:
		return "", newError(refVal)
	}
}

// toTime function to convert other type data to time
func toTime(refVal reflect.Value) (time.Time, error) {
	if !isTime(refVal) {
		return time.Time{}, newError(refVal)
	}
	if isPtr(refVal) {
		return toTime(reflect.ValueOf(refVal.Elem().Interface()))
	}
	var val any
	val = refVal.Interface()
	switch val.(type) {
	case time.Time:
		return val.(time.Time), nil
	default:
		return time.Time{}, newError(refVal)
	}
}

// toArray function to convert other type data to array
func toArray(refVal reflect.Value) ([]any, error) {
	switch refVal.Kind() {
	case reflect.Invalid:
		return nil, nil
	case reflect.Ptr:
		return toArray(reflect.ValueOf(refVal.Elem().Interface()))
	case reflect.Slice:
		if isBytes(refVal) {
			return toArray(reflect.ValueOf(string(refVal.Bytes())))
		}
		return nil, newError(refVal)
	case reflect.Array:
		var arr []any
		for i := 0; i < refVal.Len(); i++ {
			currentIndex := refVal.Index(i)
			if currentIndex.CanInterface() {
				arr = append(arr, currentIndex.Interface())
			} else {
				arr = append(arr, nil)
			}
		}
		return arr, nil
	case reflect.String:
		var arr []any
		if err := json.Unmarshal([]byte(refVal.String()), &arr); err != nil {
			return nil, err
		}
		return arr, nil
	}
	return nil, newError(refVal)
}

// toMap function to convert other type data to map
func toMap(refVal reflect.Value) (map[string]any, error) {
	switch refVal.Kind() {
	case reflect.Invalid:
		return nil, nil
	case reflect.Ptr:
		return toMap(reflect.ValueOf(refVal.Elem().Interface()))
	case reflect.Slice:
		if isBytes(refVal) {
			return toMap(reflect.ValueOf(string(refVal.Bytes())))
		}
		return nil, newError(refVal)
	case reflect.Struct:
		var m = make(map[string]any)
		numFieldLen := refVal.NumField()
		for i := 0; i < numFieldLen; i++ {
			currentFieldIndex := refVal.Field(i)
			if currentFieldIndex.CanInterface() {
				m[refVal.Type().Field(i).Name] = currentFieldIndex.Interface()
			} else {
				m[refVal.Type().Field(i).Name] = nil
			}
		}
		return m, nil
	case reflect.Map:
		var m = make(map[string]any)
		mapKeys := refVal.MapKeys()
		for _, key := range mapKeys {
			currentMapIndex := refVal.MapIndex(key)
			if currentMapIndex.CanInterface() {
				m[key.String()] = currentMapIndex.Interface()
			} else {
				m[key.String()] = nil
			}
		}
		return m, nil
	case reflect.String:
		var m = make(map[string]any)
		if err := json.Unmarshal([]byte(refVal.String()), &m); err != nil {
			return nil, err
		}
		return m, nil
	}
	return nil, newError(refVal)
}

// newError function to create new error
func newError(refVal reflect.Value) error {
	return fmt.Errorf("unable to casting type %T", refVal.Interface())
}
