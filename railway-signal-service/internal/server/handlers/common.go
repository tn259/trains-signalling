package handlers

type errorResponse struct {
	Error string `json:"error"`
}

// newErrorResponse wraps an error that can be returned as a JSON response.
func newErrorResponse(err error) errorResponse {
	return errorResponse{
		Error: err.Error(),
	}
}
