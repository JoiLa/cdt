package test

import (
	"fmt"
	"testing"

	"github.com/JoiLa/cdt"
)

// TestConvertString test convert string
func TestConvertString(t *testing.T) {
	convertStringEncode := cdt.NewConvert("test")
	convertStringEncode.SetCtdStructTagKV(map[string]string{
		"type": "string",
	})
	jsonByte, jsonByteErr := cdt.MarshalJSONToString(convertStringEncode)
	if jsonByteErr != nil {
		t.Fatalf("jsonByteErr: %v", jsonByteErr)
	}
	fmt.Println(fmt.Sprintf("json Str：%s", jsonByte))
	convertStringDecode := cdt.NewConvert(nil)
	convertStringDecode.SetCtdStructTagKV(map[string]string{
		"type": "string",
	})
	decodeJsonStrErr := cdt.UnmarshalJSONFromString(jsonByte, &convertStringDecode)
	if decodeJsonStrErr != nil {
		t.Fatalf("decodeJsonStrErr: %v", decodeJsonStrErr)
	}
	fmt.Println("convertStringDecode: ", convertStringDecode)
}

// TestConvertArray test convert array
func TestConvertArray(t *testing.T) {
	convertArrayEncode := cdt.NewConvert([]any{
		1,
		2,
		4,
	})
	convertArrayEncode.SetCtdStructTagKV(map[string]string{
		"type": "array",
	})
	jsonByte, jsonByteErr := cdt.MarshalJSONToString(convertArrayEncode)
	if jsonByteErr != nil {
		t.Fatalf("jsonByteErr: %v", jsonByteErr)
	}
	fmt.Println(fmt.Sprintf("json Str：%s", jsonByte))
	convertArrayDecode := cdt.NewConvert(nil)
	convertArrayDecode.SetCtdStructTagKV(map[string]string{
		"type": "array",
	})
	decodeJsonStrErr := cdt.UnmarshalJSONFromString(jsonByte, &convertArrayDecode)
	if decodeJsonStrErr != nil {
		t.Fatalf("decodeJsonStrErr: %v", decodeJsonStrErr)
	}
	fmt.Println("convertArrayDecode: ", convertArrayDecode)
}

// TestConvertMap test convert map
func TestConvertMap(t *testing.T) {
	convertMapEncode := cdt.NewConvert(map[string]interface{}{
		"1": 1,
		"2": 2,
	})
	convertMapEncode.SetCtdStructTagKV(map[string]string{
		"type": "map",
	})
	jsonByte, jsonByteErr := cdt.MarshalJSONToString(convertMapEncode)
	if jsonByteErr != nil {
		t.Fatalf("jsonByteErr: %v", jsonByteErr)
	}
	fmt.Println(fmt.Sprintf("json Str：%s", jsonByte))
	convertMapDecode := cdt.NewConvert(nil)
	convertMapDecode.SetCtdStructTagKV(map[string]string{
		"type": "map",
	})
	decodeJsonStrErr := cdt.UnmarshalJSONFromString(jsonByte, &convertMapDecode)
	if decodeJsonStrErr != nil {
		t.Fatalf("decodeJsonStrErr: %v", decodeJsonStrErr)
	}
	fmt.Println("convertMapDecode: ", convertMapDecode)
}
