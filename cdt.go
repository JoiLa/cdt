package cdt

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type convert struct {
	ctdStructTagKV map[string]string
	OriginVal      any
}

type DataRaw struct {
	convert
}

// NewConvert function to create new convert instance
func NewConvert(originVal any) *DataRaw {
	t := new(DataRaw)
	t.OriginVal = originVal
	return t
}

// Set function to set origin data
func (i *convert) Set(val any) {
	i.OriginVal = val
}

// SetCtdStructTagKV function to set ctd struct tag kv
func (i *convert) SetCtdStructTagKV(kv map[string]string) {
	i.ctdStructTagKV = kv
}

// GetOriginVal return origin value
func (i *convert) GetOriginVal() any {
	return i.OriginVal
}

// GetOriginValRef return origin value reference
func (i *convert) GetOriginValRef() reflect.Value {
	return reflect.ValueOf(i.OriginVal)
}

// IsNumeric is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
// In PHP hexadecimal (e.g. 0xf4c3b00c) is not supported, but IsNumeric is supported.
func (i *convert) IsNumeric() bool {
	return isNumeric(i.GetOriginValRef())
}

// IsPtr is_prt
func (i *convert) IsPtr() bool {
	return isPtr(i.GetOriginValRef())
}

// IsString is_string
func (i *convert) IsString() bool {
	return isString(i.GetOriginValRef())
}

// IsArray is_array
func (i *convert) IsArray() bool {
	return isArray(i.GetOriginValRef())
}

// IsBytes is_bytes
func (i *convert) IsBytes() bool {
	return isBytes(i.GetOriginValRef())
}

// IsMap is_map
func (i *convert) IsMap() bool {
	return isMap(i.GetOriginValRef())
}

// IsStruct is_struct
func (i *convert) IsStruct() bool {
	return isStruct(i.GetOriginValRef())
}

// IsBoolean is_bool
func (i *convert) IsBoolean() bool {
	return isBoolean(i.GetOriginValRef())
}

// IsInt is_int
func (i *convert) IsInt() bool {
	return isInt(i.GetOriginValRef())
}

// IsInt8 is_int8
func (i *convert) IsInt8() bool {
	return isInt8(i.GetOriginValRef())
}

// IsInt16 is_int16
func (i *convert) IsInt16() bool {
	return isInt16(i.GetOriginValRef())
}

// IsInt32 is_int32
func (i *convert) IsInt32() bool {
	return isInt32(i.GetOriginValRef())
}

// IsInt64 is_int64
func (i *convert) IsInt64() bool {
	return isInt64(i.GetOriginValRef())
}

// IsUint is_uint
func (i *convert) IsUint() bool {
	return isUint(i.GetOriginValRef())
}

// IsUint8 is_uint8
func (i *convert) IsUint8() bool {
	return isUint8(i.GetOriginValRef())
}

// IsUint16 is_uint16
func (i *convert) IsUint16() bool {
	return isUint16(i.GetOriginValRef())
}

// IsUint32 is_uint32
func (i *convert) IsUint32() bool {
	return isUint32(i.GetOriginValRef())
}

// IsUint64 is_uint64
func (i *convert) IsUint64() bool {
	return isUint64(i.GetOriginValRef())
}

// IsUintptr is_uintptr
func (i *convert) IsUintptr() bool {
	return isUintptr(i.GetOriginValRef())
}

// IsFloat is_float
func (i *convert) IsFloat() bool {
	return isFloat(i.GetOriginValRef())
}

// IsFloat32 is_float32
func (i *convert) IsFloat32() bool {
	return isFloat32(i.GetOriginValRef())
}

// IsFloat64 is_float64
func (i *convert) IsFloat64() bool {
	return isFloat64(i.GetOriginValRef())
}

// IsNil is_null
func (i *convert) IsNil() bool {
	return isNil(i.GetOriginValRef())
}

// IsTime is_time
func (i *convert) IsTime() bool {
	return isTime(i.GetOriginValRef())
}

// CopyVariable copy variable
func CopyVariable(data any) (newData any) {
	defer func() {
		if err := recover(); err != nil {
			newData = data
		}
	}()
	refVal := reflect.ValueOf(data)
	if refVal.Kind() == reflect.Invalid {
		newData = data
	} else {
		// copy data to newData
		copyRef := reflect.New(reflect.TypeOf(data).Elem())
		copyRef.Elem().Set(reflect.ValueOf(data).Elem())
		newData = copyRef.Interface()
	}
	return
}

// ctdStructTagKV struct tag option
type ctdStructTagKV map[string]string

// =====================================================
const ctdStructTag = "cdt"
const ctdStructTagParamSeparator = "="
const ctdStructTagParamInterval = ","

// example: `cdt:"type=string,time_format='2006-01-02 15:04:05'"`
// example: `cdt:"type=string,time_format=2006-01-02 15:04:05"`
// example: `cdt:"type=string,time_format=\"2006-01-02 15:04:05\""`
// example: `cdt:"type=string,time_format=` if input `time_format` tag ,but not set time_format value or `time_format` check is empty ,than set time_format `2006-01-02 15:04:05`
// example: `cdt:"type=string"`
// example: `cdt:"type=int"`
// example: `cdt:"type=float"`

// =====================================================
const ctdStructTagParamNameType = "type"
const ctdStructTagParamNameTimeFormat = "time_format"

// =====================================================
const ctdStructTagParamNameTypeString = "string"
const ctdStructTagParamNameTypeUInt = "uint"
const ctdStructTagParamNameTypeUInt8 = "uint8"
const ctdStructTagParamNameTypeUInt16 = "uint16"
const ctdStructTagParamNameTypeUInt32 = "uint32"
const ctdStructTagParamNameTypeUInt64 = "uint64"
const ctdStructTagParamNameTypeUIntPtr = "uintptr"
const ctdStructTagParamNameTypeInt = "int"
const ctdStructTagParamNameTypeInt8 = "int8"
const ctdStructTagParamNameTypeInt16 = "int16"
const ctdStructTagParamNameTypeInt32 = "int32"
const ctdStructTagParamNameTypeInt64 = "int64"
const ctdStructTagParamNameTypeFloat = "float"
const ctdStructTagParamNameTypeFloat32 = "float32"
const ctdStructTagParamNameTypeFloat64 = "float64"
const ctdStructTagParamNameTypeBoolean = "boolean"
const ctdStructTagParamNameTypeNull = "null"
const ctdStructTagParamNameTypeArray = "array"
const ctdStructTagParamNameTypeObject = "object"
const ctdStructTagParamNameTypeMap = "map"
const ctdStructTagParamNameTypeTime = "time"

// =====================================================
const ctdStructTagParamValueDefaultFormat = "2006-01-02 15:04:05"

// EncodeDataForTags encode struct data type value, with struct tags `cdt`
func EncodeDataForTags(v any) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("EncodeDataForTags:", err)
		}
	}()
	if v == nil {
		return
	}
	var value reflect.Value
	var typ reflect.Type
	switch v.(type) {
	case reflect.Value:
		value = v.(reflect.Value)
		typ = value.Type()
		break
	default:
		value = reflect.ValueOf(v)
		typ = value.Type()
		break
	}
	if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Interface {
		EncodeDataForTags(value.Elem())
		return
	}
	if typ.Kind() != reflect.Struct || value.Interface() == nil {
		return
	}
	switch value.Interface().(type) {
	case DataRaw, *DataRaw:
		if !value.CanAddr() {
			return
		}
		sourceDataRawPtr := value.Addr().Interface().(*DataRaw)
		if len(sourceDataRawPtr.ctdStructTagKV) <= 0 {
			return
		}
		encodeDataForCdtTagKV(value, sourceDataRawPtr.ctdStructTagKV)
		return
	}
	//
	for i := 0; i < typ.NumField(); i++ {
		currentFieldValueRef := value.Field(i)
		if isPtr(currentFieldValueRef) || isInterface(currentFieldValueRef) {
			for isPtr(currentFieldValueRef) || isInterface(currentFieldValueRef) {
				currentFieldValueRef = currentFieldValueRef.Elem()
			}
		}
		currentFieldValue := currentFieldValueRef.Interface()
		sourceDataRawPtr := new(DataRaw)
		switch currentFieldValue.(type) {
		case DataRaw:
			if !currentFieldValueRef.CanAddr() {
				continue
			}
			sourceDataRawPtr = currentFieldValueRef.Addr().Interface().(*DataRaw)
			break
		default:
			if currentFieldValueRef.Kind() == reflect.Struct {
				// other struct
				EncodeDataForTags(currentFieldValueRef)
				return
			}
			continue
		}
		var configTagKV ctdStructTagKV
		// check has set custom tag
		if len(sourceDataRawPtr.ctdStructTagKV) > 0 {
			// with custom tag
			configTagKV = sourceDataRawPtr.ctdStructTagKV
		} else {
			// no custom tag, that parse struct tag
			configTagKV = parseConfigCdtTagToKV(typ.Field(i))
		}
		//
		if len(configTagKV) == 0 {
			// no tag options
			continue
		}
		// each tag kv
		encodeDataForCdtTagKV(currentFieldValueRef, configTagKV)
	}
}

// DecodeDataForTags decode struct data type value, with struct tags `cdt`
func DecodeDataForTags(v any) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("DecodeDataForTags:", err)
		}
	}()
	if v == nil {
		return
	}
	var value reflect.Value
	var typ reflect.Type
	switch v.(type) {
	case reflect.Value:
		value = v.(reflect.Value)
		typ = value.Type()
		break
	default:
		value = reflect.ValueOf(v)
		typ = value.Type()
		break
	}
	if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Interface {
		DecodeDataForTags(value.Elem())
		return
	}
	if typ.Kind() != reflect.Struct || value.Interface() == nil {
		return
	}
	switch value.Interface().(type) {
	case DataRaw, *DataRaw:
		if !value.CanAddr() {
			return
		}
		sourceDataRawPtr := value.Addr().Interface().(*DataRaw)
		if len(sourceDataRawPtr.ctdStructTagKV) <= 0 {
			return
		}
		decodeDataForCdtTagKV(value, sourceDataRawPtr.ctdStructTagKV)
		return
	}
	for i := 0; i < typ.NumField(); i++ {
		currentFieldValueRef := value.Field(i)
		if isPtr(currentFieldValueRef) || isInterface(currentFieldValueRef) {
			for isPtr(currentFieldValueRef) || isInterface(currentFieldValueRef) {
				currentFieldValueRef = currentFieldValueRef.Elem()
			}
		}
		currentFieldValue := currentFieldValueRef.Interface()
		sourceDataRawPtr := new(DataRaw)
		switch currentFieldValue.(type) {
		case DataRaw:
			if !currentFieldValueRef.CanAddr() {
				return
			}
			sourceDataRawPtr = currentFieldValueRef.Addr().Interface().(*DataRaw)
			break
		default:
			if currentFieldValueRef.Kind() == reflect.Struct {
				// other struct
				DecodeDataForTags(currentFieldValueRef)
				return
			}
			continue
		}
		var configTagKV ctdStructTagKV
		// check has set custom tag
		if len(sourceDataRawPtr.ctdStructTagKV) > 0 {
			// with custom tag
			configTagKV = sourceDataRawPtr.ctdStructTagKV
		} else {
			// no custom tag, that parse struct tag
			configTagKV = parseConfigCdtTagToKV(typ.Field(i))
		}
		//
		if len(configTagKV) == 0 {
			// no tag options
			continue
		}
		// each tag kv
		decodeDataForCdtTagKV(currentFieldValueRef, configTagKV)
	}
}

// encodeDataForCdtTagKV encode data for tag kv
func encodeDataForCdtTagKV(currentFieldValueRef reflect.Value, configTagKV ctdStructTagKV) {
	if !currentFieldValueRef.CanAddr() {
		return
	}
	sourceDataRawPtr := currentFieldValueRef.Addr().Interface().(*DataRaw)
	for configTagKeyStr, configTagValueStr := range configTagKV {
		switch strings.ToLower(configTagKeyStr) {
		case ctdStructTagParamNameType:
			// has set type tag
			switch strings.ToLower(configTagValueStr) {
			case ctdStructTagParamNameTypeString:
				// data type is string
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsString() {
					continue
				}
				convertStrValue, convertStrValueErr := sourceDataRawPtr.ToStringE()
				if convertStrValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertStrValue)
				break
			case ctdStructTagParamNameTypeUInt, ctdStructTagParamNameTypeUInt8, ctdStructTagParamNameTypeUInt16, ctdStructTagParamNameTypeUInt32, ctdStructTagParamNameTypeUInt64, ctdStructTagParamNameTypeUIntPtr,
				ctdStructTagParamNameTypeInt, ctdStructTagParamNameTypeInt8, ctdStructTagParamNameTypeInt16, ctdStructTagParamNameTypeInt32, ctdStructTagParamNameTypeInt64:
				// data type is uint || int
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsInt() {
					continue
				}
				convertInt64Value, convertInt64ValueErr := sourceDataRawPtr.ToInt64E()
				if convertInt64ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertInt64Value)
				break
			case ctdStructTagParamNameTypeFloat, ctdStructTagParamNameTypeFloat32, ctdStructTagParamNameTypeFloat64:
				// data type is float
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsFloat() {
					continue
				}
				convertFloat64Value, convertFloat64ValueErr := sourceDataRawPtr.ToFloat64E()
				if convertFloat64ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertFloat64Value)
				break
			case ctdStructTagParamNameTypeBoolean:
				// data type is boolean
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsBoolean() {
					continue
				}
				convertBooleanValue, convertBooleanValueErr := sourceDataRawPtr.ToBooleanE()
				if convertBooleanValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertBooleanValue)
				break
			case ctdStructTagParamNameTypeNull:
				// data type is null
				if sourceDataRawPtr.IsNil() {
					continue
				}
				sourceDataRawPtr.Set(nil)
				break
			case ctdStructTagParamNameTypeTime:
				// data type is time
				// check sourceDataRawPtr data value is time.Time
				if sourceDataRawPtr.IsNil() {
					continue
				}
				if sourceDataRawPtr.IsString() || sourceDataRawPtr.IsBytes() {
					timeLayoutFormatStr, timeLayoutFormatStrHasExists := configTagKV[ctdStructTagParamNameTimeFormat]
					if !timeLayoutFormatStrHasExists {
						continue
					}
					if len(timeLayoutFormatStr) == 0 {
						timeLayoutFormatStr = ctdStructTagParamValueDefaultFormat
					}
					convertStrValue, convertStrValueErr := sourceDataRawPtr.ToStringE()
					if convertStrValueErr != nil {
						continue
					}
					parsedTimeFormat, parsedTimeFormatErr := time.Parse(timeLayoutFormatStr, convertStrValue)
					if parsedTimeFormatErr != nil {
						// parse time with default time format
						parsedTimeFormat, parsedTimeFormatErr = time.Parse(time.RFC3339Nano, convertStrValue)
					}
					if parsedTimeFormatErr != nil {
						continue
					}
					sourceDataRawPtr.Set(parsedTimeFormat.Format(timeLayoutFormatStr))
				}
				if !sourceDataRawPtr.IsTime() {
					continue
				}
				convertTimeValue, convertTimeValueErr := sourceDataRawPtr.ToTimeE()
				if convertTimeValueErr != nil {
					continue
				}
				timeLayoutFormatStr, timeLayoutFormatStrHasExists := configTagKV[ctdStructTagParamNameTimeFormat]
				if !timeLayoutFormatStrHasExists {
					continue
				}
				if len(timeLayoutFormatStr) == 0 {
					timeLayoutFormatStr = ctdStructTagParamValueDefaultFormat
				}
				sourceDataRawPtr.Set(convertTimeValue.Format(timeLayoutFormatStr))
				break
			case ctdStructTagParamNameTypeArray, ctdStructTagParamNameTypeObject, ctdStructTagParamNameTypeMap:
				// data type is array || object
				if sourceDataRawPtr.IsNil() {
					continue
				}
				if (sourceDataRawPtr.IsArray() || sourceDataRawPtr.IsStruct() || sourceDataRawPtr.IsMap()) && !sourceDataRawPtr.IsBytes() {
					continue
				}
				convertStrValue, convertStrValueErr := sourceDataRawPtr.ToStringE()
				if convertStrValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertStrValue)
				break
			}
			break
		}
	}
}

// decodeDataForCdtTagKV decode data for tag kv
func decodeDataForCdtTagKV(currentFieldValueRef reflect.Value, configTagKV ctdStructTagKV) {
	if !currentFieldValueRef.CanAddr() {
		return
	}
	sourceDataRawPtr := currentFieldValueRef.Addr().Interface().(*DataRaw)
	// each tag kv
	for configTagKeyStr, configTagValueStr := range configTagKV {
		switch strings.ToLower(configTagKeyStr) {
		case ctdStructTagParamNameType:
			// has set type tag
			switch strings.ToLower(configTagValueStr) {
			case ctdStructTagParamNameTypeString:
				// tag type is string, need convert data value to string
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsString() {
					continue
				}
				convertStrValue, convertStrValueErr := sourceDataRawPtr.ToStringE()
				if convertStrValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertStrValue)
				break
			case ctdStructTagParamNameTypeUInt:
				// tag type is uint, need convert data value to uint
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertUIntValue, convertUIntValueErr := sourceDataRawPtr.ToUintE()
				if convertUIntValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertUIntValue)
				break
			case ctdStructTagParamNameTypeUInt8:
				// tag type is uint8, need convert data value to uint8
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertUint8Value, convertUint8ValueErr := sourceDataRawPtr.ToUint8E()
				if convertUint8ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertUint8Value)
				break
			case ctdStructTagParamNameTypeUInt16:
				// tag type is uint16, need convert data value to uint16
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertUint16Value, convertUint16ValueErr := sourceDataRawPtr.ToUint16E()
				if convertUint16ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertUint16Value)
				break
			case ctdStructTagParamNameTypeUInt32:
				// tag type is uint32, need convert data value to uint32
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertUint32Value, convertUint32ValueErr := sourceDataRawPtr.ToUint32E()
				if convertUint32ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertUint32Value)
				break
			case ctdStructTagParamNameTypeUInt64:
				// tag type is uint64, need convert data value to uint64
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertUint64Value, convertUint64ValueErr := sourceDataRawPtr.ToUint64E()
				if convertUint64ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertUint64Value)
				break
			case ctdStructTagParamNameTypeUIntPtr:
				// tag type is uintptr, need convert data value to uintptr
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertUintptrValue, convertUintptrValueErr := sourceDataRawPtr.ToUintptrE()
				if convertUintptrValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertUintptrValue)
				break
			case ctdStructTagParamNameTypeInt:
				// tag type is int, need convert data value to int
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertIntValue, convertIntValueErr := sourceDataRawPtr.ToIntE()
				if convertIntValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertIntValue)
				break
			case ctdStructTagParamNameTypeInt8:
				// tag type is int8, need convert data value to int8
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertInt8Value, convertInt8ValueErr := sourceDataRawPtr.ToInt8E()
				if convertInt8ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertInt8Value)
				break
			case ctdStructTagParamNameTypeInt16:
				// tag type is int16, need convert data value to int16
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertInt16Value, convertInt16ValueErr := sourceDataRawPtr.ToInt16E()
				if convertInt16ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertInt16Value)
				break
			case ctdStructTagParamNameTypeInt32:
				// tag type is int32, need convert data value to int32
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertInt32Value, convertInt32ValueErr := sourceDataRawPtr.ToInt32E()
				if convertInt32ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertInt32Value)
				break
			case ctdStructTagParamNameTypeInt64:
				// tag type is int64, need convert data value to int64
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertInt64Value, convertInt64ValueErr := sourceDataRawPtr.ToInt64E()
				if convertInt64ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertInt64Value)
				break
			case ctdStructTagParamNameTypeFloat32:
				// tag type is float32, need convert data value to float32
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertFloat32Value, convertFloat32ValueErr := sourceDataRawPtr.ToFloat32E()
				if convertFloat32ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertFloat32Value)
				break
			case ctdStructTagParamNameTypeFloat, ctdStructTagParamNameTypeFloat64:
				// tag type is float, need convert data value to float64
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertFloat64Value, convertFloat64ValueErr := sourceDataRawPtr.ToFloat64E()
				if convertFloat64ValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertFloat64Value)
				break
			case ctdStructTagParamNameTypeBoolean:
				// tag type is boolean, need convert data value to boolean
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsBoolean() {
					continue
				}
				convertBooleanValue, convertBooleanValueErr := sourceDataRawPtr.ToBooleanE()
				if convertBooleanValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertBooleanValue)
				break
			case ctdStructTagParamNameTypeNull:
				// tag type is null, need convert data value to nil
				if sourceDataRawPtr.IsNil() {
					continue
				}
				sourceDataRawPtr.Set(nil)
				break
			case ctdStructTagParamNameTypeTime:
				// tag type is time, need convert data value to time
				// check data type is time.Time
				if sourceDataRawPtr.IsTime() {
					continue
				}
				// check data value is nil
				if sourceDataRawPtr.IsNil() {
					continue
				}
				convertStrValue, convertStrValueErr := sourceDataRawPtr.ToStringE()
				if convertStrValueErr != nil {
					continue
				}
				timeLayoutFormatStr, timeLayoutFormatStrHasExists := configTagKV[ctdStructTagParamNameTimeFormat]
				if !timeLayoutFormatStrHasExists {
					continue
				}
				if len(timeLayoutFormatStr) == 0 {
					timeLayoutFormatStr = ctdStructTagParamValueDefaultFormat
				}
				parsedTimeFormat, parsedTimeFormatErr := time.Parse(timeLayoutFormatStr, convertStrValue)
				if parsedTimeFormatErr != nil {
					// parse time with default time format
					parsedTimeFormat, parsedTimeFormatErr = time.Parse(time.RFC3339Nano, convertStrValue)
				}
				if parsedTimeFormatErr != nil {
					continue
				}
				sourceDataRawPtr.Set(parsedTimeFormat)
				break
			case ctdStructTagParamNameTypeArray:
				// data type is array
				if sourceDataRawPtr.IsNil() {
					continue
				}
				if sourceDataRawPtr.IsArray() && !sourceDataRawPtr.IsBytes() {
					continue
				}
				convertArrayValue, convertArrayValueErr := sourceDataRawPtr.ToArrayE()
				if convertArrayValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertArrayValue)
				break
			case ctdStructTagParamNameTypeObject, ctdStructTagParamNameTypeMap:
				// data type is object
				if sourceDataRawPtr.IsNil() || sourceDataRawPtr.IsStruct() || sourceDataRawPtr.IsMap() {
					continue
				}
				convertMapValue, convertMapValueErr := sourceDataRawPtr.ToMapE()
				if convertMapValueErr != nil {
					continue
				}
				sourceDataRawPtr.Set(convertMapValue)
				break
			}
			break
		}
	}
}

// =====================================================
// parseConfigCdtTagToKV parse config tags to kv
func parseConfigCdtTagToKV(fieldRef reflect.StructField) ctdStructTagKV {
	readConfigTagStr := strings.TrimSpace(fieldRef.Tag.Get(ctdStructTag))
	if len(readConfigTagStr) == 0 {
		return nil
	}
	configTagsArray := strings.Split(readConfigTagStr, ctdStructTagParamInterval)
	if len(configTagsArray) == 0 {
		return nil
	}
	configTagKV := make(ctdStructTagKV, 0)
	for _, configTagStr := range configTagsArray {
		configTagStrToLowerStr := strings.ToLower(configTagStr)
		if len(configTagStrToLowerStr) == 0 {
			continue
		}
		// read : after str
		midStrIndex := strings.Index(configTagStrToLowerStr, ctdStructTagParamSeparator)
		if midStrIndex == -1 {
			continue
		}
		tagNameLowerStr := strings.TrimSpace(configTagStrToLowerStr[:midStrIndex])
		if len(tagNameLowerStr) == 0 {
			continue
		}
		if midStrIndex+1 >= len(configTagStrToLowerStr) {
			continue
		}
		var tagValueLowerStr string
		switch tagNameLowerStr {
		case ctdStructTagParamNameTimeFormat:
			tagValueLowerStr = strings.TrimSpace(configTagStrToLowerStr[midStrIndex+1:])
			// remove ' from string
			if strings.Index(tagValueLowerStr, "'") > -1 {
				tagValueLowerStr = strings.ReplaceAll(tagValueLowerStr, "'", "")
			}
			// remove " from string
			if strings.Index(tagValueLowerStr, "\"") > -1 {
				tagValueLowerStr = strings.ReplaceAll(tagValueLowerStr, "\"", "")
			}
			if len(tagValueLowerStr) == 0 {
				// default time format
				tagValueLowerStr = ctdStructTagParamValueDefaultFormat
			}
			break
		default:
			tagValueLowerStr = strings.TrimSpace(configTagStrToLowerStr[midStrIndex+1:])
			break
		}
		configTagKV[tagNameLowerStr] = tagValueLowerStr
	}
	return configTagKV
}
