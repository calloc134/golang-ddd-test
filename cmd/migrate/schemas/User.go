package schemas

import "github.com/uptrace/bun"

type UserTable struct {
	bun.BaseModel `bun:"table:users"`
	ULID          string           `bun:",pk,type:vargchar(26)"`
	UserDetail    *UserDetailTable `bun:"rel:has-one,join:ulid=user_id"`
	Version       int
}

type UserDetailTable struct {
	bun.BaseModel `bun:"table:user_details"`
	Name          string `bun:"type:varchar(255)"`
	Age           int
	UserID        string `bun:"type:varchar(26)"`
}
