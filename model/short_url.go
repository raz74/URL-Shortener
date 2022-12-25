package model

type ShortedUrl struct {
	Id int `json:"id"`
	// Redirect    string `json:"redirect"`
	Url        string `json:"url"`
	ShortedUrl string `json:"shorted_url"`
}

type UrlCreationRequest struct {
	LongUrl string `json:"url"`

	//Random bool `json:"random"`
}