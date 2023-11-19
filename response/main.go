package response

type ResponseView struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Response(message string) ResponseView {
	return ResponseView{
		Success: true,
		Message: message,
	}
}

type ErrorView struct {
	Success bool   `json:"success"`
	Error   string `json:"err"`
}

func Error(err string) ErrorView {
	return ErrorView{
		Success: false,
		Error:   err,
	}
}
