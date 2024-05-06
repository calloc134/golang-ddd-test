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
		ulid, err := domain.NewULID(postTable.PostULID)

		if err != nil {
			return nil, err
		}

		userUlid, err := domain.NewULID(postTable.UserULID)

		if err != nil {
			return nil, err
		}

		title, err := domain.NewTitle(postTable.PostDetail.Title)

		if err != nil {
			return nil, err
		}

		content, err := domain.NewContent(postTable.PostDetail.Content)

		if err != nil {
			return nil, err
		}

		posts = append(posts, domain.Post{
			ULID: ulid,
			PostDetail: &domain.PostDetail{
				Title:   title,
				Content: content,
			},
			Version:  postTable.Version,
			UserULID: userUlid,
		})
	}

	return posts, nil
}

func (pr PostRepository) FindByUlid(context context.Context, ulid domain.UlidValue) (*domain.Post, error) {

	postTable := schemas.PostTable{}

	err := pr.db.NewSelect().Model(&postTable).Where("post_ulid = ?", ulid.String()).Relation("PostDetail").Scan(context)

	if err != nil {
		return nil, err
	}

	ulid, err = domain.NewULID(postTable.PostULID)

	if err != nil {
		return nil, err
	}

	userUlid, err := domain.NewULID(postTable.UserULID)

	if err != nil {
		return nil, err
	}

	title, err := domain.NewTitle(postTable.PostDetail.Title)

	if err != nil {
		return nil, err
	}

	content, err := domain.NewContent(postTable.PostDetail.Content)

	if err != nil {
		return nil, err
	}

	return &domain.Post{
		ULID: ulid,
		PostDetail: &domain.PostDetail{
			Title:   title,
			Content: content,
		},
		Version:  postTable.Version,
		UserULID: userUlid,
	}, nil
}

func (pr PostRepository) Save(context context.Context, post *domain.Post) (*domain.Post, error) {

	tx, err := pr.db.BeginTx(context, nil)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var version int
	versionCheckErr := pr.db.NewSelect().ColumnExpr("version").Model(&schemas.PostTable{}).Where("post_ulid = ?", post.ULID.String()).Scan(context, &version)

	// データがあってVersionが一致しない場合はエラー
	if versionCheckErr != nil && version != post.Version {
		return nil, errors.New("楽観的ロックエラー")
	}

	postTable := schemas.PostTable{
		PostULID:       post.ULID.String(),
		UserULID:       post.UserULID.String(),
		Version:        post.Version + 1,
		PostDetailULID: post.PostDetail.ULID.String(),
	}

	postDetailTable := schemas.PostDetailTable{
		PostDetailULID: post.PostDetail.ULID.String(),
		Title:          post.PostDetail.Title.String(),
		Content:        post.PostDetail.Content.String(),
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
