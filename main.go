package main

import (
	"log"

	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth"
	atom_flashcard "github.com/KelXR/FlashNoteBE/atom/flashcard"
	atom_folder "github.com/KelXR/FlashNoteBE/atom/folder"
	atom_note "github.com/KelXR/FlashNoteBE/atom/note"
	"github.com/KelXR/FlashNoteBE/config"
	routes "github.com/KelXR/FlashNoteBE/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// middleware "github.com/KelXR/FlashNoteBE/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitGoogleAuth()

	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&atom_flashcard.FolderFlashcard{},
		&atom_folder.Folder{},
		&atom_auth.User{},
		&atom_note.Note{},
		&atom_flashcard.Flashcard{},
		&atom_flashcard.Question{},
		&atom_auth.RefreshToken{},
	)

	router := gin.Default()

	router.Use(routes.CorsMiddleware())

	routes.AuthRoutes(router)

	// protected := router.Group("/api/protected", middleware.AuthMiddleware())

	router.Run(":8080")
}
