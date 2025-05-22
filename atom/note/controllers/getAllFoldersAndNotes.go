package atom_note

import (
	"log"
	"net/http"

	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/gin-gonic/gin"
)

func GetAllFoldersAndNotes(context *gin.Context) {
	folders, err := atom_note.GetAllFoldersAndNotesUseCase()

	if err != nil {
		log.Print("[atom][note][controllers][getAllFoldersAndNotes] Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve all folders and ungrouped notes",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    folders,
		"message": "Folders retrieved successfully",
	})
}
