package application

import (
	"context"

	"github.com/calloc134/golang-ddd-test/src/mutation/domain"
)

type IPostRepository interface {
	FindAll(context context.Context) ([]domain.Post, error)
	FindByUlid(context context.Context, ulid domain.UlidValue) (*domain.Post, error)
	Save(context context.Context, post *domain.Post) (*domain.Post, error)
}

type PostApplication struct {
	PostRepository IPostRepository
}

func NewPostApplication(postRepository IPostRepository) PostApplication {
	return PostApplication{PostRepository: postRepository}
}

func (pa *PostApplication) FindAll(context context.Context) ([]domain.Post, error) {
	return pa.PostRepository.FindAll(context)
}

func (pa *PostApplication) FindByUlid(context context.Context, ulidString string) (*domain.Post, error) {

	ulid, err := domain.NewULID(ulidString)

	if err != nil {
		return nil, err
	}

	return pa.PostRepository.FindByUlid(context, ulid)
}

func (pa *PostApplication) NewPost(context context.Context, userULIDString, titleString, contentString string) (*domain.Post, error) {
	// TODO: ログインに対応させる

	userULID, err := domain.NewULID(userULIDString)

	if err != nil {
		return nil, err
	}

	title, err := domain.NewTitle(titleString)

	if err != nil {
		return nil, err
	}

	content, err := domain.NewContent(contentString)

	if err != nil {
		return nil, err
	}

	post, err := domain.NewPost(userULID, title, content)

	if err != nil {
		return nil, err
	}

	return pa.PostRepository.Save(context, post)
}

func (pa *PostApplication) UpdateTitleByUlid(context context.Context, ulidString, titleString string) (*domain.Post, error) {

	title, err := domain.NewTitle(titleString)

	if err != nil {
		return nil, err
	}

	ulid, err := domain.NewULID(ulidString)

	if err != nil {
		return nil, err
	}

	post, err := pa.PostRepository.FindByUlid(context, ulid)

	if err != nil {
		return nil, err
	}

	if err := post.SetTitle(title); err != nil {
		return nil, err
	}

	return pa.PostRepository.Save(context, post)
}

func (pa *PostApplication) UpdateContentByUlid(context context.Context, ulidString, contentString string) (*domain.Post, error) {

	content, err := domain.NewContent(contentString)

	if err != nil {
		return nil, err
	}

	ulid, err := domain.NewULID(ulidString)

	if err != nil {
		return nil, err
	}

	post, err := pa.PostRepository.FindByUlid(context, ulid)

	if err != nil {
		return nil, err
	}

	if err := post.SetContent(content); err != nil {
		return nil, err
	}

	return pa.PostRepository.Save(context, post)
}
