package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type Member struct {
	UserId             string
	Password	   string
	Email		   string
	HashedPassword     []byte
}

func (m *Member) String() string {
	return fmt.Sprintf("User(%s)", m.UserId)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (memb *Member) Validate(v *revel.Validation) {
	v.Check(memb.UserId,
		revel.Required{},
		revel.MaxSize{16},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, memb.Password).
		Key("memb.Password")

	v.Check(memb.Email,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{5},
	)
}