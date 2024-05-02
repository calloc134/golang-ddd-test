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

func NewPost(userUlidString string, titleString string, contentString string) (*Post, error) {

	// ULIDのバリデーションはここでやってしまう
	userUlid, err := NewULID(userUlidString)

	if err != nil {
		return nil, err
	}

	title, err := NewTitle(titleString)

	if err != nil {
		return nil, err
	}

	content, err := NewContent(contentString)

	if err != nil {
		return nil, err
	}

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

func (p *Post) SetTitle(titleString string) error {

	ulid, err := GenerateULID()

	if err != nil {
		return err
	}

	title, err := NewTitle(titleString)

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

func (p *Post) SetContent(contentString string) error {

	ulid, err := GenerateULID()

	if err != nil {
		return err
	}

	content, err := NewContent(contentString)

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
