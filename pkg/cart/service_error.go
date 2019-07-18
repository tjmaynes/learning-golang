package cart

// ServiceStatusCode ..
type ServiceStatusCode string

// ServiceError ..
type ServiceError interface {
	Message() string
	StatusCode() ServiceStatusCode
}

const (
	// ItemNotFound ..
	ItemNotFound ServiceStatusCode = "ItemNotFound"

	// InvalidItem ..
	InvalidItem ServiceStatusCode = "InvalidItem"

	// UnknownException ..
	UnknownException ServiceStatusCode = "UnknownException"
)

// CreateServiceError ..
func CreateServiceError(message string, statusCode ServiceStatusCode) ServiceError {
	return &serviceError{message: message, statusCode: statusCode}
}

type serviceError struct {
	message    string
	statusCode ServiceStatusCode
}

func (s *serviceError) Message() string {
	return s.message
}

func (s *serviceError) StatusCode() ServiceStatusCode {
	return s.statusCode
}
