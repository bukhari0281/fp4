package response

const DEFAULT_ERROR_MESSAGE_UNAUTHORIZED = "You are not authorized."
const DEFAULT_ERROR_BAD_REQUEST = "Bad Request. Check again the header or request body"
const DEFAULT_ERROR_TOKEN_NOT_VALID = "Token is not valid."

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BuildErrResponse(message string, code int) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}
