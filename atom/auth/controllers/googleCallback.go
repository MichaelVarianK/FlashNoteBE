package atom_auth

import (
	"net/http"

	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth"
	"github.com/gin-gonic/gin"
)

func GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	storedState, _ := c.Cookie("oauthstate")

	if state != storedState {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid OAuth state parameter",
		})
		return
	}

	code := c.Query("code")
	result, err := atom_auth.GoogleCallbackUseCase(c.Request.Context(), code, state, storedState)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": result,
	})
}
