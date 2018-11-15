package response

type CardList struct {
	Result bool         `json:"result"`
	Data   CardListData `json:"data"`
}

type CardListData []Card
type Card struct {
	ID         int    `json:"id"`
	PAN        string `json:"pan_last_4"`
	ExpireDate string `json:"expire_date"`
}
type emptySuccess struct {
	Result bool `json:"result"`
}

type Error struct {
	Result bool
	Error  struct {
		Message   string
		Code      int
		Exception string
	}
}
