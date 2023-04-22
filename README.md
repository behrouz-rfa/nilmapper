# nilmapper
A simple and easy go tools for auto mapper  struct to struct, that support nil
MapStruct maps the fields of a source struct or slice to a destination struct or slice.
If nested is true, it recursively maps nested structs or slices.

Example usage:
```go
package main

import (
	"fmt"
    "github.com/behrouz-rfa/nilmapper"
)

type SourceStruct struct {
 FieldA string
 FieldB int
 FieldC *string
}

type DestStruct struct {
 FieldA *string
 FieldB int
 FieldC string
}

func main() {

 src := SourceStruct{
  FieldA: "Test1",
  FieldB: 123,
  FieldC: nil,
 }

 var dest DestStruct
	nilmapper.MapStruct(src, &dest)

 fmt.Println(dest.FieldA, dest.FieldB, dest.FieldC)
 //Output:Test1123""

 srcSlice := []SourceStruct{
  {
   FieldA: "Test1",
   FieldB: 123,
   FieldC: nil,
  },
  {
   FieldA: "Test2",
   FieldB: 456,
   FieldC: nilmapper.ToValue("Value"),
  },
 }

 var destSlice []DestStruct
	nilmapper.MapSlice(srcSlice, &destSlice)

 fmt.Println(destSlice[0].FieldA, destSlice[0].FieldB, destSlice[0].FieldC)
 fmt.Println(destSlice[1].FieldA, destSlice[1].FieldB, destSlice[1].FieldC)
 //Output:
 // Test1
 // 123
 // ""
 // Test2
 // 456
 // Value

}

```