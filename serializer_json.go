package cdt

import (
	"encoding/json"
	"fmt"
)

// String return string value
func (i convert) String() string {
	return fmt.Sprintf("%v", i.OriginVal)
}

// MarshalJSON to output non base64 encoded []byte
func (c DataRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.OriginVal)
}

// UnmarshalJSON to deserialize []byte
func (c *DataRaw) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &c.OriginVal)
	if err != nil {
		return err
	}
	return nil
}

// String return string value
func (c DataRaw) String() string {
	return fmt.Sprintf("%v", c.OriginVal)
}

// MarshalJSON marshal json
func MarshalJSON(data any) ([]byte, error) {
	newData := CopyVariable(data)
	EncodeDataForTags(newData)
	return json.Marshal(newData)
}

// MarshalJSONToString marshal json to string
func MarshalJSONToString(data any) (string, error) {
	result, err := MarshalJSON(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// MarshalJSONToIndent marshal json to string
func MarshalJSONToIndent(data any, prefix, indent string) ([]byte, error) {
	newData := CopyVariable(data)
	EncodeDataForTags(newData)
	return json.MarshalIndent(newData, prefix, indent)
}

// UnmarshalJSONFromString unmarshal json from string
func UnmarshalJSONFromString(data string, v any) error {
	return UnmarshalJSONFromByte([]byte(data), v)
}

// UnmarshalJSONFromByte unmarshal json from byte
func UnmarshalJSONFromByte(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	DecodeDataForTags(v)
	return nil
}
