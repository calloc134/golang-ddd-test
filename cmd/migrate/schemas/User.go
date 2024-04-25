package schemas

import "github.com/uptrace/bun"

type UserTable struct {
	bun.BaseModel `bun:"table:users"`
	ULID 		string `bun:"pk,type:vargchar(26)"`
	Name 		string `bun:"type:varchar(255)"`
	Age 		int 
	Version 	int `bun:"version"`
}

