package models

import (
	"encoding/json"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
)

type Company struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Employees   int         `json:"employees"`
	Registered  bool        `json:"registered"`
	Type        CompanyType `json:"type"`
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
