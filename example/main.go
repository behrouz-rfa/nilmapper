package main

import (
	"fmt"
	"github.com/behrouz-rfa/nilmapper"
)

type SourceStruct struct {
	FieldA string
	FieldB int
	FieldC *string
	Object Object
}
type Object struct {
	Name string
}
type DestStruct struct {
	FieldA *string
	FieldB int
	FieldC string
	Object *Object
}

func main() {

	src := SourceStruct{
		FieldA: "Test1",
		FieldB: 123,
		FieldC: nil,
		Object: Object{
			Name: "NilMapper",
		},
	}

	var dest DestStruct
	nilmapper.Copy(src, &dest)

	fmt.Println(dest.FieldA, dest.FieldB, dest.FieldC)
	//Output: Test1 123""

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
	nilmapper.CopySlice(srcSlice, &destSlice)

	fmt.Println(*destSlice[0].FieldA, destSlice[0].FieldB, destSlice[0].FieldC)
	fmt.Println(*destSlice[1].FieldA, destSlice[1].FieldB, destSlice[1].FieldC)
	//Output:
	// Test1 123 ""
	// Test2 456 Value

}
