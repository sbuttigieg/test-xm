package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sbuttigieg/test-xm/xm_app/handler/helpers"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) GetToken(c *gin.Context) {
	var req *TokenRequest

	var user *models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Abort()

		c.JSON(http.StatusBadRequest, ErrMsg{
			Code:  http.StatusBadRequest,
			Error: InvalidRequest,
		})

		return
	}

	// check if email exists and password is correct
	user, err := h.service.GetByEmail(c.Request.Context(), req.Email)
	if err != nil {
		c.Abort()

		switch {
		case strings.Contains(err.Error(), NotFound):
			c.JSON(http.StatusNotFound, ErrMsg{
				Code:  http.StatusNotFound,
				Error: IncorrectCredentials,
			})
		default:
			c.JSON(http.StatusInternalServerError, ErrMsg{
				Code:  http.StatusInternalServerError,
				Error: GetTokenError,
			})
		}

		return
	}

	credentialError := user.CheckPassword(req.Password)
	if credentialError != nil {
		c.Abort()
		c.JSON(http.StatusUnauthorized, ErrMsg{
			Code:  http.StatusUnauthorized,
			Error: IncorrectCredentials,
		})

		return
	}

	token, err := helpers.GenerateJWT(h.config, user.Email, user.Username)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, ErrMsg{
			Code:  http.StatusInternalServerError,
			Error: GetTokenError,
		})

		return
	}

	c.JSON(http.StatusOK, OKMsg{
		Code: http.StatusOK,
		Data: fmt.Sprintf("token: %s", token),
	})
}
