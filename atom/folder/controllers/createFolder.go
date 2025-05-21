package atom_folder

import (
	"net/http"
	atom_folder "github.com/KelXR/FlashNoteBE/atom/folder"
	"github.com/gin-gonic/gin"
)

func CreateFolder(context *gin.Context) {
	var folder atom_folder.Folder

	if err := context.BindJSON(&folder); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request body",
		})
		return
	}

	if err := atom_folder.CreateFolderUseCase(folder); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create folder",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Folder created successfully",
	})
}
