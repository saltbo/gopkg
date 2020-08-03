package httputil

// JSONResponse
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// JSONResponse
func NewJSONResponse(data interface{}) *JSONResponse {
	return &JSONResponse{Code: 0, Msg: "ok", Data: data}
}
