package companies

import (
	"context"
	"fmt"
	mathRand "math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sbuttigieg/test-xm/xm_app/services"
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

	mathRand.Seed(time.Now().UnixNano())
	a := mathRand.Intn(100000)
	fmt.Println("**** RANDOM **** ", a)

	kk, err := services.Produce(context.Background(), "hello", a)
	fmt.Println("**** KAFKA **** ", kk, err)

	c.JSON(http.StatusOK, OKMsg{
		Code: http.StatusOK,
		Data: company,
	})
}
