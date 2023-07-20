package handler

type UserMultiGetQuery struct {
	Keys []string `query:"keys" validate:"required,min=1,dive"`
}
