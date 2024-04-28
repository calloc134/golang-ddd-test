package schemas

import "github.com/uptrace/bun"

type PostTable struct {
	bun.BaseModel  `bun:"table:posts"`
	PostULID       string           `bun:",pk,type:varchar(26)"`
	PostDetail     *PostDetailTable `bun:"rel:belongs-to,join:post_detail_ulid=post_detail_ulid"`
	PostDetailULID string
	Version        int
}

type PostDetailTable struct {
	bun.BaseModel  `bun:"table:post_details"`
	PostDetailULID string `bun:",pk,type:varchar(26)"`
	Title          string `bun:"type:varchar(255)"`
	Content        string `bun:"type:text"`
	UserULID       string `bun:"type:varchar(26)"`
}
