package apimodels

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResLogin struct {
	Token string `json:"token"`
}
