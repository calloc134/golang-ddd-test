package domain

import (
	"errors"
)

type Post struct {
	ULID       UlidValue
	PostDetail *PostDetail
	Version    int
	UserULID   UlidValue
}

type PostDetail struct {
	ULID    UlidValue
	Title   TitleValue
	Content ContentValue
}

type TitleValue struct {
	value string
}

type ContentValue struct {
	value string
}

func (t TitleValue) String() string {
	return t.value
}

func (c ContentValue) String() string {
	return c.value
}

func NewTitle(value string) (TitleValue, error) {
	if value == "" {
		return TitleValue{}, ErrEmptyTitle
	}

	if len(value) > 255 {
		return TitleValue{}, ErrInvalidTitle
	}

	return TitleValue{value: value}, nil
}

func NewContent(value string) (ContentValue, error) {
	if value == "" {
		return ContentValue{}, ErrEmptyContent
	}

	if len(value) > 65535 {
		return ContentValue{}, ErrInvalidContent
	}

	return ContentValue{value: value}, nil
}

func NewPost(userUlid UlidValue, title TitleValue, content ContentValue) (*Post, error) {

	ulid, err := GenerateULID()

	if err != nil {
		return nil, err
	}

	post := &Post{
		ULID:     ulid,
		Version:  0,
		UserULID: userUlid,
		PostDetail: &PostDetail{
			Title:   title,
			Content: content,
		},
	}

	return post, nil
}

func (p *Post) SetTitle(title TitleValue) error {
	ulid, err := GenerateULID()

	if err != nil {
		return err
	}

	postDetail := &PostDetail{
		ULID:    ulid,
		Title:   title,
		Content: p.PostDetail.Content,
	}

	p.PostDetail = postDetail
	return nil
}

func (p *Post) SetContent(content ContentValue) error {

	ulid, err := GenerateULID()

	if err != nil {
		return err
	}

	postDetail := &PostDetail{
		ULID:    ulid,
		Title:   p.PostDetail.Title,
		Content: content,
	}

	p.PostDetail = postDetail

	return nil
}

var (
	ErrEmptyTitle      = errors.New("title is empty")
	ErrInvalidTitle    = errors.New("title is invalid")
	ErrEmptyContent    = errors.New("content is empty")
	ErrInvalidContent  = errors.New("content is invalid")
	ErrEmptyUserUlid   = errors.New("user ulid is empty")
	ErrInvalidUserUlid = errors.New("user ulid is invalid")
)
