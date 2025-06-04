package tmdb

type CreditId string

type Credit interface {
	Id() CreditId
	Person() Person
	OriginalName() string
}