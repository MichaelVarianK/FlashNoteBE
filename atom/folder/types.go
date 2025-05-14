package atom_folder

import "time"

type Folder struct {
	FolderId uint `gorm:"primaryKey" json:"folder_id"`
	FolderName string `gorm:"not null" json:"folder_name"`
	CreatedDate time.Time `json:"created_date"`
	DeletedDate time.Time `json:"deleted_date"`
}