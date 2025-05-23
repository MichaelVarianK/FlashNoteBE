package main

import (
	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth"
	atom_flashcard "github.com/KelXR/FlashNoteBE/atom/flashcard"
	atom_folder "github.com/KelXR/FlashNoteBE/atom/folder"
	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/KelXR/FlashNoteBE/config"
	"github.com/KelXR/FlashNoteBE/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&atom_flashcard.FolderFlashcard{},
		&atom_folder.Folder{},
		&atom_auth.User{}, 
		&atom_note.Note{},
		&atom_flashcard.Flashcard{},
		&atom_flashcard.Question{}, 
	)

	router := gin.Default()
	router.Use(routes.CorsMiddleware())
	routes.FolderRoute(router)
	routes.NoteRoute(router)
	router.Run(":8080")
}