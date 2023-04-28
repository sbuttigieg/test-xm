package companies

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

type UpdateRequest struct {
	Company *models.Company `json:"company" binding:"required"`
	Fields  []string        `json:"fields" binding:"required"`
}

func (h *Handler) Update(c *gin.Context) {
	var req *UpdateRequest

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, ErrMsg{
			Code:  http.StatusBadRequest,
			Error: InvalidRequest,
		})

		return
	}

	id := c.Param("id")

	company, err := h.service.Update(c.Request.Context(), id, req.Company, req.Fields)
	if err != nil {
		c.Abort()

		switch {
		case strings.Contains(err.Error(), FieldError):
			c.JSON(http.StatusBadRequest, ErrMsg{
				Code:  http.StatusBadRequest,
				Error: InvalidRequest,
			})
		default:
			c.JSON(http.StatusInternalServerError, ErrMsg{
				Code:  http.StatusInternalServerError,
				Error: UpdateError,
			})
		}

		return
	}

	c.JSON(http.StatusOK, OKMsg{
		Code: http.StatusOK,
		Data: company,
	})
}
