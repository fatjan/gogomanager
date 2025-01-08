package delivery

type Response struct {
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
