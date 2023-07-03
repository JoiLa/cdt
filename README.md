<p align="center" style="text-align: center">
  <img src="./assets/images/logo.png" width="55%"><br/>
</p>

<p align="center">
  convert between different data types easily and conveniently
  <br/>
  <br/>

  <a href="https://github.com/Joila/cdt/tags" rel="nofollow">
    <img alt="GitHub tag (latest SemVer pre-release)" src="https://img.shields.io/github/v/tag/Joila/cdt?include_prereleases&label=version"/>
  </a>
</p>

<div align="center">
<strong>
<samp>

[English](./README.md) · [简体中文](./README.zh-Hans.md)

</samp>
</strong>
</div>

# CDT (Common Data Types) Library

CDT is a Go library that aims to provide simple and flexible data type handling utilities.
It offers a set of functions and structs for working with different types of data and provides convenient methods for data type conversions and operations.

# Features

- **Common Data Types**: CDT supports handling various common data types, including strings, integers, floats, booleans, arrays, objects(map), and timestamps.
- **Flexible Tag Settings**: By using the `cdt` tag on struct fields, it's easy to specify the data type and format information for the field.
- **Simple Value Setting**: The provided `Set` method makes it easy to set values of different types to their respective fields.
- **JSON Serialization and Deserialization**: CDT provides functionality to serialize data structures to JSON format and deserialize them from JSON data conveniently.
- **SQL Serialization and Deserialization**: The CDT library provides the function of serializing data structures into SQL format, and can easily deserialize data structures from SQL data.
- **Type Checking and Conversion**: CDT offers a range of methods for checking the data type of fields and provides convenient type conversion methods for easy conversion between different types.



# Getting Started

## Installing

To start using CDT, install Go and run `go get`:

```sh
$ go get github.com/JoiLa/cdt
```

This will retrieve the library.

##  Example

```go
package main

import (
	`fmt`
	`time`
	`github.com/JoiLa/cdt`
)

// testDataStruct is a test data struct
type testDataStruct struct {
	String  cdt.DataRaw `json:"string" cdt:"type=string"`
	Int     cdt.DataRaw `json:"int" cdt:"type=int"`     // type=int || type=int8 || type=int16 || type=int32 || type=int64 || type=uint || type=uint8 || type=uint16 || type=uint32 || type=uint64 || type=uintptr
	Float   cdt.DataRaw `json:"float" cdt:"type=float"` // type=float || type=float32 || type=float64
	Boolean cdt.DataRaw `json:"boolean" cdt:"type=boolean"`
	Null    cdt.DataRaw `json:"null" cdt:"type=null"`
	Array   cdt.DataRaw `json:"array" cdt:"type=array"`
	Object  cdt.DataRaw `json:"object" cdt:"type=object"` // type=object || type=map
	Time    cdt.DataRaw `json:"time" cdt:"type=time,time_format=2006-01-02 15:04:05"`
	Time2   cdt.DataRaw `json:"time2" cdt:"type=time,time_format='2006-01-02 15:04:05'"`
	Time3   cdt.DataRaw `json:"time3" cdt:"type=time,time_format=\"2006-01-02 15:04:05\""`
}

func main() {
	testData := new(testDataStruct)
	var name = "zhang-san"
	testData.String.Set(&name)
	testData.Int.Set(float64(18))
	testData.Float.Set(1.8)
	testData.Boolean.Set(true)
	testData.Null.Set(nil)
	testData.Array.Set([]any{
		"zhang-san",
		"li-si",
		"wang-wu",
	})
	testData.Object.Set(map[string]any{
		"index": 1,
		"name":  "zhang-san",
	})
	fixedTimeStr := "2020-01-15 15:19:30"
	fixedTime, fixedTimeErr := time.Parse("2006-01-02 15:04:05", fixedTimeStr)
	if fixedTimeErr != nil {
		fmt.Println("fixedTimeErr:", fixedTimeErr)
	}
	testData.Time.Set(fixedTime)
	testData.Time2.Set(fixedTimeStr)
	testData.Time3.Set(fixedTime)
	resultBytes, _ := cdt.MarshalJSONToIndent(testData, "", "\t")
	fmt.Println("// json echo")
	fmt.Println(string(resultBytes))
	// json parse
	parseTestData := new(testDataStruct)
	_ = cdt.UnmarshalJSONFromByte(resultBytes, &parseTestData)
	fmt.Println("// check data type")
	fmt.Println("parseTestData String is string:", parseTestData.String.IsString())
	fmt.Println("parseTestData Int is int:", parseTestData.Int.IsInt())
	fmt.Println("parseTestData Int is numeric:", parseTestData.Int.IsNumeric())
	fmt.Println("parseTestData Float is int:", parseTestData.Float.IsInt())
	fmt.Println("parseTestData Float is float:", parseTestData.Float.IsFloat())
	fmt.Println("parseTestData Float is numeric:", parseTestData.Float.IsNumeric())
	fmt.Println("parseTestData Boolean is boolean:", parseTestData.Boolean.IsBoolean())
	fmt.Println("parseTestData Null is null:", parseTestData.Null.IsNil())
	fmt.Println("parseTestData Array is array:", parseTestData.Array.IsArray())
	fmt.Println("parseTestData Object is object:", parseTestData.Object.IsMap())
	fmt.Println("parseTestData Time is time:", parseTestData.Time.IsTime())
	fmt.Println("parseTestData Time2 is time:", parseTestData.Time2.IsTime())
	fmt.Println("parseTestData Time3 is time:", parseTestData.Time3.IsTime())
	fmt.Println("// data value type convert")
	fmt.Print("parseTestData Int ToStringE: ")
	fmt.Println(parseTestData.Int.ToStringE())
	fmt.Print("parseTestData Int ToString: ")
	fmt.Println(parseTestData.Int.ToString())
	fmt.Print("parseTestData Int ToStringPtr: ")
	fmt.Println(parseTestData.Int.ToStringPtr())
	fmt.Print("parseTestData Int ToStringPtrE: ")
	fmt.Println(parseTestData.Int.ToStringPtrE())

	// other convert method
	// parseTestData.*.ToInt
	// parseTestData.*.ToInt8
	// parseTestData.*.ToInt16
	// parseTestData.*.ToInt32
	// parseTestData.*.ToInt64
	// parseTestData.*.ToUint
	// parseTestData.*.ToUint8
	// parseTestData.*.ToUint16
	// parseTestData.*.ToUint32
	// parseTestData.*.ToUint64
	// parseTestData.*.ToFloat32
	// parseTestData.*.ToUintptr
	// parseTestData.*.ToFloat64
	// parseTestData.*.ToString
	// parseTestData.*.ToBoolean
	// parseTestData.*.ToMap
	// parseTestData.*.ToArray
	// parseTestData.*.ToTime

	// other convert method
	// parseTestData.*.To*Ptr
	// parseTestData.*.To*PtrE
	// parseTestData.*.To*E
}
```

## This will print:

```sh
// json echo
{
        "string": "zhang-san",
        "int": 18,
        "float": 1.8,
        "boolean": true,
        "null": null,
        "array": [
                "zhang-san",
                "li-si",
                "wang-wu"
        ],
        "object": {
                "index": 1,
                "name": "zhang-san"
        },
        "time": "2020-01-15 15:19:30",
        "time2": "2020-01-15 15:19:30",
        "time3": "2020-01-15 15:19:30"
}
// check data type
parseTestData String is string: true
parseTestData Int is int: true
parseTestData Int is numeric: true
parseTestData Float is int: false
parseTestData Float is float: true
parseTestData Float is numeric: true
parseTestData Boolean is boolean: true
parseTestData Null is null: true
parseTestData Array is array: true
parseTestData Object is object: true
parseTestData Time is time: true
parseTestData Time2 is time: true
parseTestData Time3 is time: true
// data value type convert
parseTestData Int ToStringE: 18 <nil>
parseTestData Int ToString: 18
parseTestData Int ToStringPtr: 0xc000127d40
parseTestData Int ToStringPtrE: 0xc000127d50 <nil>
```



# Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.



# License

CDT library is licensed under the [MIT License](https://opensource.org/licenses/MIT).
