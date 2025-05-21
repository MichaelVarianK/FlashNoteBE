package atom_note

import (
	"net/http"

	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/gin-gonic/gin"
)

func UpdateNote(context *gin.Context) {
	var req atom_note.UpdateNoteRequest
	id := context.Param("id")

	if err := context.BindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Invalid JSON",
		})
		return
	}

	if err := atom_note.UpdateNoteUseCase(id, req); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to update note",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Note update successfully",
	})
}