package atom_note

import "log"

func GetAllFoldersAndNotesUseCase() (map[string]interface{}, error) {
	result, err := GetAllFoldersAndNotesDB()
	if err != nil {
		log.Print("[atom][note][resource.go][GetAllFoldersUseCase] Error: ", err)
		return nil, err
	}

	return result, nil
}

func GetNotesByFolderIdUseCase(id string) ([]Note, error) {
	return GetNotesByFolderIdDB(id)
}

func SearchFoldersAndNotesUseCase(query string) ([]SearchSuggestion, error) {
	return SearchFoldersAndNotesDB(query)
}

func CreateNoteUseCase(note Note, userIDs []uint) error {
	return CreateNoteDB(&note, userIDs)
}

func UpdateNoteUseCase(id string, req UpdateNoteRequest) error {
	return UpdateNoteDB(id, &req)
}

func DeleteNoteUseCase(id string) error {
	return DeleteNoteDB(id)
}