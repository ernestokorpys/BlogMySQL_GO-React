package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserID string `json:"userid"`
	User   User   `json:"user";gorm:"foreignKey:UserID"` //Revisa que el id exista en la otra tabla
}
