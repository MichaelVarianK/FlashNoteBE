package routes

import (
	"github.com/gin-gonic/gin"
	atom_note "github.com/KelXR/FlashNoteBE/atom/note/controllers"
)

func NoteRoute(router *gin.Engine) {
	router.GET("/get/all-folders-notes", atom_note.GetAllFoldersAndNotes)
	router.GET("/get/notes-by-folder-id/:id", atom_note.GetNotesByFolderId)
	router.GET("/get/search-folders-notes", atom_note.SearchFoldersAndNotesDB)
	router.POST("/post/create-note", atom_note.CreateNote)
	router.DELETE("/delete/delete-note/:id", atom_note.DeleteNote)
	router.PUT("/put/update-note/:id", atom_note.UpdateNote)
}