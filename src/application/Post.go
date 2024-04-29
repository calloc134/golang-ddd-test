package application

import (
	"context"

	"github.com/calloc134/golang-ddd-test/src/domain"
)

type IPostRepository interface {
	FindAll(context context.Context) ([]domain.Post, error)
	FindByUlid(context context.Context, uuid string) (*domain.Post, error)
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

func (pa *PostApplication) NewPost(context context.Context, userULID, title, content string) (*domain.Post, error) {
	// TODO: ログインに対応させる
	post, err := domain.NewPost(userULID, title, content)

	if err != nil {
		return nil, err
	}

	return pa.PostRepository.Save(context, post)
}

func (pa *PostApplication) UpdateTitleByUlid(context context.Context, uuid string, title string) (*domain.Post, error) {
	post, err := pa.PostRepository.FindByUlid(context, uuid)

	if err != nil {
		return nil, err
	}

	err = post.SetTitle(title)

	if err != nil {
		return nil, err
	}

	return pa.PostRepository.Save(context, post)
}

func (pa *PostApplication) UpdateContentByUlid(context context.Context, uuid string, content string) (*domain.Post, error) {
	post, err := pa.PostRepository.FindByUlid(context, uuid)

	if err != nil {
		return nil, err
	}

	err = post.SetContent(content)

	if err != nil {
		return nil, err
	}

	return pa.PostRepository.Save(context, post)
}
