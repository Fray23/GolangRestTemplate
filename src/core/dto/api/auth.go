package dto

type SignUpDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}

type SignUpResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}

