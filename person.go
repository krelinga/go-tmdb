package tmdb

type PersonId int

type Gender int

const (
	GenderUnknown   Gender = 0
	GenderFemale    Gender = 1
	GenderMale      Gender = 2
	GenderNonBinary Gender = 3
)

type Person struct {
	Id PersonId
	
	Adult *bool
	Gender *Gender
	KnownForDepartment *string
	Name *string
	Popularity *float64
	Profile *Image
}
