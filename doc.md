<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# nilmapper

```go
import "github.com/behrouz-rfa/nilmapper"
```

## Index

- [func Copy(source interface{}, destination interface{})](<#func-copy>)
- [func CopySlice(source interface{}, destination interface{})](<#func-copyslice>)
- [func ToValue[T any](s T) *T](<#func-tovalue>)


## func Copy

```go
func Copy(source interface{}, destination interface{})
```

Copy maps the fields of a source struct or slice to a destination struct or slice. If nested is true, it recursively maps nested structs or slices.

Example usage:

```
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

src := SourceStruct{
    FieldA: "Test1",
    FieldB: 123,
    FieldC: nil,
}

var dest DestStruct
Copy(src, &dest)

fmt.Println(dest.FieldA, dest.FieldB, dest.FieldC)
// Output: Test1 123 ""
```

## func CopySlice

```go
func CopySlice(source interface{}, destination interface{})
```

## func ToValue

```go
func ToValue[T any](s T) *T
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)