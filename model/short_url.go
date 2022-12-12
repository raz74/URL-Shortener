package model

type ShortendUrl struct {
	Id       int    `json:"id"`
	Redirect string `json:"redirect"`
	Url      string `json:"url"`
}

type UrlCreationRequest struct {
	LongUrl string `json:"url"`

	//Random bool `json:"random"`
}
