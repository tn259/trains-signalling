package handlers

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(err error) errorResponse {
	return errorResponse{
		Error: err.Error(),
	}
}
