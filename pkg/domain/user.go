package domain

type Users struct {
	ID    uint   `json:"id" gorm:"unique;not null"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
