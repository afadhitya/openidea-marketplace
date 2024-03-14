package signin

import "context"

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InportResponse struct {
	Token    string
	Username string
	Name     string
}
