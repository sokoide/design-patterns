package domain

type Server interface {
	HandleRequest(url, method string) (int, string)
}
