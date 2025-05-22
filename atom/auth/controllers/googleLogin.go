package atom_auth

import (
	"net/http"

	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth"
	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	state, authURL, err := atom_auth.GoogleLoginUserCase()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.SetCookie("oauthstate", state, 3600, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}
