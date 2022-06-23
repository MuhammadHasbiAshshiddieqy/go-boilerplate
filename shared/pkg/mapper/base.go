package mapper

import "microservice/shared/dto"

func BaseResponse(status, message string, data interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
