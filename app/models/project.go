package models

import "time"

import (
	"fmt"
	"github.com/revel/revel"
)


type Project struct {
	ProjectId      int       `db:"id,"`
	ProjectName    string    `db:"name,size:64"`
	ProjectComment string    `db:"comment"`
	ShortComment   string
	CompanyName    string    `db:"company,size:128"`
	Position       string    `db:"position,size:16"`
	StartDate      string
	EndDate        string
	StartDateTime  time.Time `db:"start_date"`
	EndDateTime    time.Time `db:"end_date"`
}

func (p *Project) String() string {
	return fmt.Sprintf("Portfolio(%s)", p.ProjectName)
}

func (p *Project) Validate(v *revel.Validation) {
	v.Check(p.ProjectId,
		revel.Required{},
	)

	v.Check(p.ProjectName,
		revel.Required{},
		revel.MaxSize{64},
	)

	v.Check(p.ProjectName,
		revel.Required{},
		revel.MaxSize{128},
	)

	v.Check(p.Position,
		revel.Required{},
		revel.MaxSize{16},
	)

	p.StartDateTime = ValidDate(v, p.StartDate)
	p.EndDateTime = ValidDate(v, p.EndDate)
}

func ValidDate(v *revel.Validation, date string) time.Time {
	t, err := time.Parse("yyyymmdd", date)
	if err != nil {
		v.Error("Date Parsing Failed!!")
	}
	return t
}