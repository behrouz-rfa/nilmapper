package nilmapper

import "reflect"

func ToValue[T any](s T) *T {
	return &s
}
func getComplex(srcFieldValue reflect.Value) complex128 {
	var f complex128
	if srcFieldValue.Kind() == reflect.Ptr {
		f = srcFieldValue.Elem().Complex()
	} else {
		f = srcFieldValue.Complex()
	}
	return f
}

func getFloat(srcFieldValue reflect.Value) float64 {
	var f float64
	if srcFieldValue.Kind() == reflect.Ptr {
		f = srcFieldValue.Elem().Float()
	} else {
		f = srcFieldValue.Float()
	}
	return f
}
func getUint(srcFieldValue reflect.Value) uint64 {
	var f uint64
	if srcFieldValue.Kind() == reflect.Ptr {
		f = srcFieldValue.Elem().Uint()
	} else {
		f = srcFieldValue.Uint()
	}
	return f
}

func getIntFromSrcValue(srcFieldValue reflect.Value) int64 {
	var f int64
	if srcFieldValue.Kind() == reflect.Ptr {
		f = srcFieldValue.Elem().Int()
	} else {
		f = srcFieldValue.Int()
	}
	return f
}
