package dto

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	TraceID   string      `json:"trace_id"`
	Timestamp int64       `json:"timestamp"`
}

type ErrorResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
	TraceID   string `json:"trace_id,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type ListResponse struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
}
