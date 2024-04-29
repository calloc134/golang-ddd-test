package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type Post struct {
	ULID       string
	PostDetail *PostDetail
	Version    int
	UserULID   string
}

type PostDetail struct {
	ULID    string
	Title   string
	Content string
}

func NewPost(userUlid, title, content string) (*Post, error) {

	// ULIDのバリデーションはここでやってしまう
	if userUlid == "" {
		return nil, ErrEmptyUserUlid
	}

	if len(userUlid) != 26 {
		return nil, ErrInvalidUserUlid
	}

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return nil, err
	}

	post := &Post{
		ULID:     ulid.String(),
		Version:  0,
		UserULID: userUlid,
		PostDetail: &PostDetail{
			Title:   "",
			Content: "",
		},
	}

	if err := post.SetTitle(title); err != nil {
		return nil, err
	}

	if err := post.SetContent(content); err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Post) SetTitle(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}

	if len(title) > 255 {
		return ErrInvalidTitle
	}

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return err
	}

	postDetail := &PostDetail{
		ULID:    ulid.String(),
		Title:   title,
		Content: p.PostDetail.Content,
	}

	p.PostDetail = postDetail
	return nil
}

func (p *Post) SetContent(content string) error {
	if content == "" {
		return ErrEmptyContent
	}

	if len(content) > 65535 {
		return ErrInvalidContent
	}

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return err
	}

	postDetail := &PostDetail{
		ULID:    ulid.String(),
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
