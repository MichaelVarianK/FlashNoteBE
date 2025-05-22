package atom_folder

import "log"



func CreateFolderUseCase(folder Folder) (error) {
	if err := CreateFolderDB(&folder); err != nil {
		log.Print("[atom][folder][resource.go][CreateFolderUseCase] Error: ", err)
		return err
	}
	return nil
}

func UpdateFolderUseCase(id string, folder Folder) (error) {
	if err := UpdateFolderDB(id, &folder); err != nil {
		log.Print("[atom][folder][resource.go][UpdateFolderUseCase] Error: ", err)
		return err
	}
	
	return nil
}

func DeleteFolderUseCase(id string) (error) {
	if err := DeleteFolderDB(id); err != nil {
		log.Print("[atom][folder][resource.go][DeleteFolderUseCase] Error: ", err)
		return err
	}

	return nil
}