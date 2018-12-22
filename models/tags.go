package models

type Tag struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	Name   string `json:"name" binding:"required,numeric"`
	Model
}

func (t *Tag) StoreTag() error {
	err := db.Insert(t)
	if err != nil {
		return err
	}
	return nil
}