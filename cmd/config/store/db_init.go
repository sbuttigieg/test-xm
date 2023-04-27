package store

import (
	"gorm.io/gorm"

	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func DBInit(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Company{})
	if err != nil {
		return err
	}

	var count int64

	db.Table("companies").Count(&count)

	if count == 0 {
		var companies = []models.Company{
			{ID: "6cc4ee0d-9919-4857-a70d-9b7283957e16", Name: "Google", Description: "house of Golang", Employees: 50000, Registered: true, Type: "Corporations"},
		}

		db.Create(&companies)
	}

	return nil
}
