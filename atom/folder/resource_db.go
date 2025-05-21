package atom_folder

import (
	"fmt"

	"github.com/KelXR/FlashNoteBE/config"
)

func CreateFolderDB(folder *Folder) (error) {
	return config.DB.Create(folder).Error
}

func FindFolderByID(id string) (*Folder, error) {
	var folder Folder
	result := config.DB.First(&folder, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("folder with ID %s not found", id)
	}
	return &folder, nil
}

func UpdateFolderDB(id string, folder *Folder) (error) {
	existingFolder, err := FindFolderByID(id)
	if err != nil {
		return err
	}

	existingFolder.FolderName = folder.FolderName

	if err := config.DB.Save(&existingFolder).Error; err != nil {
		return err
	}

	return nil
}

func DeleteFolderDB(id string) (error) {
	existingFolder, err := FindFolderByID(id)
	if err != nil {
		return err
	}

	return config.DB.Delete(&existingFolder, id).Error
}