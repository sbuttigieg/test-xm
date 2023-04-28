package users

import (
	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/services/users"
)

const (
	CreateError          = "error creating user"
	GetTokenError        = "error getting token"
	IncorrectCredentials = "incorrect credentials"
	InvalidRequest       = "invalid request"
	NotFound             = "not found"
	Successful           = "successful"
)

type ErrMsg struct {
	Code  int
	Error string
}

type OKMsg struct {
	Code int
	Data interface{}
}

func New(cfg *app.Config, serv users.Service) *Handler {
	return &Handler{
		config:  cfg,
		service: serv,
	}
}

type Handler struct {
	config  *app.Config
	service users.Service
}
