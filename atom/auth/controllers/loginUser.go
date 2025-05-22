package atom_auth

import (
	"net/http"

	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var input atom_auth.LoginInput

	if inputError := c.ShouldBindJSON(&input); inputError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  inputError.Error(),
		})
		return
	}

	user, err := atom_auth.LoginUserUseCase(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   user,
	})
}
