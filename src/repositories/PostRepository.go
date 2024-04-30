package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/calloc134/golang-ddd-test/cmd/migrate/schemas"
	"github.com/calloc134/golang-ddd-test/src/application"
	"github.com/calloc134/golang-ddd-test/src/domain"
	"github.com/uptrace/bun"
)

type PostRepository struct {
	db *bun.DB
}

func NewPostRepository(db *bun.DB) application.IPostRepository {
	return PostRepository{db: db}
}

func (pr PostRepository) FindAll(context context.Context) ([]domain.Post, error) {

	postTables := []schemas.PostTable{}

	err := pr.db.NewSelect().Model(&postTables).Relation("PostDetail").Scan(context)

	if err != nil {
		return nil, err
	}

	var posts []domain.Post

	for _, postTable := range postTables {
		posts = append(posts, domain.Post{
			ULID: postTable.PostULID,
			PostDetail: &domain.PostDetail{
				Title:   postTable.PostDetail.Title,
				Content: postTable.PostDetail.Content,
			},
			Version:  postTable.Version,
			UserULID: postTable.UserULID,
		})
	}

	return posts, nil
}

func (pr PostRepository) FindByUlid(context context.Context, uuid string) (*domain.Post, error) {

	postTable := schemas.PostTable{}

	err := pr.db.NewSelect().Model(&postTable).Where("post_ulid = ?", uuid).Relation("PostDetail").Scan(context)

	if err != nil {
		return nil, err
	}

	return &domain.Post{
		ULID: postTable.PostULID,
		PostDetail: &domain.PostDetail{
			Title:   postTable.PostDetail.Title,
			Content: postTable.PostDetail.Content,
		},
		Version:  postTable.Version,
		UserULID: postTable.UserULID,
	}, nil
}

func (pr PostRepository) Save(context context.Context, post *domain.Post) (*domain.Post, error) {

	tx, err := pr.db.BeginTx(context, nil)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var version int
	versionCheckErr := pr.db.NewSelect().ColumnExpr("version").Model(&schemas.PostTable{}).Where("post_ulid = ?", post.ULID).Scan(context, &version)

	// データがあってVersionが一致しない場合はエラー
	if versionCheckErr != nil && version != post.Version {
		return nil, errors.New("楽観的ロックエラー")
	}

	postTable := schemas.PostTable{
		PostULID:       post.ULID,
		UserULID:       post.UserULID,
		Version:        post.Version + 1,
		PostDetailULID: post.PostDetail.ULID,
	}

	postDetailTable := schemas.PostDetailTable{
		PostDetailULID: post.PostDetail.ULID,
		Title:          post.PostDetail.Title,
		Content:        post.PostDetail.Content,
	}

	_, err = tx.NewInsert().Model(&postDetailTable).Exec(context)

	if err != nil {
		return nil, err
	}

	if versionCheckErr != nil && versionCheckErr.Error() == "sql: no rows in result set" {
		fmt.Println("insert")
		_, err = tx.NewInsert().Model(&postTable).Exec(context)

		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("update")
		_, err = tx.NewUpdate().Model(&postTable).Where("version = ?, post_detail_ulid = ?, user_ulid = ?", postTable.Version, postTable.PostDetailULID, postTable.UserULID).Exec(context)

		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return post, nil
}
