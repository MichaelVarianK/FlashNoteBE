package atom_folder

import (
	"net/http"

	atom_folder "github.com/KelXR/FlashNoteBE/atom/folder"
	"github.com/gin-gonic/gin"
)

func DeleteFolder(context *gin.Context) {
	id := context.Param("id")

	if err := atom_folder.DeleteFolderUseCase(id); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to delete folder",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Delete folder successfully",
	})
}