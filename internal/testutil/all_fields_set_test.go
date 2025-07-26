package testutil

import (
	"fmt"
	"testing"
)

// Test structures for testing the AssertAllFieldsSet function
type SimpleStruct struct {
	StringField string
	IntField    int
	BoolField   bool
}

type StructWithPointers struct {
	StringPtr *string
	IntPtr    *int
	BoolPtr   *bool
}

type StructWithSlices struct {
	StringSlice []string
	IntSlice    []int
	PtrSlice    []*string
}

type StructWithMapsAndChannels struct {
	StringMap map[string]string
	IntMap    map[int]*string
	Channel   chan string
	Func      func() string
}

type NestedStruct struct {
	Simple   SimpleStruct
	WithPtrs *StructWithPointers
	Slice    []*SimpleStruct
}

type DeepNestedStruct struct {
	Level1 *NestedStruct
	Level2 *StructWithSlices
}

// Interface types for testing interface handling
type TestInterface interface {
	DoSomething()
}

type TestInterfaceValue interface {
	DoSomethingValue()
}

type StructWithInterface struct {
	MyInterface TestInterface
}

type StructWithInterfaceValue struct {
	MyInterface TestInterfaceValue
}

type InterfaceImpl struct {
	Name *string
	Age  *int
}

func (i *InterfaceImpl) DoSomething() {}

type InterfaceImplValue struct {
	Name *string
	Age  *int
}

func (i InterfaceImplValue) DoSomethingValue() {}

type InterfaceImplWithNested struct {
	Basic *SimpleStruct
	Name  *string
}

func (i *InterfaceImplWithNested) DoSomething() {}

func TestAssertAllFieldsSet_SimpleStruct(t *testing.T) {
	// Test with a simple struct (no pointers/slices) - should always pass
	simple := &SimpleStruct{
		StringField: "test",
		IntField:    42,
		BoolField:   true,
	}

	result := AssertAllFieldsSet(t, simple)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for simple struct")
	}
}

func TestAssertAllFieldsSet_StructWithPointers_AllSet(t *testing.T) {
	// Test with all pointers set
	str := "test"
	num := 42
	flag := true

	withPtrs := &StructWithPointers{
		StringPtr: &str,
		IntPtr:    &num,
		BoolPtr:   &flag,
	}

	result := AssertAllFieldsSet(t, withPtrs)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true when all pointers are set")
	}
}

func TestAssertAllFieldsSet_StructWithPointers_SomeNil(t *testing.T) {
	// Test with some nil pointers - should fail
	str := "test"

	withPtrs := &StructWithPointers{
		StringPtr: &str,
		IntPtr:    nil, // This should cause failure
		BoolPtr:   nil, // This should cause failure
	}

	// Capture the test errors by creating a mock test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withPtrs)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false when some pointers are nil")
	}
}

func TestAssertAllFieldsSet_StructWithSlices_AllSet(t *testing.T) {
	// Test with all slices properly initialized
	str1 := "test1"
	str2 := "test2"

	withSlices := &StructWithSlices{
		StringSlice: []string{"a", "b", "c"},
		IntSlice:    []int{1, 2, 3},
		PtrSlice:    []*string{&str1, &str2},
	}

	result := AssertAllFieldsSet(t, withSlices)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true when all slices are set")
	}
}

func TestAssertAllFieldsSet_StructWithSlices_SomeNil(t *testing.T) {
	// Test with some nil slices - should fail
	withSlices := &StructWithSlices{
		StringSlice: []string{"a", "b", "c"},
		IntSlice:    nil, // This should cause failure
		PtrSlice:    nil, // This should cause failure
	}

	// Capture the test errors by creating a mock test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withSlices)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false when some slices are nil")
	}
}

func TestAssertAllFieldsSet_StructWithSlices_NilElementsInSlice(t *testing.T) {
	// Test with nil elements inside slices - should fail
	str1 := "test1"

	withSlices := &StructWithSlices{
		StringSlice: []string{"a", "b", "c"},
		IntSlice:    []int{1, 2, 3},
		PtrSlice:    []*string{&str1, nil}, // nil element should cause failure
	}

	// Capture the test errors by creating a mock test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withSlices)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false when slice contains nil elements")
	}
}

func TestAssertAllFieldsSet_NestedStruct_AllSet(t *testing.T) {
	// Test with nested struct, all fields set
	str := "test"
	num := 42
	flag := true

	nested := &NestedStruct{
		Simple: SimpleStruct{
			StringField: "simple",
			IntField:    10,
			BoolField:   false,
		},
		WithPtrs: &StructWithPointers{
			StringPtr: &str,
			IntPtr:    &num,
			BoolPtr:   &flag,
		},
		Slice: []*SimpleStruct{
			{StringField: "slice1", IntField: 1, BoolField: true},
			{StringField: "slice2", IntField: 2, BoolField: false},
		},
	}

	result := AssertAllFieldsSet(t, nested)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for properly nested struct")
	}
}

func TestAssertAllFieldsSet_NestedStruct_NilPointer(t *testing.T) {
	// Test with nested struct containing nil pointer - should fail
	nested := &NestedStruct{
		Simple: SimpleStruct{
			StringField: "simple",
			IntField:    10,
			BoolField:   false,
		},
		WithPtrs: nil, // This should cause failure
		Slice: []*SimpleStruct{
			{StringField: "slice1", IntField: 1, BoolField: true},
		},
	}

	// Capture the test errors by creating a mock test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, nested)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false when nested pointer is nil")
	}
}

func TestAssertAllFieldsSet_DeepNested_AllSet(t *testing.T) {
	// Test with deeply nested structure, all fields set
	str := "test"
	num := 42
	flag := true

	deep := &DeepNestedStruct{
		Level1: &NestedStruct{
			Simple: SimpleStruct{
				StringField: "simple",
				IntField:    10,
				BoolField:   false,
			},
			WithPtrs: &StructWithPointers{
				StringPtr: &str,
				IntPtr:    &num,
				BoolPtr:   &flag,
			},
			Slice: []*SimpleStruct{
				{StringField: "slice1", IntField: 1, BoolField: true},
			},
		},
		Level2: &StructWithSlices{
			StringSlice: []string{"a", "b"},
			IntSlice:    []int{1, 2},
			PtrSlice:    []*string{&str},
		},
	}

	result := AssertAllFieldsSet(t, deep)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for properly deep nested struct")
	}
}

func TestAssertAllFieldsSet_NilInput(t *testing.T) {
	// Test with nil input - should fail
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, nil)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for nil input")
	}
}

func TestAssertAllFieldsSet_NonPointerInput(t *testing.T) {
	// Test with non-pointer input - should fail
	simple := SimpleStruct{
		StringField: "test",
		IntField:    42,
		BoolField:   true,
	}

	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, simple) // Not a pointer

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for non-pointer input")
	}
}

func TestAssertAllFieldsSet_NonStructPointer(t *testing.T) {
	// Test with pointer to non-struct - should fail
	str := "test"

	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, &str) // Pointer to string, not struct

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for pointer to non-struct")
	}
}

// Test with empty slice (should pass)
func TestAssertAllFieldsSet_EmptySlice(t *testing.T) {
	withSlices := &StructWithSlices{
		StringSlice: []string{}, // Empty slice should be ok
		IntSlice:    []int{},
		PtrSlice:    []*string{},
	}

	result := AssertAllFieldsSet(t, withSlices)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for empty slices")
	}
}

func TestAssertAllFieldsSet_MapsAndChannels_AllSet(t *testing.T) {
	// Test with maps and channels set
	str := "test"

	withExtra := &StructWithMapsAndChannels{
		StringMap: map[string]string{"key": "value"},
		IntMap:    map[int]*string{1: &str},
		Channel:   make(chan string),
		Func:      func() string { return "test" },
	}

	result := AssertAllFieldsSet(t, withExtra)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true when maps and channels are set")
	}
}

func TestAssertAllFieldsSet_MapsAndChannels_SomeNil(t *testing.T) {
	// Test with some nil maps and channels - should fail
	withExtra := &StructWithMapsAndChannels{
		StringMap: map[string]string{"key": "value"},
		IntMap:    nil, // This should cause failure
		Channel:   nil, // This should cause failure
		Func:      nil, // This should cause failure
	}

	// Capture the test errors by creating a mock test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withExtra)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false when some maps/channels/functions are nil")
	}
}

func TestAssertAllFieldsSet_MapWithNilValues(t *testing.T) {
	// Test with map containing nil values - should fail
	withExtra := &StructWithMapsAndChannels{
		StringMap: map[string]string{"key": "value"},
		IntMap:    map[int]*string{1: nil}, // nil value in map should cause failure
		Channel:   make(chan string),
		Func:      func() string { return "test" },
	}

	// Capture the test errors by creating a mock test
	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withExtra)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false when map contains nil values")
	}
}

func TestAssertAllFieldsSet_LargeSliceIndices(t *testing.T) {
	// Test with larger slice indices to verify proper index formatting
	strs := make([]*string, 15)
	for i := range strs {
		s := fmt.Sprintf("string%d", i)
		strs[i] = &s
	}

	withSlices := &StructWithSlices{
		StringSlice: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"},
		IntSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		PtrSlice:    strs,
	}

	result := AssertAllFieldsSet(t, withSlices)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for large slices with all elements set")
	}

	// Now test with a nil element at a higher index
	strs[12] = nil // Set element at index 12 to nil

	mockT := &testing.T{}
	result2 := AssertAllFieldsSet(mockT, withSlices)
	if result2 {
		t.Error("Expected AssertAllFieldsSet to return false when slice has nil element at higher index")
	}
}

func TestAssertAllFieldsSet_Interface_PointerToStruct_AllSet(t *testing.T) {
	// Test interface containing a pointer to a struct with all fields set
	name := "test"
	age := 25

	impl := &InterfaceImpl{
		Name: &name,
		Age:  &age,
	}

	withInterface := &StructWithInterface{
		MyInterface: impl, // Interface holds a pointer to struct
	}

	result := AssertAllFieldsSet(t, withInterface)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for interface containing pointer to struct with all fields set")
	}
}

func TestAssertAllFieldsSet_Interface_PointerToStruct_WithNilFields(t *testing.T) {
	// Test interface containing a pointer to a struct with nil fields - should fail
	name := "test"

	impl := &InterfaceImpl{
		Name: &name,
		Age:  nil, // This should cause failure
	}

	withInterface := &StructWithInterface{
		MyInterface: impl, // Interface holds a pointer to struct with nil field
	}

	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withInterface)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for interface containing pointer to struct with nil fields")
	}
}

func TestAssertAllFieldsSet_Interface_StructValue_AllSet(t *testing.T) {
	// Test interface containing a struct value (not pointer) - should pass for non-nil fields
	name := "test"
	age := 25

	impl := InterfaceImplValue{
		Name: &name,
		Age:  &age,
	}

	withInterface := &StructWithInterfaceValue{
		MyInterface: impl, // Interface holds a struct value
	}

	result := AssertAllFieldsSet(t, withInterface)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for interface containing struct value with all fields set")
	}
}

func TestAssertAllFieldsSet_Interface_StructValue_WithNilFields(t *testing.T) {
	// Test interface containing a struct value with nil fields - should fail
	name := "test"

	impl := InterfaceImplValue{
		Name: &name,
		Age:  nil, // This should cause failure
	}

	withInterface := &StructWithInterfaceValue{
		MyInterface: impl, // Interface holds a struct value with nil field
	}

	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withInterface)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for interface containing struct value with nil fields")
	}
}

func TestAssertAllFieldsSet_Interface_NestedStruct_AllSet(t *testing.T) {
	// Test interface containing a pointer to a struct with nested struct fields
	simple := &SimpleStruct{
		StringField: "nested",
		IntField:    42,
		BoolField:   true,
	}
	name := "test"

	impl := &InterfaceImplWithNested{
		Basic: simple,
		Name:  &name,
	}

	withInterface := &StructWithInterface{
		MyInterface: impl,
	}

	result := AssertAllFieldsSet(t, withInterface)
	if !result {
		t.Error("Expected AssertAllFieldsSet to return true for interface with nested struct fields all set")
	}
}

func TestAssertAllFieldsSet_Interface_NestedStruct_WithNilFields(t *testing.T) {
	// Test interface containing a pointer to a struct with nil nested struct
	impl := &InterfaceImplWithNested{
		Basic: nil, // This should cause failure
		Name:  nil, // This should also cause failure
	}

	withInterface := &StructWithInterface{
		MyInterface: impl,
	}

	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withInterface)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for interface with nil nested struct fields")
	}
}

func TestAssertAllFieldsSet_Interface_NilInterface(t *testing.T) {
	// Test with nil interface - should fail
	withInterface := &StructWithInterface{
		MyInterface: nil, // Interface itself is nil
	}

	mockT := &testing.T{}
	result := AssertAllFieldsSet(mockT, withInterface)

	if result {
		t.Error("Expected AssertAllFieldsSet to return false for nil interface")
	}
}
