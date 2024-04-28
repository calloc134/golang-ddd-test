package schemas

import "github.com/uptrace/bun"

type CommentTable struct {
	bun.BaseModel     `bun:"table:comments"`
	CommentULID       string              `bun:",pk,type:varchar(26)"`
	CommentDetail     *CommentDetailTable `bun:"rel:belongs-to,join:comment_detail_ulid=comment_detail_ulid"`
	CommentDetailULID string
	Version           int
}

type CommentDetailTable struct {
	bun.BaseModel     `bun:"table:comment_details"`
	CommentDetailULID string `bun:",pk,type:varchar(26)"`
	Content           string `bun:"type:text"`
	UserULID          string `bun:"type:varchar(26)"`
	PostULID          string `bun:"type:varchar(26)"`
}
