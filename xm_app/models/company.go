package models

import (
	"encoding/json"
	"time"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
)

type Company struct {
	ID          string      `json:"id"`
	Name        string      `json:"name" binding:"required,max=15"`
	Description string      `json:"description" binding:"max=3000"`
	Employees   int         `json:"employees"  binding:"required"`
	Registered  *bool       `json:"registered" binding:"required"`
	Type        CompanyType `json:"type" binding:"required"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type CompanyType string

const (
	CompanyTypeCorporations       CompanyType = "Corporations"
	CompanyTypeNonProfit          CompanyType = "NonProfit"
	CompanyTypeCooperative        CompanyType = "Cooperative"
	CompanyTypeSoleProprietorship CompanyType = "Sole Proprietorship"
)

func (t *CompanyType) Parse(s string) error {
	switch s {
	case "Corporations":
		*t = CompanyTypeCorporations
	case "NonProfit":
		*t = CompanyTypeNonProfit
	case "Cooperative":
		*t = CompanyTypeCooperative
	case "Sole Proprietorship":
		*t = CompanyTypeSoleProprietorship
	default:
		return errors.BadRequest("invalid Company type '%s'", s)
	}

	return nil
}

func (t CompanyType) String() string {
	return string(t)
}

func (t *CompanyType) UnmarshalJSON(b []byte) error {
	var compType string

	err := json.Unmarshal(b, &compType)
	if err != nil {
		return errors.Internal("JSON marshaling of company type failed '%s'", err)
	}

	return t.Parse(compType)
}

var CompanyField = struct {
	ID          string
	Name        string
	Description string
	Employees   string
	Registered  string
	Type        string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	Employees:   "employees",
	Registered:  "registered",
	Type:        "type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var CompanyUpdateFields = map[string]bool{
	CompanyField.ID:          false,
	CompanyField.Name:        true,
	CompanyField.Description: true,
	CompanyField.Employees:   true,
	CompanyField.Registered:  true,
	CompanyField.Type:        true,
	CompanyField.CreatedAt:   false,
	CompanyField.UpdatedAt:   false,
}
