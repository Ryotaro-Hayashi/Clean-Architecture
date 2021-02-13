package interfaces

type Error struct {
	Code        int
	Message     string
	Description string
}

func ResponseError(code int, message string, description string) Error {
	resErr := Error{
		Code:        code,
		Message:     message,
		Description: description,
	}

	return resErr
}
