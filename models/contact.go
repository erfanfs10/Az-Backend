package models

type ContactCreate struct {
	Name    *string `db:"name" json:"name" form:"name" validate:"required"`
	Email   *string `db:"email" json:"email" form:"email" validate:"required,email"`
	Subject *string `db:"subject" json:"subject" form:"subject" validate:"required"`
	Message *string `db:"message" json:"message" form:"message" validate:"required"`
}
