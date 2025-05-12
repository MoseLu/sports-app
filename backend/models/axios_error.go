package models

import "time"

// AxiosError 前端错误日志模型
type AxiosError struct {
	ErrorType    string    `json:"error_type"`
	Message      string    `json:"message"`
	StatusCode   int       `json:"status_code"`
	Path         string    `json:"path"`
	Method       string    `json:"method"`
	RequestBody  string    `json:"request_body"`
	ResponseBody string    `json:"response_body"`
	CreatedAt    time.Time `json:"created_at"`
} 