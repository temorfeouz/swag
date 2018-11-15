package response

type (
	FeeRate struct {
		Result bool        `json:"result"`
		Data   FeeRateData `json:"data"`
	}
	FeeRateData struct {
		Deposit  CurrencyRate `json:"deposit"`
		Withdraw CurrencyRate `json:"withdraw"`
	}
	CurrencyRate map[string]PS
	PS           map[string]PSFeeRate

	PSFeeRate struct {
		OurFee FeeRateInfo `json:"our_fee"`
		PSFee  FeeRateInfo `json:"ps_fee"`
		Limits Limit       `json:"limits"`
	}
	FeeRateInfo struct {
		Min     string `json:"min"`
		Max     string `json:"max"`
		Percent string `json:"percent"`
	}
	Limit struct {
		Check string `json:"check"`
		Min   string `json:"min"`
		Max   string `json:"max"`
	}
)

type Error struct {
	Result bool
	Error  struct {
		Message   string
		Code      int
		Exception string
	}
}
