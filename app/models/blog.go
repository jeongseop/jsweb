package models

import  (
	"fmt"
	"github.com/revel/revel"
)


type Blog struct {
	BlogId         int       `db:"id"`
	BlogName    string    `db:"name,size:64"`
	BlogComment string    `db:"comment"`
	Category    string    `db:"company,size:128"`
	Position       string    `db:"position,size:16"`
	StartDateTime  int64     `db:"start_date"`
	EndDateTime    int64     `db:"end_date"`
	LaunchUrl      string    `db:"launch_url"`

	StartDate      string    `db:"-"`
	EndDate        string    `db:"-"`
	CommentList  []string    `db:"-"`
}

func (p *Blog) String() string {
	return fmt.Sprintf("Portfolio(%s)", p.BlogName)
}

func (p *Blog) Validate(v *revel.Validation) {
	v.Check(p.BlogName,
		revel.Required{},
		revel.MaxSize{64},
	).Message("ProjectName 입력 오류!!")
	v.Check(p.BlogComment,
		revel.Required{},
		revel.MaxSize{2048},
	).Message("ProjectComment 입력 오류!!")
	v.Check(p.Category,
		revel.Required{},
		revel.MaxSize{128},
	).Message("CompanyName 입력 오류!!")
	v.Check(p.Position,
		revel.Required{},
		revel.MaxSize{16},
	).Message("Position 입력 오류!!")
	v.Check(p.LaunchUrl,
		revel.MaxSize{255},
	).Message("LaunchUrl 입력 오류!!")

	var err error
	if p.StartDateTime, err = ValidDate(p.StartDate);  err != nil {
		v.Error("StartDate 입력 오류!!")
	}
	if p.EndDateTime, err = ValidDate(p.EndDate);  err != nil {
		v.Error("EndDate 입력 오류!!")
	}
}