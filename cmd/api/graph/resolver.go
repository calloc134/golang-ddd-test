package graph

import "github.com/calloc134/golang-ddd-test/src/application"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserApplication application.UserApplication
}

func NewResolver(userApplication application.UserApplication) *Resolver {
	return &Resolver{UserApplication: userApplication}
}