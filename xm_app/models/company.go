package models

type Company struct {
	ID          string `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(15);not null;unique;"`
	Description string `json:"description" gorm:"type:varchar(3000);"`
	Employees   int    `json:"employees" gorm:"type:int;not null;"`
	Registered  bool   `json:"registered" gorm:"type:bool;not null;default:false"`
	Type        string `json:"last_activity" gorm:"type:varchar(30);not null;"`
}
