package schemas

import "github.com/uptrace/bun"

type UserTable struct {
	bun.BaseModel `bun:"table:users,alias:users"`
	UUID 		string `bun:"id,pk,type:uuid"`
	Name 		string `bun:"type:varchar(255)"`
	Age 		int
}

