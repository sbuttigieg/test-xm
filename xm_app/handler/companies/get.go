package companies

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	company, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.Abort()

		switch {
		case strings.Contains(err.Error(), NotFound):
			c.JSON(http.StatusNotFound, ErrMsg{
				Code:  http.StatusNotFound,
				Error: InexistentCompany,
			})
		case err.Error() == InvalidUUID:
			c.JSON(http.StatusBadRequest, ErrMsg{
				Code:  http.StatusBadRequest,
				Error: InvalidRequest,
			})
		default:
			c.JSON(http.StatusInternalServerError, ErrMsg{
				Code:  http.StatusInternalServerError,
				Error: GetError,
			})
		}

		return
	}

	c.JSON(http.StatusOK, OKMsg{
		Code: http.StatusOK,
		Data: company,
	})
}
