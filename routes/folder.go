package routes

import (
	"github.com/gin-gonic/gin"
	atom_folder "github.com/KelXR/FlashNoteBE/atom/folder/controllers"
)

func FolderRoute(router *gin.Engine) {
	router.POST("/post/create-folder", atom_folder.CreateFolder)
	router.PUT("/put/update-folder/:id", atom_folder.UpdateFolder)
	router.DELETE("/delete/delete-folder/:id", atom_folder.DeleteFolder)
}