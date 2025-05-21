package atom_folder

import (
	"net/http"

	atom_folder "github.com/KelXR/FlashNoteBE/atom/folder"
	"github.com/gin-gonic/gin"
)

func UpdateFolder(context *gin.Context) {
	var folder atom_folder.Folder
	id := context.Param("id")

	if err := context.BindJSON(&folder); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Invalid JSON",
		})
		return
	}

	if err := atom_folder.UpdateFolderUseCase(id, folder); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to update folder",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Folder update successfully",
	})
}