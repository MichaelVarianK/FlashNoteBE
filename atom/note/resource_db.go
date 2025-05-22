package atom_note

import (
	"fmt"

	"github.com/KelXR/FlashNoteBE/config"
)

func GetAllFoldersAndNotesDB() (map[string]interface{}, error) {
	var folders []Folder
	var ungroupedNotes []Note

	if err := config.DB.Find(&folders).Error; err != nil  {
		return nil, err
	}

	if err := config.DB.
		Where("folder_id IS NULL or folder_id = 0").
		Find(&ungroupedNotes).Error; err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"folders": folders,
		"ungrouped_notes": ungroupedNotes,
	}
	
	return result, nil
}

func GetNotesByFolderIdDB(id string) ([]Note, error) {
	var notes []Note

	if err := config.DB.Where("folder_id = ?", id).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func SearchFoldersAndNotesDB(query string) ([]SearchSuggestion, error) {
	var suggestions []SearchSuggestion
	var folders []Folder
	var notes []Note

	likeQuery := "%" + query + "%"

	if err := config.DB.
		Where("folder_name ILIKE ?", likeQuery).
		Find(&folders).Error; err != nil {
			return nil, err
	}

	for _, f := range folders {
		suggestions = append(suggestions, SearchSuggestion{
			Id: f.FolderId,
			Title: f.FolderName,
			Type: "Folder",
		})
	}

	if err := config.DB.
		Where("note_title ILIKE ?", likeQuery).
		Find(&notes).Error; err != nil {
			return nil, err
	}

	for _, n := range notes {
		suggestions = append(suggestions, SearchSuggestion{
			Id: n.NoteId,
			Title: n.NoteTitle,
			Type: "Note",
		})
	}

	return suggestions, nil

}

func CreateNoteDB(note *Note, userIDs []uint) error {
	if len(userIDs) > 0 {
		var users []User
		if err := config.DB.Where("user_id IN ?", userIDs).Find(&users).Error; err != nil {
			return err
		}
		note.Users = users
	}
	return config.DB.Create(note).Error
}

func FindNoteByID(id string) (*Note, error) {
	var note Note
	result := config.DB.First(&note, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("note with ID %s not found", id)
	}
	return &note, nil
}

func UpdateNoteDB(id string, req *UpdateNoteRequest) error {
	existingNote, err := FindNoteByID(id)
	if err != nil {
		return err
	}

	existingNote.NoteTitle = req.NoteTitle
	existingNote.NoteContent = req.NoteContent
	existingNote.FolderId = req.FolderId
	
	if err := config.DB.Save(&existingNote).Error; err != nil {
		return err
	}

	if len(req.UserIDs) > 0 {
		var users []User
		if err := config.DB.Where("user_id IN ?", req.UserIDs).Find(&users).Error; err != nil {
			return err
		}

		if err := config.DB.Model(&existingNote).Association("Users").Replace(users); err != nil {
			return err
		}
	}

	return nil
}

func DeleteNoteDB(id string) error {
	existingNote, err := FindNoteByID(id)
	if err != nil {
		return err
	}

	return config.DB.Delete(&existingNote, id).Error
}