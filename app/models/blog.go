package models

import (
	"github.com/revel/revel"
)

type Blog struct {
	Id		int	`db:"id"`
	Title		string	`db:"title"`
	Content		string	`db:"content"`
	Category	string	`db:"category"`
	Type		string	`db:"type"`
	WriteDateTime	int64	`db:"write_date_time"`
	MainImage	string	`db:"main_image"`
	MainLinkUrl	string	`db:"main_link_url"`
}

type BlogReply struct {
	Id		int	`db:"id"`
	Comment		string	`db:"comment"`
	WriterName	string	`db:"writer_name"`
	Password	[]byte	`db:"password"`
	WriteDateTime	int64	`db:"write_date_time"`
	BlogId		int	`db:"blog_id"`
	ReplyGroup	int	`db:"reply_group"`
	Depth		int	`db:"depth"`
	Order		int	`db:"order"`

	InputPassword	string	`db:"-"`
}

func (b *Blog) Validate(v *revel.Validation){
	v.Check(b.Title,
		revel.Required{},
		revel.MaxSize{255})
	v.Check(b.Category,
		revel.Required{})
	v.Check(b.Type,
		revel.Required{})
}

func (br *BlogReply) Validate(v *revel.Validation){
	v.Check(br.Comment,
		revel.Required{})
	v.Check(br.BlogId,
		revel.Required{})
	v.Check(br.WriterName,
		revel.Required{},
		revel.MaxSize{32})
	v.Check(br.ReplyGroup,
		revel.Required{})
	v.Check(br.Depth,
		revel.Required{})

	ValidatePassword(v, br.InputPassword).Key("br.Password")
}