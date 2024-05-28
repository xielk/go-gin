package controllers

// BaseController 是所有控制器的基础
type BaseController struct{}

// ListResponse 定义了通用的列表返回结构
type ListResponse struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

// NewListResponse 返回一个新的列表响应
func (bc *BaseController) NewListResponse(page, pageSize int, total int64, data interface{}) ListResponse {
	return ListResponse{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     data,
	}
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Code:    1,
		Message: message,
		Data:    data,
	}
}
