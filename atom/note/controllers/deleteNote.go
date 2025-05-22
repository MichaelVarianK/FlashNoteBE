package atom_note

import (
	"net/http"

	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/gin-gonic/gin"
)

func DeleteNote(context *gin.Context) {
	id := context.Param("id")

	if err := atom_note.DeleteNoteUseCase(id); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to delete note",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Delete note successfully",
	})
}