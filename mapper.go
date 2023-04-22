package nilmapper

import (
	"reflect"
	"strings"
)

// The CopySlice function maps a slice of source struct values to a slice of
// destination struct values.
// It takes two parameters - source and destination - both of which are interfaces.
// The source parameter should be a slice of structs, while the destination parameter
// should be a pointer to a slice of structs.
// The function first gets the length of the source slice and creates a new
// slice of the same length with the type of the destination slice.
// It then iterates over each element of the source slice, maps it to a new
// element of the destination slice, and sets the value of the new element
// in the destination slice. The mapStruct function is called for each
// element to perform the actual mapping.
// This function can be used to quickly and easily map entire slices
// of structs from one type to another.
//
// Example:
//
//	func ExampleMapSlice() {
//		type SourceStruct struct {
//			FieldA string
//			FieldB int
//			FieldC *string
//		}
//
//		type DestStruct struct {
//			FieldA *string
//			FieldB int
//			FieldC string
//		}
//
//		src := []SourceStruct{
//			{
//				FieldA: "Test1",
//				FieldB: 123,
//				FieldC: nil,
//			},
//			{
//				FieldA: "Test2",
//				FieldB: 456,
//				FieldC: ToValue("Value"),
//			},
//		}
//
//		var dest []DestStruct
//		CopySlice(src, &dest)
//
//		fmt.Println(dest[0].FieldA, dest[0].FieldB, dest[0].FieldC)
//		fmt.Println(dest[1].FieldA, dest[1].FieldB, dest[1].FieldC)
//		// Output:
//		// Test1 123
//		// Test2 456 Value
//	}

func CopySlice(source interface{}, destination interface{}) {
	srcValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(destination).Elem()
	mapSlice(srcValue, destValue)

}

func mapSlice(srcValue reflect.Value, destValue reflect.Value) {

	srcLen := srcValue.Len()
	destType := destValue.Type().Elem()
	destSlice := reflect.MakeSlice(destValue.Type(), srcLen, srcLen)
	for i := 0; i < srcLen; i++ {
		srcElem := srcValue.Index(i)
		destElem := reflect.New(destType).Elem()
		mapStruct(srcElem.Interface(), destElem.Addr().Interface(), false)
		destSlice.Index(i).Set(destElem)
	}
	destValue.Set(destSlice)
}

// Copy maps the fields of a source struct or slice to a destination struct or slice.
// If nested is true, it recursively maps nested structs or slices.
//
// Example usage:
//
//	type SourceStruct struct {
//	    FieldA string
//	    FieldB int
//	    FieldC *string
//	}
//
//	type DestStruct struct {
//	    FieldA *string
//	    FieldB int
//	    FieldC string
//	}
//
//	src := SourceStruct{
//	    FieldA: "Test1",
//	    FieldB: 123,
//	    FieldC: nil,
//	}
//
//	var dest DestStruct
//	Copy(src, &dest)
//
//	fmt.Println(dest.FieldA, dest.FieldB, dest.FieldC)
//	// Output: Test1 123 ""
func Copy(source interface{}, destination interface{}) {
	mapStruct(source, destination, false)
}

func mapStruct(source interface{}, destination interface{}, nested bool) {
	srcValue := reflect.ValueOf(source)
	srcValue2 := reflect.TypeOf(source)
	destValue := reflect.ValueOf(destination).Elem()
	if srcValue.Kind() == reflect.Slice || destValue.Kind() == reflect.Slice && !nested {
		mapSlice(srcValue, destValue)
		return
	}
	var size int
	if nested {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		size = srcValue.NumField()
	} else {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		size = srcValue.NumField()
	}
	for i := 0; i < size; i++ {
		var name string
		if srcValue2.Kind() == reflect.Ptr {
			srcFileTypeName := srcValue2.Elem().Field(i)
			name = srcFileTypeName.Name
		} else {
			srcFileTypeName := srcValue2.Field(i)
			name = srcFileTypeName.Name
		}

		srcFieldValue := srcValue.Field(i)

		destFieldValue := destValue.FieldByName(name)
		if !destFieldValue.IsValid() {
			destFieldValue = destValue.FieldByNameFunc(func(s string) bool {
				if strings.ToLower(s) == strings.ToLower(name) {
					return true
				}
				return false
			})
			if !destFieldValue.IsValid() {
				continue
			}
		}
		if !destFieldValue.CanSet() {
			continue
		}
		//destFieldValue := destValue.Field(i)

		srcFieldType := srcFieldValue.Type()
		destFieldType := destFieldValue.Type()

		if srcFieldType.Kind() == reflect.Ptr {
			srcFieldType = srcFieldType.Elem()
		}

		if destFieldType.Kind() == reflect.Ptr {
			destFieldType = destFieldType.Elem()
		}

		if srcFieldType == destFieldType || destFieldType.Kind() == reflect.Interface {
			if srcFieldValue.Kind() == reflect.Ptr && srcFieldValue.IsNil() {
				if srcFieldValue.IsNil() {
					continue
				}
				destFieldValue.Set(reflect.Zero(destFieldType))
			} else {
				if srcFieldType.Kind() == reflect.Struct {
					newDestValue := reflect.New(destFieldType)
					mapStruct(srcFieldValue.Interface(), newDestValue.Interface(), true)
					//assignStructField(destFieldValue, newDestValue.Elem(), destFieldType)
					if destFieldValue.Kind() == reflect.Ptr {
						if destFieldValue.IsNil() {
							destFieldValue.Set(newDestValue)
						}
						destFieldValue = destFieldValue.Elem()
						destFieldValue.Set(newDestValue.Elem())

					} else {
						if newDestValue.Kind() == reflect.Ptr {
							destFieldValue.Set(newDestValue.Elem())
						} else {
							destFieldValue.Set(newDestValue)
						}

					}
				} else if srcFieldType.Kind() == reflect.Slice {
					srcSlice := srcFieldValue

					destSlice := reflect.MakeSlice(destFieldType, srcSlice.Len(), srcSlice.Len())
					for j := 0; j < srcSlice.Len(); j++ {
						if srcSlice.Index(j).Type().Kind() == reflect.Struct {
							newDestValue := reflect.New(destFieldType.Elem())
							mapStruct(srcSlice.Index(j).Interface(), newDestValue.Interface(), false)
							assignSliceElement(destSlice, newDestValue.Elem(), j)
						} else {
							assignSliceElement(destSlice, srcSlice.Index(j), j)
						}
					}
					destFieldValue.Set(destSlice)
				} else {

					assignValue(destFieldValue, srcFieldValue)
				}
			}
		} else {
			if destFieldType.Kind() == reflect.Struct {
				if destFieldValue.Kind() == reflect.Ptr {
					newDestValue := reflect.New(destFieldType)
					mapStruct(srcFieldValue.Interface(), newDestValue.Interface(), false)
					destFieldValue.Set(newDestValue.Elem().Addr())
				} else {
					newDestValue := reflect.New(destFieldType)
					mapStruct(srcFieldValue.Interface(), newDestValue.Interface(), false)
					destFieldValue.Set(newDestValue.Elem())
				}

			}
		}

	}
}

func assignStructField(destFieldValue reflect.Value, newDestValue reflect.Value, fieldType reflect.Type) {
	if destFieldValue.Kind() == reflect.Ptr {
		newDestValue := reflect.New(fieldType)
		if destFieldValue.IsNil() {
			destFieldValue.Set(newDestValue)
		}
		destFieldValue = destFieldValue.Elem()
		destFieldValue.Set(newDestValue.Elem())
	} else {
		destFieldValue.Set(newDestValue)
	}
}

func assignSliceElement(destSlice reflect.Value, value reflect.Value, index int) {
	if value.Type().Kind() == reflect.Ptr {
		destSlice.Index(index).Set(value.Elem())
	} else {
		destSlice.Index(index).Set(value)
	}
}
func assignValue(destFieldValue reflect.Value, srcFieldValue reflect.Value) {
	if destFieldValue.Kind() == reflect.Ptr {
		if destFieldValue.Kind() == reflect.Ptr {
			switch destFieldValue.Type().Elem().Kind() {
			case reflect.String:
				var str string
				if srcFieldValue.Kind() == reflect.Ptr {
					str = srcFieldValue.Elem().String()
				} else {
					str = srcFieldValue.String()
				}

				ptr := new(string)
				*ptr = str
				destFieldValue.Set(reflect.ValueOf(ptr))

			case reflect.Float32:
				f := getFloat(srcFieldValue)
				ptr := new(float32)
				*ptr = float32(f)
				destFieldValue.Set(reflect.ValueOf(ptr))

			case reflect.Float64:
				f := getFloat(srcFieldValue)
				ptr := new(float64)
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))

			case reflect.Uint:
				f := getUint(srcFieldValue)
				ptr := new(uint)
				*ptr = uint(f)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Uint8:
				f := getUint(srcFieldValue)
				ptr := new(uint8)
				*ptr = uint8(f)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Uint16:
				f := getUint(srcFieldValue)
				ptr := new(uint16)
				*ptr = uint16(f)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Uint32:
				f := getUint(srcFieldValue)
				ptr := new(uint32)
				*ptr = uint32(f)
				destFieldValue.Set(reflect.ValueOf(ptr))

			case reflect.Uint64:
				f := getUint(srcFieldValue)
				ptr := new(uint64)
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Uintptr:
				f := srcFieldValue.Interface().(uintptr)
				ptr := new(uintptr)
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))

			case reflect.Int:
				i := getIntFromSrcValue(srcFieldValue)
				ptr := new(int)
				*ptr = int(i)
				destFieldValue.Set(reflect.ValueOf(ptr))

			case reflect.Int16:
				i := getIntFromSrcValue(srcFieldValue)
				ptr := new(int16)
				*ptr = int16(i)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Int32:
				f := getIntFromSrcValue(srcFieldValue)
				ptr := new(int32)
				*ptr = int32(f)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Int64:
				f := getIntFromSrcValue(srcFieldValue)
				ptr := new(int64)
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Complex64:
				f := getComplex(srcFieldValue)
				ptr := new(complex64)
				*ptr = complex64(f)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Complex128:
				f := getComplex(srcFieldValue)
				ptr := new(complex128)
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Map:
				f := srcFieldValue.Interface().(map[string]interface{})
				ptr := new(map[string]interface{})
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Int8:
				f := getIntFromSrcValue(srcFieldValue)
				ptr := new(int8)
				*ptr = int8(f)
				destFieldValue.Set(reflect.ValueOf(ptr))
			case reflect.Bool:
				var f bool
				if srcFieldValue.Kind() == reflect.Ptr {
					f = srcFieldValue.Elem().Bool()
				} else {
					f = srcFieldValue.Bool()
				}
				ptr := new(bool)
				*ptr = f
				destFieldValue.Set(reflect.ValueOf(ptr))
			// add more cases for other types
			default:
				assignValue(destFieldValue, srcFieldValue)
			}

		} else {
			destFieldValue.Set(reflect.ValueOf(srcFieldValue.Interface()))
		}
	} else {
		if srcFieldValue.Kind() == reflect.Ptr {
			destFieldValue.Set(srcFieldValue.Elem())
		} else {
			destFieldValue.Set(srcFieldValue)
		}
	}
}
