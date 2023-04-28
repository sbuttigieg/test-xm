package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (h *Handler) Create(c *gin.Context) {
	var req *models.User

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, ErrMsg{
			Code:  http.StatusBadRequest,
			Error: InvalidRequest,
		})

		return
	}

	if err := req.HashPassword(req.Password); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, ErrMsg{
			Code:  http.StatusInternalServerError,
			Error: CreateError,
		})

		return
	}

	id, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, ErrMsg{
			Code:  http.StatusInternalServerError,
			Error: CreateError,
		})

		return
	}

	c.JSON(http.StatusCreated, OKMsg{
		Code: http.StatusCreated,
		Data: fmt.Sprintf("%s: %s", Successful, id),
	})
}
