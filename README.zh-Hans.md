<p align="center" style="text-align: center">
  <img src="./assets/images/logo.png" width="55%"><br/>
</p>

<p align="center">
  轻松方便地在不同数据类型之间进行转换
  <br/>
  <br/>

  <a href="https://github.com/Joila/cdt/tags" rel="nofollow">
    <img alt="GitHub tag (latest SemVer pre-release)" src="https://img.shields.io/github/v/tag/Joila/cdt?include_prereleases&label=version"/>
  </a>
</p>

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md)

</samp>
</strong>
</div>

# CDT（通用数据类型）库

CDT（Common Data Types）是一个Go语言库，旨在提供简单而灵活的数据类型处理工具。它提供了一组函数和结构体，用于处理不同类型的数据，并提供了方便的方法进行数据类型转换和操作。



# 特点

- **通用数据类型**：CDT库支持处理多种常见的数据类型，包括字符串、整数、浮点数、布尔值、数组、对象（MAP）和时间。
- **灵活的标签设置**：通过使用结构体字段上的`cdt`标签，可以方便地指定字段的数据类型、格式等信息。
- **简单的值设置**：通过提供的`Set`方法，可以轻松地将不同类型的值设置到相应的字段中。
- **JSON序列化与反序列化**：CDT库提供了将数据结构序列化为JSON格式的功能，并且可以方便地从JSON数据中反序列化为数据结构。
- **SQL序列化与反序列化**：CDT库提供了将数据结构序列化为SQL格式的功能，并且可以方便地从SQL数据中反序列化为数据结构。
- **类型检查和转换**：CDT库提供了一系列方法，用于检查字段的数据类型，并提供了方便的类型转换方法，使得在不同类型之间进行转换变得简单。



# 快速开始

## 安装

To start using CDT, install Go and run `go get`:

```sh
$ go get github.com/JoiLa/cdt
```

This will retrieve the library.

## 使用示例

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

## 运行结果

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



# 贡献

欢迎贡献代码和提出问题。如果您发现任何错误或有改进的建议，请在GitHub存储库上提出问题或提交拉取请求。



# 许可证

CDT 库根据 [MIT]((https://opensource.org/licenses/MIT)) 许可证获得许可。
