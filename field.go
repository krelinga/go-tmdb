package tmdb

import (
	"fmt"
	"math"
)

type Field[T any] interface {
	Get() (T, bool)
	GetDefault() T
	GetPanic() T
	Has() bool
}

type AllowedType interface {
	~bool | ~string | ~float64 | ~Array | ~Object
}

type ObjectField[T AllowedType] struct {
	Object  Object
	Key     string
	Default T
}

func (f ObjectField[T]) Get() (T, bool) {
	if f.Object == nil {
		return f.Default, false
	}
	anyVal, ok := f.Object[f.Key]
	if !ok {
		return f.Default, false
	}
	val, ok := anyVal.(T)
	if !ok {
		return f.Default, false
	}
	return val, true
}

func (f ObjectField[T]) GetDefault() T {
	val, _ := f.Get()
	return val
}

func (f ObjectField[T]) GetPanic() T {
	val, ok := f.Get()
	if !ok {
		panic(fmt.Sprintf("ObjectField.GetPanic: key %s not found or not of type %T", f.Key, val))
	}
	return val
}

func (f ObjectField[T]) Has() bool {
	_, ok := f.Get()
	return ok
}

func (f ObjectField[T]) WithDefault(defaultValue T) ObjectField[T] {
	f.Default = defaultValue
	return f
}

type WrappedField[In, Out any] struct {
	In      Field[In]
	Wrap    func(In) (Out, bool)
	Default Out
}

func (wf WrappedField[In, Out]) Get() (Out, bool) {
	inVal, ok := wf.In.Get()
	if !ok {
		return wf.Default, false
	}
	outVal, ok := wf.Wrap(inVal)
	if !ok {
		return wf.Default, false
	}
	return outVal, true
}

func (wf WrappedField[In, Out]) GetDefault() Out {
	val, _ := wf.Get()
	return val
}

func (wf WrappedField[In, Out]) GetPanic() Out {
	val, ok := wf.Get()
	if !ok {
		panic(fmt.Sprintf("WrappedField.GetPanic: key not found or not of type %T", val))
	}
	return val
}

func (wf WrappedField[In, Out]) Has() bool {
	_, ok := wf.Get()
	return ok
}

func (wf WrappedField[In, Out]) WithDefault(defaultValue Out) WrappedField[In, Out] {
	wf.Default = defaultValue
	return wf
}

func GetField[T AllowedType](obj Object, key string) ObjectField[T] {
	return ObjectField[T]{
		Object:  obj,
		Key:     key,
		Default: *new(T), // Default to zero value of T
	}
}

func wrapObject[T ~Object](obj Object) (T, bool) {
	return T(obj), true
}

func canFloat64BeInt32(f float64) bool {
	return f >= float64(int32(math.MinInt32)) && f <= float64(int32(math.MaxInt32)) && f == float64(int32(f))
}

func wrapNumberAsInt32(f float64) (int32, bool) {
	if canFloat64BeInt32(f) {
		return int32(f), true
	}
	return 0, false
}

func wrapArray[Out any](in Array) ([]Out, bool) {
	out := make([]Out, len(in))
	for i, v := range in {
		if outVal, ok := v.(Out); ok {
			out[i] = outVal
		} else {
			return nil, false
		}
	}
	return out, true
}

func WrapSubObject[Out ~Object](in Object, key string) Field[Out] {
	return WrappedField[Object, Out]{
		In: ObjectField[Object]{
			Object: in,
			Key:    key,
		},
		Wrap: wrapObject[Out],
	}
}

func WrapNumberAsInt32(in Object, key string) Field[int32] {
	return WrappedField[float64, int32]{
		In: ObjectField[float64]{
			Object: in,
			Key:    key,
		},
		Wrap: wrapNumberAsInt32,
	}
}

func WrapArray[Out any](in Object, key string) Field[[]Out] {
	return WrappedField[Array, []Out]{
		In: ObjectField[Array]{
			Object: in,
			Key:    key,
		},
		Wrap: wrapArray[Out],
	}
}