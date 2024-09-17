package dto

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RepositoryResult struct {
	Result interface{}
	Error  error
}
