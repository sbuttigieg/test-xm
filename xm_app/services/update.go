package services

import (
	"context"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *service) Update(ctx context.Context, id string, data *models.Company, fields []string) (*models.Company, error) {
	if len(fields) == 0 {
		return nil, errors.BadRequest("no fields to update")
	}

	// check if fields are valid
	for i := range fields {
		enabled, ok := models.CompanyUpdateFields[fields[i]]
		if !ok {
			return nil, errors.BadRequest("requested field not found:'%s'", fields[i])
		}

		if !enabled {
			return nil, errors.BadRequest("not able to update field:'%s'", fields[i])
		}
	}

	// add modified
	fields = append(fields, models.CompanyField.UpdatedAt)
	data.UpdatedAt = s.timeFunc()

	company, err := s.store.Update(ctx, id, data, fields)
	if err != nil {
		return nil, errors.BadRequestWrap(err, "company update")
	}

	return company, nil
}
