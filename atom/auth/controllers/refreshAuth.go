package atom_auth

import (
	"net/http"

	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth"
	"github.com/gin-gonic/gin"
)

func RefreshAuth(c *gin.Context) {
	var input atom_auth.RefreshTokenInput

	if inputError := c.ShouldBindJSON(&input); inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  inputError.Error(),
		})
		return
	}

	tokens, err := atom_auth.RefreshAuthUseCase(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"tokens": tokens,
	})
}
