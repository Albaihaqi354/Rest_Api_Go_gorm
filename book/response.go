package book

type BookResponse struct {
	Id          int    `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
	Rating      int    `json:"Rating"`
}
