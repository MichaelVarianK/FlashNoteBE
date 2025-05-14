package atom_note

import "time"

type User struct {
	UserId uint `gorm:"primaryKey" json:"user_id"`
	FullName string `gorm:"not null" json:"full_name"`
	Email string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedDate time.Time `json:"created_date"`

	Notes []Note `gorm:"many2many:UserNotes;"`
}

type Folder struct {
	FolderId uint `gorm:"primaryKey" json:"folder_id"`
	FolderName string `gorm:"not null" json:"folder_name"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`

	Notes []Note `gorm:"foreignKey:FolderId;references:FolderId" json:"notes,omitempty"`
}

// Seperate

type Note struct {
	NoteId uint `gorm:"primaryKey" json:"note_id"`
	FolderId uint `json:"folder_id"`
	NoteTitle string `gorm:"not null" json:"note_title"`
	NoteContent string `gorm:"not null" json:"note_content"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`

	Users []User `gorm:"many2many:UserNotes;"`
	Folder *Folder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"folder,omitempty"`
}