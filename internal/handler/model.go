package handler

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RequestID struct {
	ID int `json:"id"`
}
