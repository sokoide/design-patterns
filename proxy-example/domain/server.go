package domain

type Server interface {
	HandleRequest(url, method string) (int, string)
}

type Logger interface {
	Log(message string)
}
