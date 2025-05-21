package atom_note

import (
	"log"
	"net/http"

	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/gin-gonic/gin"
)

func GetNotesByFolderId(context *gin.Context) {
	id := context.Param("id")
	var notes []atom_note.Note

	notes, err := atom_note.GetNotesByFolderIdUseCase(id)
	if (err != nil) {
		log.Print("[atom][note][controllers][getNotesByFolderId] Error: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to retrieve notes",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": notes,
		"message": "Retrieve notes successfully",
	})
}