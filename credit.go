package tmdb

type CreditId string
type CastId int

type CreditKey struct {
	Id CreditId
}

type Credit struct {
	CreditKey

	Person *Person
	// TODO: does this really need to be a pointer?
	// They are always set for movie cast & crew.
	OriginalName *string

	// Set for movie cast.
	CastId *CastId
	Character *string
	Order *int

	// Set for movie crew.
	Department *string
	Job *string
}