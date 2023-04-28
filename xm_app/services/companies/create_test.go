package companies_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies"
	"github.com/sbuttigieg/test-xm/xm_app/services/companies/mocks"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		it  string
		req *models.Company

		CreateInput    *models.Company
		CreateResultID string
		CreateError    error

		expectedError    string
		expectedResultID string
	}{
		{
			it: "it returns the id of the Company created",
			req: &models.Company{
				ID:   "mock-id",
				Name: "mock-name",
			},
			CreateInput: &models.Company{
				ID:   "mock-id",
				Name: "mock-name",
			},
			CreateResultID:   "5e0771f5-9e09-4f1f-a9f5-f0cbc4b296b4",
			expectedResultID: "5e0771f5-9e09-4f1f-a9f5-f0cbc4b296b4",
		},
		{
			it: "it returns error on store Create",
			req: &models.Company{
				ID:   "mock-id",
				Name: "mock-name",
			},
			CreateInput: &models.Company{
				ID:   "mock-id",
				Name: "mock-name",
			},
			CreateError:   errors.Wrap("mock-error"),
			expectedError: "create company: mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			checkIs := is.New(t)

			store := &mocks.StoreMock{
				CreateFunc: func(contextMoqParam context.Context, company *models.Company) (string, error) {
					checkIs.Equal(company, tc.CreateInput)

					return tc.CreateResultID, tc.CreateError
				},
			}

			service := companies.New(nil, nil, store, nil, nil)

			id, err := service.Create(context.Background(), tc.req)
			if err != nil {
				checkIs.Equal(err.Error(), tc.expectedError)
			}
			checkIs.Equal(id, tc.expectedResultID)
		})
	}
}
