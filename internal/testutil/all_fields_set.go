package testutil

import (
	"fmt"
	"reflect"
	"testing"
)

// AssertAllFieldsSet recursively traverses all fields of a struct and confirms
// that every field is not nil. The function treats the second argument as a
// pointer to a struct and uses reflection to check all fields including nested
// structs, pointers to scalars, and slices.
//
// If any nil fields are found, the function generates appropriate errors in the
// *testing.T and returns false. If no nil fields are found, returns true.
func AssertAllFieldsSet(t *testing.T, structPtr any) bool {
	t.Helper()

	if structPtr == nil {
		t.Error("struct pointer is nil")
		return false
	}

	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Ptr {
		t.Error("expected a pointer to a struct")
		return false
	}

	structVal := val.Elem()
	if structVal.Kind() != reflect.Struct {
		t.Error("expected a pointer to a struct")
		return false
	}

	return checkFieldsRecursively(t, structVal, "")
}

// checkFieldsRecursively performs the actual recursive field checking
func checkFieldsRecursively(t *testing.T, val reflect.Value, fieldPath string) bool {
	t.Helper()
	
	allFieldsSet := true
	valType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := valType.Field(i)
		
		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}
		
		currentPath := fieldType.Name
		if fieldPath != "" {
			currentPath = fieldPath + "." + fieldType.Name
		}

		if !checkField(t, field, currentPath) {
			allFieldsSet = false
		}
	}

	return allFieldsSet
}

// checkField checks a single field for nil values and handles different types
func checkField(t *testing.T, field reflect.Value, fieldPath string) bool {
	t.Helper()
	
	switch field.Kind() {
	case reflect.Ptr:
		if field.IsNil() {
			t.Errorf("field %s is nil", fieldPath)
			return false
		}
		// If it's a pointer to a struct, recursively check its fields
		if field.Elem().Kind() == reflect.Struct {
			return checkFieldsRecursively(t, field.Elem(), fieldPath)
		}
		
	case reflect.Slice:
		if field.IsNil() {
			t.Errorf("field %s (slice) is nil", fieldPath)
			return false
		}
		// Check each element in the slice
		for j := 0; j < field.Len(); j++ {
			elem := field.Index(j)
			elemPath := fmt.Sprintf("%s[%d]", fieldPath, j)
			if !checkField(t, elem, elemPath) {
				return false
			}
		}
		
	case reflect.Map:
		if field.IsNil() {
			t.Errorf("field %s (map) is nil", fieldPath)
			return false
		}
		// Check each value in the map
		for _, key := range field.MapKeys() {
			value := field.MapIndex(key)
			valuePath := fmt.Sprintf("%s[%v]", fieldPath, key.Interface())
			if !checkField(t, value, valuePath) {
				return false
			}
		}
		
	case reflect.Chan:
		if field.IsNil() {
			t.Errorf("field %s (channel) is nil", fieldPath)
			return false
		}
		
	case reflect.Func:
		if field.IsNil() {
			t.Errorf("field %s (function) is nil", fieldPath)
			return false
		}
		
	case reflect.Struct:
		// Recursively check nested struct fields
		return checkFieldsRecursively(t, field, fieldPath)
		
	case reflect.Interface:
		if field.IsNil() {
			t.Errorf("field %s (interface) is nil", fieldPath)
			return false
		}
		// Check the concrete value inside the interface
		if field.Elem().Kind() == reflect.Struct {
			return checkFieldsRecursively(t, field.Elem(), fieldPath)
		}
	}
	
	// For non-pointer, non-slice, non-map, non-channel, non-func, non-interface types, 
	// we don't check for nil as they cannot be nil (they have zero values instead)
	return true
}