package atom_note

import (
	"net/http"
	"time"

	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/gin-gonic/gin"
)

func CreateNote(context *gin.Context) {
	var req atom_note.CreateNoteRequest

	if err := context.BindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	note := atom_note.Note {
		NoteTitle:   req.NoteTitle,
		NoteContent: req.NoteContent,
		FolderId:    req.FolderId,
		CreatedDate: time.Now(),
	}

	if err := atom_note.CreateNoteUseCase(note, req.UserIDs); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Failed to create note",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Note created successfully",
	})
}