package mantra_request

type CreateMantraReqDTO struct {
	Content string `json:"content"`
	Speaker string `json:"speaker"`
	Writer  string `json:"writer"`
}
