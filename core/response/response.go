package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
	Code int `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}

type PageableParam struct {
	TotalElements int `json:"total_elements"`
	PageSize int `json:"page_size"`
	PrevPage int `json:"prev_page"`
	CurrentPage int `json:"current_page"`
}

type PageableResponse struct {
	Content interface{} `json:"content"`
	MetaData PageableParam `json:"meta_data"`
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, status int, success bool, message string, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := Response{
		Success: success,
		Code: status,
		Message: message,
		Data: body,
	}

	_ = json.NewEncoder(w).Encode(response)
}

func ErrorOccured(w http.ResponseWriter, status int, success bool, message string, body interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Errors: body,
		Message: message,
		Success: success,
		Code: status,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func JSONPageableResponse(w http.ResponseWriter, status int, message string, content interface{}, meta PageableParam) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	pgResp := PageableResponse{
		MetaData: meta,
		Content: content,
		Message: message,
	}

	_ = json.NewEncoder(w).Encode(pgResp)
}
