package tmdb

import "fmt"

type CompanyId int

type Company interface {
	Id() CompanyId

	// TODO: support Upgrade()

	CompanyData
}

type CompanyDataCol int

const (
	companyDataNone CompanyDataCol = iota
	companyDataMin
	// TODO: add more columns as needed
	companyDataMax
)

func (d CompanyDataCol) String() string {
	switch d {
	case companyDataNone:
		return "companyDataNone"
	case companyDataMin:
		return "companyDataMin"
	case companyDataMax:
		return "companyDataMax"
	default:
		return fmt.Sprintf("CompanyDataCol(%d)", d)
	}
}

func companyNoDataError(field string, col CompanyDataCol) error {
	var colPart string
	if col != companyDataNone {
		colPart = fmt.Sprintf(" with %s", col)
	}
	return fmt.Errorf("cannot access %s on CompanyData without calling Upgrade()%s first", field, colPart)
}

var (
	ErrCompanyNoDataLogo          = companyNoDataError("Logo", companyDataNone)
	ErrCompanyNoDataName          = companyNoDataError("Name", companyDataNone)
	ErrCompanyNoDataOriginCountry = companyNoDataError("OriginCountry", companyDataNone)
)

type CompanyData interface {
	Logo() Image
	Name() string
	OriginCountry() Country
}

type companyNoData struct{}

func (c companyNoData) Logo() Image {
	panic(ErrCompanyNoDataLogo)
}

func (c companyNoData) Name() string {
	panic(ErrCompanyNoDataName)
}

func (c companyNoData) OriginCountry() Country {
	panic(ErrCompanyNoDataOriginCountry)
}

type company struct {
	id CompanyId
	CompanyData
}

func (c *company) Id() CompanyId {
	return c.id
}
