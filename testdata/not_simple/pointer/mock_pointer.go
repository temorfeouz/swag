package pointer

type update struct {
	ID1      *int    `schema:"id1,required" swagger:"merchant's id'"`
	Name1    *string `schema:"name1" swagger:"Merchant name"`
	Testing1 *bool   `schema:"testing1" swagger:"Is merchant in test mode"`
	ID2      int     `schema:"id2,required" swagger:"merchant's id'"`
	Name2    string  `schema:"name2" swagger:"Merchant name"`
	Testing2 bool    `schema:"testing2" swagger:"Is merchant in test mode"`
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

// MOCK. godoc
// @Summary MOCK.
// @tags merchant
// @Description MOCK.
// @Accept x-www-form-urlencoded
// @Produce json
// @Auth true
// @Params pointer.CardList query
// @Success 200 {object} pointer.emptySuccess
// @Failure 500 {object} pointer.Error
// @Router /some/update [post]
func Update() {}
