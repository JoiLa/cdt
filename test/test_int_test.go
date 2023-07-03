package test

import (
	"encoding/json"
	"reflect"
	"testing"

	"cdt"
)

func TestParseInt(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToIntE()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToIntE()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("should be int", func(t *testing.T) {
		var expectedDataType = reflect.Int
		var i any = 20

		val, err := cdt.NewConvert(i).ToIntE()
		actualDataType := reflect.TypeOf(val).Kind()

		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}

		if actualDataType != expectedDataType {
			t.Fatalf("Expects to be type of %v, but got type of %v.", expectedDataType, actualDataType)
		}
	})

}

func TestParseInt8(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		convertedNum, _ := cdt.NewConvert(nil).ToInt8E()
		if convertedNum != 0 {
			t.Fatalf("return must be 0 instead of %v", convertedNum)
		}
	})
	t.Run("Json Number", func(t *testing.T) {
		convertedNum, _ := cdt.NewConvert(json.Number("100")).ToInt8E()
		if convertedNum != 100.0 {
			t.Fatalf("return must be 100.0 instead of %v", convertedNum)
		}
	})
	t.Run("Error value", func(t *testing.T) {
		_, err := cdt.NewConvert("test").ToInt8E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseInt16(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		convertedNum, _ := cdt.NewConvert(nil).ToInt16E()
		if convertedNum != 0 {
			t.Fatalf("return must be 0 instead of %v", convertedNum)
		}
	})
	t.Run("Json Number", func(t *testing.T) {
		convertedNum, _ := cdt.NewConvert(json.Number("100")).ToInt16E()
		if convertedNum != 100.0 {
			t.Fatalf("return must be 100.0 instead of %v", convertedNum)
		}
	})
	t.Run("Error value", func(t *testing.T) {
		_, err := cdt.NewConvert("test").ToInt16E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseInt32(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToInt32E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToInt32E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("should be int32", func(t *testing.T) {
		var expectedDataType = reflect.Int32
		var i any = 20

		val, err := cdt.NewConvert(i).ToInt32E()
		actualDataType := reflect.TypeOf(val).Kind()

		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}

		if actualDataType != expectedDataType {
			t.Fatalf("Expects to be type of %v, but got type of %v.", expectedDataType, actualDataType)
		}
	})
}

func TestParseInt64(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToInt64E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToInt64E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}
