package cdt

import (
	"context"
	"encoding/json"
	"reflect"

	"gorm.io/gorm/schema"
)

type gormSerializerValuerInterface interface {
	Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error)
}
type gormSerializerInterface interface {
	Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error
	gormSerializerValuerInterface
}

// Scan ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (c *DataRaw) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) error {
	c.OriginVal = dbValue
	configTagKV := parseConfigCdtTagToKV(field.StructField)
	if len(configTagKV) > 0 {
		currentFieldValueRef := reflect.ValueOf(c)
		decodeDataForCdtTagKV(currentFieldValueRef.Elem(), configTagKV)
	}
	return nil
}

// Value ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (c *DataRaw) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	if fieldValue == nil {
		return nil, nil
	}
	refVal := reflect.ValueOf(fieldValue)
	switch refVal.Kind() {
	case reflect.Ptr:
		return c.Value(ctx, field, dst, refVal.Elem().Interface())
	default:
		val := refVal.Interface()
		switch val.(type) {
		case DataRaw:
			sourceDataRaw := val.(DataRaw)
			if sourceDataRaw.IsPtr() {
				return c.Value(ctx, field, dst, *NewConvert(sourceDataRaw.GetOriginValRef().Elem().Interface()))
			}
			if sourceDataRaw.IsNil() || sourceDataRaw.IsString() ||
				sourceDataRaw.IsBoolean() || sourceDataRaw.IsBytes() ||
				sourceDataRaw.IsTime() {
				return sourceDataRaw.OriginVal, nil
			}
			if sourceDataRaw.IsNumeric() {
				return sourceDataRaw.ToFloat64E()
			}
			str, _ := sourceDataRaw.ToStringE()
			return str, nil
		default:
			return json.Marshal(fieldValue)
		}
	}
}
