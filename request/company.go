package request

import validation "github.com/go-ozzo/ozzo-validation"

type (
	CompanyRequest struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
		Address string `json:"address"`
	}

	TopupCompanyBalance struct {
		Balance int `json:"balance" validate:"required"`
	}
)

// TODO: tuliskan validasi untuk CompanyRequest dengan rule semua field required
func (r CompanyRequest) Validate() interface{} {

	return nil
}

func (req TopupCompanyBalance) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Balance, validation.Required),
	)
}
