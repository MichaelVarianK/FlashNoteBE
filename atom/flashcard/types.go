package atom_flashcard

import "time"

type Folder struct {
	FolderId uint `gorm:"primaryKey" json:"folder_id"`
	FolderName string `gorm:"not null" json:"folder_name"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`

	Notes []Note `gorm:"foreignKey:FolderId;references:FolderId" json:"notes,omitempty"`
}

type Note struct {
	NoteId uint `gorm:"primaryKey" json:"note_id"`
	FolderId uint `json:"folder_id"`
	NoteTitle string `gorm:"not null" json:"note_title"`
	NoteContent string `gorm:"not null" json:"note_content"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`
	
	Folder *Folder `gorm:"foreignKey:FolderId;references:FolderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"folder,omitempty"`
	Flashcards []Flashcard `gorm:"foreignKey:NoteId;references:NoteId" json:"notes,omitempty"`
}

// Seperate

type FolderFlashcard struct {
	FolderFlashcardId uint `gorm:"primaryKey" json:"folder_flashcard_id"`
	FolderFlashcardName string `gorm:"not null" json:"folder_flashcard_name"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`

	FlashCards []Flashcard `gorm:"foreignKey:FolderFlashcardId;references:FolderFlashcardId" json:"notes,omitempty"`
}

type Flashcard struct {
	FlashcardId uint `gorm:"primaryKey" json:"flashcard_id"`
	NoteId uint `json:"note_id"`
	FolderFlashcardId uint `json:"folder_flashcard_id"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`

	Note *Note `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"note,omitempty"`
	FolderFlashcard *FolderFlashcard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"folder_flashcard,omitempty"`
	Questions []Question `gorm:"foreignKey:FlashcardId;references:FlashcardId" json:"notes,omitempty"`
}

type Question struct {
	QuestionId uint `gorm:"primaryKey" json:"question_id"`
	FlashcardId uint `json:"flashcard_id"`
	Question string `gorm:"not null" json:"question"`
	Answer string `gorm:"not null" json:"answer"`

	Flashcard *Flashcard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

