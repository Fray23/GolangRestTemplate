package dto

type SignUpDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}

