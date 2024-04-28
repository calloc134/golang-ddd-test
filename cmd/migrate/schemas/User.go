package schemas

import "github.com/uptrace/bun"

type UserTable struct {
	bun.BaseModel  `bun:"table:users"`
	UserULID       string           `bun:",pk,type:vargchar(26)"`
	UserDetail     *UserDetailTable `bun:"rel:belongs-to,join:user_detail_ulid=user_detail_ulid"`
	UserDetailULID string
	Version        int
}

type UserDetailTable struct {
	bun.BaseModel  `bun:"table:user_details"`
	UserDetailULID string `bun:",pk,type:varchar(26)"`
	Name           string `bun:"type:varchar(255)"`
	Age            int
}
