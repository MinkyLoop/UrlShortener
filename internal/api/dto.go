package api

type CreateShortenReq struct {
	Url string `json:"url" binding:"required"`
}

type CreateShortenRes struct {
	ShortUrl string `json:"shorten_url"`
}

type ErrorRes struct {
	Error string `json:"error"`
}
