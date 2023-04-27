package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

func (s *store) Update(ctx context.Context, id string, data *models.Company, fields []string) (*models.Company, error) {
	company, err := s.Get(ctx, id)
	if err != nil {
		return nil, errors.BadRequestWrap(err, "company get")
	}

	updateFields := make([]string, 0, len(fields))
	updateArgs := make([]interface{}, 0, len(fields))

	for i := range fields {
		switch fields[i] {
		case models.CompanyField.Name:
			company.Name = data.Name
			updateArgs = append(updateArgs, data.Name)
		case models.CompanyField.Description:
			company.Description = data.Description
			updateArgs = append(updateArgs, data.Description)
		case models.CompanyField.Employees:
			company.Employees = data.Employees
			updateArgs = append(updateArgs, data.Employees)
		case models.CompanyField.Registered:
			company.Registered = data.Registered
			updateArgs = append(updateArgs, data.Registered)
		case models.CompanyField.Type:
			company.Type = data.Type
			updateArgs = append(updateArgs, data.Type)
		case models.CompanyField.UpdatedAt:
			company.UpdatedAt = data.UpdatedAt
			updateArgs = append(updateArgs, data.UpdatedAt)
		}

		updateFields = append(updateFields, fmt.Sprintf("%s = $%d", fields[i], i+1))
	}

	updateArgs = append(updateArgs, id)

	_, err = s.db.ExecContext(ctx, fmt.Sprintf(`
		UPDATE companies
		SET %s
		WHERE id = $%d;`, strings.Join(updateFields, ","), len(fields)+1), updateArgs...)
	if err != nil {
		return nil, errors.BadRequestWrap(err, "update company")
	}

	return company, nil
}
