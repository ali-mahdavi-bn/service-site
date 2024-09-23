package entrypoints

type (
	UserRequestModel struct {
		Name  string `json:"name" form:"name" validate:"required,min=2"`
		Email string `json:"email" form:"email" validate:"required,email"`
	}
)
