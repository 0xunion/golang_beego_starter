package types

type MasterResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) *MasterResponse {
	return &MasterResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func ErrorResponse(code int, message string) *MasterResponse {
	return &MasterResponse{
		Code:    code,
		Message: message,
	}
}

func (r *MasterResponse) IsSuccess() bool {
	return r.Code == 0
}

func (r *MasterResponse) IsError() bool {
	return r.Code < 0
}