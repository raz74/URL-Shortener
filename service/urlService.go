package service

type UrlService interface {
	AddUrl(string) (string, error)
	AddCustomUrl(string, string) (string, error)
	GetUrl(string) (string, bool)
}
