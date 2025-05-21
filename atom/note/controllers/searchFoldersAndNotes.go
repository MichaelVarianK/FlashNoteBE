package atom_note

import (
	"net/http"

	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/gin-gonic/gin"
)

func SearchFoldersAndNotesDB(context *gin.Context) {
	query := context.Query("q")

	if query == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Query parameter 'q' is required",
		})
		return
	}

	suggestions, err := atom_note.SearchFoldersAndNotesUseCase(query)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to search folders and notes",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": suggestions,
		"message": "Search folders and notes successfully",
	})
}