package handler

import "github.com/sbuttigieg/test-xm/xm_app/services"

const (
	IncorrectCredentials    = "incorrect credentials"
	MissingCredentialsError = "missing credentials"
	GetError                = "error getting company by id"
	InvalidUUID             = "id is not a valid uuid"
	InexistentCompany       = "company does not exist"
	NotFound                = "get company: : not found company by"
)

type ErrMsg struct {
	Code  int
	Error string
}

func New(serv services.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

type Handler struct {
	service services.Service
}
