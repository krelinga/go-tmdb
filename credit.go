package tmdb

import "github.com/krelinga/go-jsonflex"

type Credit Object

func (c Credit) Adult() (bool, error) {
	return jsonflex.GetField(c, "adult", jsonflex.AsBool())
}

func (c Credit) Gender() (int32, error) {
	return jsonflex.GetField(c, "gender", jsonflex.AsInt32())
}

func (c Credit) ID() (int32, error) {
	return jsonflex.GetField(c, "id", jsonflex.AsInt32())
}

func (c Credit) KnownForDepartment() (string, error) {
	return jsonflex.GetField(c, "known_for_department", jsonflex.AsString())
}

func (c Credit) Name() (string, error) {
	return jsonflex.GetField(c, "name", jsonflex.AsString())
}

func (c Credit) OriginalName() (string, error) {
	return jsonflex.GetField(c, "original_name", jsonflex.AsString())
}

func (c Credit) Popularity() (float64, error) {
	return jsonflex.GetField(c, "popularity", jsonflex.AsFloat64())
}

func (c Credit) ProfilePath() (string, error) {
	return jsonflex.GetField(c, "profile_path", jsonflex.AsString())
}

func (c Credit) CastID() (int32, error) {
	return jsonflex.GetField(c, "cast_id", jsonflex.AsInt32())
}

func (c Credit) Character() (string, error) {
	return jsonflex.GetField(c, "character", jsonflex.AsString())
}

func (c Credit) CreditID() (string, error) {
	return jsonflex.GetField(c, "credit_id", jsonflex.AsString())
}

func (c Credit) Order() (int32, error) {
	return jsonflex.GetField(c, "order", jsonflex.AsInt32())
}

func (c Credit) Department() (string, error) {
	return jsonflex.GetField(c, "department", jsonflex.AsString())
}

func (c Credit) Job() (string, error) {
	return jsonflex.GetField(c, "job", jsonflex.AsString())
}

type Credits Object

func (c Credits) ID() (int32, error) {
	return jsonflex.GetField(c, "id", jsonflex.AsInt32())
}

func (c Credits) Cast() ([]Credit, error) {
	return jsonflex.GetField(c, "cast", jsonflex.AsArray(jsonflex.AsObject[Credit]()))
}

func (c Credits) Crew() ([]Credit, error) {
	return jsonflex.GetField(c, "crew", jsonflex.AsArray(jsonflex.AsObject[Credit]()))
}
