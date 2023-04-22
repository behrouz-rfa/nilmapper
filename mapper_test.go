package nilmapper

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

type SourceStruct struct {
	FieldA string
	FieldB float32
	FieldC *string
	FieldD *[]string
}

type DestStruct struct {
	FieldA *string
	FieldB float32
	FieldC string
	FieldD *[]string
}
type DestStruct22 struct {
	FieldA string
	FieldB *float32
	FieldC string
}

func TestMapStruct(t *testing.T) {
	t.Run("Non-Nested Struct", func(t *testing.T) {
		src := SourceStruct{
			FieldA: "Test",
			FieldB: 123,
			FieldC: nil,
		}
		dest := DestStruct{}
		mapStruct(src, &dest, false)
		if *dest.FieldA != "Test" || dest.FieldB != 123 || dest.FieldC != "" {
			t.Errorf("Expected dest to be %+v, but got %+v", DestStruct{FieldA: &src.FieldA, FieldB: src.FieldB, FieldC: ""}, dest)
		}
	})

	t.Run("Nested Struct", func(t *testing.T) {
		src := struct {
			FieldA string
			FieldB SourceStruct
		}{
			FieldA: "NestedTest",
			FieldB: SourceStruct{
				FieldA: "Test",
				FieldB: 123,
				FieldC: nil,
			},
		}
		dest := struct {
			FieldA *string
			FieldB DestStruct
		}{}
		mapStruct(src, &dest, true)
		if *dest.FieldA != "NestedTest" || *dest.FieldB.FieldA != "Test" || dest.FieldB.FieldB != 123 || dest.FieldB.FieldC != "" {
			t.Errorf("Expected dest to be %+v, but got %+v", struct {
				FieldA *string
				FieldB DestStruct
			}{
				FieldA: &src.FieldA,
				FieldB: DestStruct{FieldA: &src.FieldB.FieldA, FieldB: src.FieldB.FieldB, FieldC: ""},
			}, dest)
		}
	})

	t.Run("Slice of Structs", func(t *testing.T) {
		src := []SourceStruct{
			{
				FieldA: "Test",
				FieldB: 123,
				FieldC: nil,
			},
			{
				FieldA: "Test2",
				FieldB: 456,
				FieldC: nil,
			},
		}
		var dest []DestStruct22
		CopySlice(src, &dest)
		if len(dest) != 2 || dest[0].FieldA != "Test" || *dest[0].FieldB != 123 || dest[0].FieldC != "" || dest[1].FieldA != "Test2" || *dest[1].FieldB != 456 || dest[1].FieldC != "" {
			t.Errorf("Expected dest to be %+v, but got %+v", []DestStruct{{FieldA: &src[0].FieldA, FieldB: src[0].FieldB, FieldC: ""}, {FieldA: &src[1].FieldA, FieldB: src[1].FieldB, FieldC: ""}}, dest)
		}
	})

}

type SourceNestedStruct struct {
	FieldD string
}

type SourceStructWithNested struct {
	FieldA string
	FieldB int
	FieldC *SourceNestedStruct
}

type DestNestedStruct struct {
	FieldD string
}

type DestStructWithNested struct {
	FieldA string
	FieldB int
	FieldC DestNestedStruct
}

func TestMapStructWithNested(t *testing.T) {
	src := SourceStructWithNested{
		FieldA: "Test",
		FieldB: 123,
		FieldC: &SourceNestedStruct{FieldD: "NestedTest"},
	}
	dest := DestStructWithNested{}
	mapStruct(src, &dest, false)
	if dest.FieldC.FieldD != "NestedTest" {
		t.Errorf("Expected dest to have FieldC.FieldD=%q, but got %q", "NestedTest", dest.FieldC.FieldD)
	}
}

type SourceSliceStruct struct {
	FieldA string
	FieldB float32
}

type DestSliceStruct struct {
	FieldA *string
	FieldB float32
	FieldC string
}

func TestMapSlice(t *testing.T) {
	src := []SourceSliceStruct{
		{
			FieldA: "Test",
			FieldB: 123,
		},
		{
			FieldA: "Test2",
			FieldB: 456,
		},
	}
	var dest []DestSliceStruct
	CopySlice(src, &dest)
	if len(dest) != 2 || *dest[0].FieldA != "Test" || dest[0].FieldB != 123 || dest[0].FieldC != "" || *dest[1].FieldA != "Test2" || dest[1].FieldB != 456 || dest[1].FieldC != "" {
		t.Errorf("Expected dest to be %+v, but got %+v", []DestStruct{{FieldA: &src[0].FieldA, FieldB: src[0].FieldB, FieldC: ""}, {FieldA: &src[1].FieldA, FieldB: src[1].FieldB, FieldC: ""}}, dest)
	}
}

func Test_ptrString(t *testing.T) {
	src := []SourceStruct{
		{
			FieldA: "Test",
			FieldB: 123,
			FieldC: ToValue[string]("Hello"),
		},
		{
			FieldA: "Test2",
			FieldB: 456,
			FieldC: ToValue[string]("World"),
		},
	}
	var dest []DestStruct
	CopySlice(src, &dest)
	if len(dest) != 2 || *dest[0].FieldA != "Test" || dest[0].FieldB != 123 || dest[0].FieldC != "Hello" || *dest[1].FieldA != "Test2" || dest[1].FieldB != 456 || dest[1].FieldC != "World" {
		t.Errorf("Expected dest to be %+v, but got %+v", []DestStruct{{FieldA: &src[0].FieldA, FieldB: src[0].FieldB, FieldC: *src[0].FieldC}, {FieldA: &src[1].FieldA, FieldB: src[1].FieldB, FieldC: *src[1].FieldC}}, dest)
	}

}

type Src struct {
	I64        int64
	Iface      string
	Name       string
	Age        int
	F32        float32
	F64        float64
	I16        int16
	I8         int8
	I32        int32
	Array      []string
	Addresses  []Address
	Addresses2 *[]Address
}

type Dst struct {
	Name       string
	Age        int
	Iface      interface{}
	F32        float32
	I32        int32
	F64        float64
	I64        int64
	I8         int8
	I16        int16
	Array      []string
	Addresses  []Address2
	Addresses2 *[]Address2
}

func Test_MapStruct(t *testing.T) {
	src := Src{
		Name:  "Mehrdad",
		Age:   34,
		Iface: "string",
		F32:   1.1,
		F64:   64.64,
		I64:   64,
		I32:   32,
		I16:   16,
		I8:    8,
		Array: []string{"test"},
		Addresses: []Address{{
			Address: "TEST",
			Code:    nil,
		}},
		Addresses2: &[]Address{{
			Address: "TEST",
			Code:    nil,
		}},
	}
	dest := Dst{}

	Copy(src, &dest)
	assert.Equal(t, isEqual(src, dest), true)
}

type Src2 struct {
	I64  int64
	Name string
	Age  *int
	F32  float32
	F64  float64
	I16  *int16
	I8   int8
	I32  *int32
	B    *bool
	C    *complex128
}

type Dst2 struct {
	Name *string
	Age  int
	F32  *float32
	I32  int32
	F64  float64
	I64  *int64
	I8   int8
	I16  *int16
	B    bool
	C    complex128
}

func Test_MapStructWithNil(t *testing.T) {
	src := Src2{
		Name: "Mehrdad",
		Age:  ToValue[int](34),
		F32:  1.1,
		F64:  64.64,
		I64:  64,
		I32:  ToValue[int32](32),
		I16:  ToValue[int16](16),
		I8:   8,
		B:    ToValue[bool](true),
		C:    ToValue[complex128](13123123132132132132131232132),
	}
	dest := Dst2{}

	Copy(src, &dest)
	assert.Equal(t, isEqual2(src, dest), true)
}

type Src3 struct {
	Age      int
	Name     string
	Address  Address
	Address2 Address
}

type Dst3 struct {
	Name     *string
	Age      int
	Address  Address2
	Address2 *Address2
}
type Address struct {
	Address string
	Code    *string
}
type Address2 struct {
	Address string
	Code    *string
}

func Test_MapStructWithNestedStruct(t *testing.T) {
	src := Src3{
		Name: "Mehrdad",
		Age:  64,
		Address: Address{
			Address: "Hi",
			Code:    ToValue("asdasd"),
		},
		Address2: Address{
			Address: "Hi",
			Code:    ToValue("asdasd"),
		},
	}
	dest := Dst3{}

	Copy(src, &dest)
	assert.Equal(t, isEqual3(src, dest), true)
}

type SourceStruct1 struct {
	User User
}
type User struct {
	Name string
}
type User2 struct {
	Name string
}
type DestStruct1 struct {
	User User
}

func Test_MapStructWithNestedNilObject(t *testing.T) {
	src := SourceStruct1{
		User: User{
			Name: "NilMapper",
		},
	}

	var dest DestStruct1

	Copy(src, &dest)
	assert.Equal(t, src.User.Name, dest.User.Name)
}

type CreatePlant struct {
	Title                string
	ProductionLicenseID  string
	ProductionLocationID string
	M                    map[string]interface{}
}
type CreatePlanRequest struct {
	Title                string
	ProductionLicenseId  string
	ProductionLocationId string
	M                    map[string]interface{}
}

func Test_MapStructWithNestedGrpc(t *testing.T) {
	m := make(map[string]interface{})
	m["test"] = "test"
	src := CreatePlanRequest{
		Title:                "sada",
		ProductionLicenseId:  "asd",
		ProductionLocationId: "asd",
		M:                    m,
	}

	var dest CreatePlant

	Copy(src, &dest)
	assert.Equal(t, src.Title, dest.Title)
}

func isEqual(src Src, dst Dst) bool {
	return src.I64 == dst.I64 &&
		src.Name == dst.Name &&
		src.Age == dst.Age &&
		src.F32 == dst.F32 &&
		src.F64 == dst.F64 &&
		src.I16 == dst.I16 &&
		src.I8 == dst.I8 &&
		src.I32 == dst.I32
}
func isEqual2(src Src2, dst Dst2) bool {
	return src.I64 == *dst.I64 &&
		src.Name == *dst.Name &&
		*src.Age == dst.Age &&
		src.F32 == *dst.F32 &&
		src.F64 == dst.F64 &&
		*src.I16 == *dst.I16 &&
		src.I8 == dst.I8 &&
		*src.I32 == dst.I32 &&
		*src.B == dst.B &&
		*src.C == dst.C
}
func isEqual3(src Src3, dst Dst3) bool {
	return src.Name == *dst.Name && src.Age == dst.Age &&
		src.Address.Address == dst.Address.Address &&
		*src.Address.Code == *src.Address.Code &&
		src.Address2.Address == dst.Address2.Address
}
