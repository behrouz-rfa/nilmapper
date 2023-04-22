# nilmapper

nilmapper is a Go library that provides a way to
map values from one struct to another, while also handling nil values.

# Installation

To use nilmapper in your Go project, you can simply run:

```
go get github.com/behrouz-rfa/nilmapper
```

# Usage
Here's a simple example of how you can use nilmapper to map values from one struct to another

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

```
In this example, we have two structs, Source and Destination, with different fields.
We want to map values from Source to Destination, but since the types of the fields
are different, we need to use nilmapper to handle the mapping.
To do this, we call the Map function and pass in the source and destination structs 
as arguments. Map will map the values from Source to Destination, taking 
care of nil values in the process.

# Contributing
If you find a bug or have a feature request, please open an issue on the GitHub repository.
Pull requests are also welcome! If you would like to contribute to nilmapper, 
please fork the repository and create a new branch for your changes. Once you have 
made your changes, submit a pull request and I will review your changes as soon as possible.

# License
nilmapper is licensed under the MIT License. See LICENSE for more information.

# Supported

- [x] support imperative type
- [x] support if src name is not same as the dest (src.FiledID  > src.FiledId)
- [x] support nil slice nil
- [x] support nil object
- [x] support nil imperative type