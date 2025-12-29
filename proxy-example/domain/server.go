package domain

// Server handles incoming requests.
type Server interface {
	HandleRequest(url, method string) (int, string)
}

// Logger abstracts logging for the domain.
type Logger interface {
	Log(message string)
}
