package model

import ( 

	"gorm.io/gorm"
	"diary_api/database"


)

type Entry struct {
	gorm.Model
	Content string `gorm: "type:text; " json: "content"`
	UserID uint

}
//creating an entry
func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
//updating an entry 
func (entry *Entry) Update() (*Entry, error) {
	err := database.Database.Save(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
//deleting an entry
func (entry *Entry) Delete(id string) (*Entry, error) {
	err := database.Database.Delete(&entry, id).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

// detail of an entry 
func (entry *Entry) Detail(id string) (*Entry, error) {
	err := database.Database.First(&entry, id).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}






