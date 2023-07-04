package test

import (
	"testing"

	"github.com/JoiLa/cdt"
)

func TestParseBoolean(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		var i any
		convertedBool, _ := cdt.NewConvert(i).ToBooleanE()
		if convertedBool != false {
			t.Fatalf("return must be false instead of %v", convertedBool)
		}
	})
	t.Run("Zero Number", func(t *testing.T) {
		var i any = 0
		convertedBool, _ := cdt.NewConvert(i).ToBooleanE()
		if convertedBool != false {
			t.Fatalf("return must be false instead of %v", convertedBool)
		}
	})
	t.Run("String 0", func(t *testing.T) {
		var i any = "0"
		convertedBool, _ := cdt.NewConvert(i).ToBooleanE()
		if convertedBool != false {
			t.Fatalf("return must be false instead of %v", convertedBool)
		}
	})
	t.Run("One Number", func(t *testing.T) {
		var i any = 1
		convertedBool, _ := cdt.NewConvert(i).ToBooleanE()
		if convertedBool != true {
			t.Fatalf("return must be true instead of %v", convertedBool)
		}
	})
	t.Run("String 1", func(t *testing.T) {
		var i any = "1"
		convertedBool, _ := cdt.NewConvert(i).ToBooleanE()
		if convertedBool != true {
			t.Fatalf("return must be true instead of %v", convertedBool)
		}
	})
	t.Run("NString", func(t *testing.T) {
		var i any = "test string"
		_, err := cdt.NewConvert(i).ToBooleanE()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("NNumber", func(t *testing.T) {
		var i any = 9999
		_, err := cdt.NewConvert(i).ToBooleanE()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}
