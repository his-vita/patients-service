package model

type Contact struct {
	WorkPhoneNumber *string `json:"work_phone_number"`
	PhoneNumber     *string `json:"phone_number"`
	Email           *string `json:"email"`
}
