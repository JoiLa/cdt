package test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/JoiLa/cdt"
)

func TestParseFloat64(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		var i any
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 0 {
			t.Fatalf("return must be 0 instead of %v", convertedNum)
		}
	})
	t.Run("Ptr Number input", func(t *testing.T) {
		var ptr = 1
		var i any = &ptr
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 1 {
			t.Fatalf("return must be 0 instead of %v", convertedNum)
		}
	})
	t.Run("Ptr Nil input", func(t *testing.T) {
		var ptr any = nil
		var i any = &ptr
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 0 {
			t.Fatalf("return must be 0 instead of %v", convertedNum)
		}
	})
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20.4")
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 20.4 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("String Number", func(t *testing.T) {
		var i = "20.4"
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 20.4 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("Byte Number", func(t *testing.T) {
		var i = []byte("20.4")
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 20.4 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("Json RawMessage Number", func(t *testing.T) {
		var i = json.RawMessage("20.4")
		convertedNum, _ := cdt.NewConvert(i).ToFloat64E()
		if convertedNum != 20.4 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20.4
		_, err := cdt.NewConvert(i).ToFloat64E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("String NOK", func(t *testing.T) {
		type str string
		var i str = "20.4"
		_, err := cdt.NewConvert(i).ToFloat64E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseFloat32(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToFloat32E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToFloat32E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("should be int", func(t *testing.T) {
		var expectedDataType = reflect.Float32
		var i any = 20

		val, err := cdt.NewConvert(i).ToFloat32E()
		actualDataType := reflect.TypeOf(val).Kind()

		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}

		if actualDataType != expectedDataType {
			t.Fatalf("Expects to be type of %v, but got type of %v.", expectedDataType, actualDataType)
		}
	})
}
