package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	company, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.Abort()

		switch err.Error() {
		case fmt.Sprintf("%s %s", NotFound, id):
			c.JSON(http.StatusNotFound, ErrMsg{
				Code:  http.StatusNotFound,
				Error: InexistentCompany,
			})
		case InvalidUUID:
			c.JSON(http.StatusBadRequest, ErrMsg{
				Code:  http.StatusBadRequest,
				Error: InvalidUUID,
			})
		default:
			c.JSON(http.StatusInternalServerError, ErrMsg{
				Code:  http.StatusInternalServerError,
				Error: GetError,
			})
		}

		return
	}

	c.JSON(http.StatusOK, company)
}
