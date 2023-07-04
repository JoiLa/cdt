package test

import (
	"encoding/json"
	"testing"

	"github.com/JoiLa/cdt"
)

func TestParseUint(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToUintE()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToUintE()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseUint8(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToUint8E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToUint8E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}
func TestParseUint16(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToUint16E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToUint16E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseUint32(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToUint32E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToUint32E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseUint64(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToUint64E()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToUint64E()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}

func TestParseUintptr(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i any = json.Number("20")
		convertedNum, _ := cdt.NewConvert(i).ToUintptrE()
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i any = 20
		_, err := cdt.NewConvert(i).ToUintptrE()
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}
