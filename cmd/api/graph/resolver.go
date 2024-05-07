package graph

import "github.com/calloc134/golang-ddd-test/src/mutation/application"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserApplication application.UserApplication
	PostApplication application.PostApplication
}

func NewResolver(userApplication application.UserApplication, postApplication application.PostApplication) *Resolver {
	return &Resolver{
		UserApplication: userApplication,
		PostApplication: postApplication,
	}
}
