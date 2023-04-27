package handler

import "github.com/sbuttigieg/test-xm/xm_app/services"

const (
	IncorrectCredentials    = "incorrect credentials"
	MissingCredentialsError = "missing credentials"

	CreateError       = "error creating company"
	GetError          = "error getting company by id"
	InvalidRequest    = "error reading request"
	InvalidUUID       = "id is not a valid uuid"
	InexistentCompany = "company does not exist"
	NotFound          = "get company: : not found company by"
	Successful        = "successful"
)

type ErrMsg struct {
	Code  int
	Error string
}

type OKMsg struct {
	Code int
	Data interface{}
}

func New(serv services.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

type Handler struct {
	service services.Service
}
