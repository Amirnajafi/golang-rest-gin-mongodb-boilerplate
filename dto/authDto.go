package dto

type AuthDto struct {
	// email field string should be unique
	Email    string `json:"email,omitempty" validate:"required" unique:"true" email:"true"`
	Password string `json:"password" binding:"required" minLength:"6"`
}
