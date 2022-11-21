package apimodels

type ReqRegistration struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name" validate:"required"`
}

type ResRegistration struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
