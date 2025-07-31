package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

type tester interface {
	test(t *testing.T, obj tmdb.Object)
	String() string
}

type testerImpl[ObjType ~tmdb.Object, ValueType comparable] struct {
	name string
	obj  ObjType
	call func(ObjType) (ValueType, error)
	want ValueType
}

func newTesterImpl[ObjType ~tmdb.Object, ValueType comparable](
	name string,
	obj ObjType,
	call func(ObjType) (ValueType, error),
	want ValueType,
) testerImpl[ObjType, ValueType] {
	return testerImpl[ObjType, ValueType]{
		name: name,
		obj:  obj,
		call: call,
		want: want,
	}
}

func (ti testerImpl[ObjType, ValueType]) String() string {
	return ti.name
}

func (ti testerImpl[ObjType, ValueType]) test(t *testing.T, obj tmdb.Object) {
	if obj == nil {
		t.Errorf("expected non-nil object, got nil")
		return
	}
	got, err := ti.call(ObjType(obj))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	if got != ti.want {
		t.Errorf("got %v, want %v", got, ti.want)
	}
}

func TestGetMovie(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	fightClub, err := tmdb.GetMovie(context.Background(), client, 550)
	if err != nil {
		t.Fatalf("failed to get movie: %v", err)
	}
	tests := []tester{
		newTesterImpl("Adult", fightClub, tmdb.Movie.Adult, false),
		newTesterImpl("ID", fightClub, tmdb.Movie.ID, int32(550)),
	}
	for _, test := range tests {
		t.Run(test.String(), func(t *testing.T) {
			test.test(t, fightClub)
		})
	}
}