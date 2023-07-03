package test

import (
	"fmt"
	"testing"
	"time"

	"cdt"
)

type jsonObject struct {
	String          cdt.DataRaw `json:"string" cdt:"type=string"`
	StringToInt     cdt.DataRaw `json:"string_to_int" cdt:"type=int"`
	StringToBoolean cdt.DataRaw `json:"string_to_boolean" cdt:"type=boolean"`
	StringToFloat   cdt.DataRaw `json:"string_to_float" cdt:"type=float"`
	StringToNull    cdt.DataRaw `json:"string_to_null" cdt:"type=null"`
	StringToArray   cdt.DataRaw `json:"string_to_array" cdt:"type=array"`
	StringToObject  cdt.DataRaw `json:"string_to_object" cdt:"type=object"`
	Time            cdt.DataRaw `json:"time" cdt:"type=time,time_format=2006-01-02 15:04:05"`
}

type jsonData struct {
	String     cdt.DataRaw `json:"string" cdt:"type=string"`
	Int        cdt.DataRaw `json:"int" cdt:"type=int"`
	Float      cdt.DataRaw `json:"float" cdt:"type=float"`
	Boolean    cdt.DataRaw `json:"boolean" cdt:"type=boolean"`
	Null       cdt.DataRaw `json:"null" cdt:"type=null"`
	Array      cdt.DataRaw `json:"array" cdt:"type=array"`
	Object     cdt.DataRaw `json:"object" cdt:"type=object"`
	Time       cdt.DataRaw `json:"time" cdt:"type=time,time_format='2006-01-02 15:04:05'"`
	JsonObject jsonObject  `json:"json_object"`
}

// TestJson encode and decode json data test
func TestJson(t *testing.T) {
	//
	encodeJsonData := new(jsonData)
	var name = "zhang-san"
	encodeJsonData.String.Set(&name)
	encodeJsonData.Int.Set(18)
	encodeJsonData.Float.Set(1.8)
	encodeJsonData.Boolean.Set(true)
	encodeJsonData.Null.Set(nil)
	encodeJsonData.Array.Set([]any{
		"zhang-san",
		"li-si",
		"wang-wu",
	})
	encodeJsonData.Object.Set(map[string]any{
		"index": 1,
		"name":  "zhang-san",
	})
	// timeValue, _ := time.Parse(ctdStructTagParamValueDefaultFormat, "2020-01-01 00:00:00")
	timeValue := time.Now()
	encodeJsonData.Time.Set(timeValue)
	encodeJsonData.JsonObject.String.Set("lisi")
	encodeJsonData.JsonObject.StringToInt.Set("18.59")
	//
	encodeJsonData.JsonObject.StringToInt.SetCtdStructTagKV(map[string]string{
		"type": "float",
	})
	encodeJsonData.JsonObject.Time.Set(&timeValue)
	encodeJsonData.JsonObject.StringToArray.Set([]any{
		"zhang-san",
		"li-si",
	})
	encodeJsonData.JsonObject.StringToObject.Set(map[string]any{
		"index": 1,
		"name":  "zhang-san55",
	})
	//
	jsonMarshalBytes, jsonMarshalBytesErr := cdt.MarshalJSON(encodeJsonData)
	if jsonMarshalBytesErr != nil {
		t.Fatalf("encodeJsonData error: %v", jsonMarshalBytesErr)
	}
	fmt.Println("encodeJsonData bytes: ", string(jsonMarshalBytes))
	parseJsonData := new(jsonData)
	parseJsonDataErr := cdt.UnmarshalJSONFromByte(jsonMarshalBytes, &parseJsonData)
	if parseJsonDataErr != nil {
		t.Fatalf("parse encodeJsonData json unmarshal error: %v", parseJsonDataErr)
	}
	t.Run("test string is string", func(t *testing.T) {
		if !parseJsonData.String.IsString() {
			t.Fatalf("parse encodeJsonData name is not string")
		}
	})
	t.Run("test int is float", func(t *testing.T) {
		if !parseJsonData.Int.IsInt() {
			t.Fatalf("parse encodeJsonData int is not int")
		}
	})
	t.Run("test float is float", func(t *testing.T) {
		if !parseJsonData.Float.IsFloat() {
			t.Fatalf("parse encodeJsonData float not float")
		}
	})
	t.Run("test int is numeric", func(t *testing.T) {
		if !parseJsonData.Int.IsNumeric() {
			t.Fatalf("parse encodeJsonData int is not int")
		}
	})
	t.Run("test float is numeric", func(t *testing.T) {
		if !parseJsonData.Float.IsNumeric() {
			t.Fatalf("parse encodeJsonData float not numeric")
		}
	})

	t.Run("test boolean is boolean", func(t *testing.T) {
		if !parseJsonData.Boolean.IsBoolean() {
			t.Fatalf("parse encodeJsonData boolean not boolean")
		}
	})
	t.Run("test null is null", func(t *testing.T) {
		if !parseJsonData.Null.IsNil() {
			t.Fatalf("parse encodeJsonData null not null")
		}
	})
	t.Run("test array is array", func(t *testing.T) {
		if !parseJsonData.Array.IsArray() {
			t.Fatalf("parse encodeJsonData array not array")
		}
	})
	t.Run("test object is map", func(t *testing.T) {
		if !parseJsonData.Object.IsMap() {
			t.Fatalf("parse encodeJsonData object not object")
		}
	})
	t.Run("test time is time", func(t *testing.T) {
		if !parseJsonData.Time.IsTime() {
			t.Fatalf("parse encodeJsonData time not time")
		}
	})

}
