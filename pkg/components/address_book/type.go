package address_book

type (
	// Response is only used for API's response
	Response struct {
		Message string      `json:"message,omitempty"`
		Status  int         `json:"status"`
		Data    interface{} `json:"data,omitempty"`
		Error   string      `json:"error,omitempty"`
	}
)
