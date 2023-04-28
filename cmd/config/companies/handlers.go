package companies

import (
	"github.com/sbuttigieg/test-xm/xm_app/handler/companies"
	services "github.com/sbuttigieg/test-xm/xm_app/services/companies"
)

func NewHandlers(service services.Service) *companies.Handler {
	return companies.New(service)
}
