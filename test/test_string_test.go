package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/JoiLa/cdt"
)

func TestParseString(t *testing.T) {
	tt := []struct {
		Description string
		Input       any
		Result      string
		Err         error
	}{
		{"empty value", "", "", nil},
		{"valid value", "valid value", "valid value", nil},
		{"nil value", nil, "", nil},
		{"different type", 1, "", errors.New("unable to casting number 1 (type int)")},
	}

	for _, tc := range tt {
		t.Run(tc.Description, func(t *testing.T) {
			result, err := cdt.NewConvert(tc.Input).ToStringE()
			if result != tc.Result {
				t.Fatalf("return must be %v instead of %v", tc.Result, result)
			}

			if fmt.Sprint(tc.Err) != fmt.Sprint(err) {
				t.Fatalf("error must be %v instead of %v", tc.Err, err)
			}
		})
	}
}
