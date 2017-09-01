package models

import "time"

import (
	"fmt"
	"github.com/revel/revel"
)


type Project struct {
	ProjectId      int       `db:"id"`
	ProjectName    string    `db:"name,size:64"`
	ProjectComment string    `db:"comment"`
	CompanyName    string    `db:"company,size:128"`
	Position       string    `db:"position,size:16"`
	StartDate      string    `db:"-"`
	EndDate        string    `db:"-"`
	StartDateTime  int64     `db:"start_date"`
	EndDateTime    int64     `db:"end_date"`
	LaunchUrl      string    `db:"-"`
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

func ValidDate(v *revel.Validation, date string) int64 {
	t, err := time.Parse("20170901", date)
	if err != nil {
		v.Error("Date Parsing Failed!!")
	}
	return t.UnixNano()
}